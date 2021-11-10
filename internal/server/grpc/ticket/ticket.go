package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	"github.com/rodkevich/ts/internal/ticket/repository/postgres"
	pb "github.com/rodkevich/ts/proto/ticket/v1"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var db *postgres.Queries

type server struct {
	pb.UnimplementedTicketServiceServer
}

func (s server) CreateTicket(ctx context.Context, request *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	v := postgres.CreateTicketParams{
		OwnerID:     uuid.MustParse(request.Ticket.OwnerId),
		NameShort:   request.Ticket.NameShort,
		NameExt:     request.Ticket.NameExt,
		Description: &request.Ticket.Description,
		Amount:      request.Ticket.Amount,
		//
		Price: request.Ticket.Price,
		//
		Currency:  request.Ticket.GetCurrency(),
		Advantage: postgres.EnumTicketsAdvantagesType(request.Ticket.Advantage),
	}

	ticket, err := db.CreateTicket(ctx, v)
	if err != nil {
		return nil, err
	}

	grpc.SetHeader(ctx, metadata.Pairs("hostname", "hello yo", "hostname", "gprc serv bro"))
	grpc.SetTrailer(ctx, metadata.Pairs("how do we use this", "ha?"))

	return &pb.CreateTicketResponse{Id: ticket.String()}, nil

}

func (s server) GetTicket(ctx context.Context, request *pb.GetTicketRequest) (*pb.GetTicketResponse, error) {
	r, err := db.ReadTicket(ctx, uuid.MustParse(request.GetId()))
	if err != nil {
		return nil, err
	}

	var pub, del timestamppb.Timestamp

	if r.PublishedAt != nil {
		pub = *timestamppb.New(*r.PublishedAt)
	}

	if r.DeletedAt != nil {
		del = *timestamppb.New(*r.DeletedAt)
	}

	t := pb.Ticket{
		Id:          r.ID.String(),
		OwnerId:     r.OwnerID.String(),
		NameShort:   r.NameShort,
		NameExt:     r.NameExt,
		Description: *r.Description,
		Amount:      r.Amount,
		Price:       r.Price,
		Currency:    r.Currency,
		Active:      r.Active,
		Advantage:   string(r.Advantage),
		PublishedAt: &pub,
		CreatedAt:   timestamppb.New(r.CreatedAt),
		UpdatedAt:   timestamppb.New(r.UpdatedAt),
		DeletedAt:   &del,
	}

	rtn := pb.GetTicketResponse{Ticket: &t}
	return &rtn, nil
}

func (s server) UpdateTicket(ctx context.Context, request *pb.UpdateTicketRequest) (*pb.GetTicketResponse, error) {
	panic("implement me")
}

func (s server) DeleteTicket(ctx context.Context, request *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	panic("implement me")
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatalln(err)
	}

	db = postgres.New(conn)

	pb.RegisterTicketServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
