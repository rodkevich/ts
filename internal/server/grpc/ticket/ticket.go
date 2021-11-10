package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rodkevich/ts/internal/ticket/repository/postgres"
	pb "github.com/rodkevich/ts/proto/ticket/v1"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var db *postgres.Queries

type server struct {
	pb.UnimplementedTicketsServiceServer
}

func (s server) CreateTicket(ctx context.Context, request *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	v := postgres.CreateTicketParams{
		OwnerID:     uuid.MustParse(request.Ticket.GetOwnerId()),
		NameShort:   request.Ticket.GetNameShort(),
		NameExt:     request.Ticket.GetNameExt(),
		Description: request.Ticket.GetDescription(),
		Amount:      request.Ticket.GetAmount(),
		Price:       request.Ticket.GetPrice(),
		Currency:    request.Ticket.GetCurrency(),
		Advantage:   postgres.EnumTicketsAdvantagesType(request.Ticket.GetAdvantage()),
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

	t, err := db.ReadTicket(ctx, uuid.MustParse(request.GetId()))
	if err != nil {
		return nil, err
	}

	rtn := ticketAsProtobuf(t)

	return &pb.GetTicketResponse{Ticket: &rtn}, nil
}

func ticketAsProtobuf(it postgres.Ticket) pb.Ticket {
	var nullablePublished, nullableDeleted *timestamppb.Timestamp

	if it.PublishedAt != nil {
		nullablePublished = timestamppb.New(*it.PublishedAt)
	}
	if it.DeletedAt != nil {
		nullableDeleted = timestamppb.New(*it.DeletedAt)
	}

	return pb.Ticket{
		Id:          it.ID.String(),
		OwnerId:     it.OwnerID.String(),
		NameShort:   it.NameShort,
		NameExt:     it.NameExt,
		Description: it.Description,
		Amount:      it.Amount,
		Price:       it.Price,
		Currency:    it.Currency,
		Active:      it.Active,
		Advantage:   string(it.Advantage),
		PublishedAt: nullablePublished,
		CreatedAt:   timestamppb.New(it.CreatedAt),
		UpdatedAt:   timestamppb.New(it.UpdatedAt),
		DeletedAt:   nullableDeleted,
	}
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

	pb.RegisterTicketsServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
