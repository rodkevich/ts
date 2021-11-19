package main

import (
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/rodkevich/ts/customer/config"
	blueprints "github.com/rodkevich/ts/customer/internal/blueprints/grpc"
	repo "github.com/rodkevich/ts/customer/internal/repository/postgres"
	"github.com/rodkevich/ts/customer/internal/usage"
	"github.com/rodkevich/ts/customer/pkg/postgres"
	v1customer "github.com/rodkevich/ts/customer/proto/v1"
	serverGRPC "github.com/rodkevich/ts/server/grpc"
)

var configuration *config.Config

func init() {
	path := config.GetConfigPath(os.Getenv("config"))
	conf, err := config.GetConfig(path)
	if err != nil {
		log.Fatalf("unable to parse cfg path %v", err)

	}
	configuration = conf
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	builder := serverGRPC.GrpcServerBuilder{}
	builder.EnableReflection(true)

	s := builder.Build()
	s.RegisterService(customerServiceRegister)

	err := s.Start("0.0.0.0:50051")

	if err != nil {
		log.Fatalf("%v", err)
	}
	s.AwaitTermination(func() {
		log.Print("Shutting down the serverGRPC")
	})
}

func customerServiceRegister(grpcServer *grpc.Server) {
	pgxConn, err := postgres.NewPgxConn(configuration)
	if err != nil {
		log.Fatal("", err.Error())
	}
	// defer pgxConn.Close()

	customerRepository := repo.NewCustomerPG(pgxConn)
	customerUseCases := usage.New(customerRepository, nil)
	customerService := blueprints.GrpcCustomerService{UseSchema: customerUseCases}
	v1customer.RegisterCustomerServiceServer(grpcServer, &customerService)
	log.Print("Customer customerService registered")
}
