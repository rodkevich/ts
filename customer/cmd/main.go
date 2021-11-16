package main

import (
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/rodkevich/ts/customer/config"

	transport "github.com/rodkevich/ts/customer/internal/transport/grpc"
	"github.com/rodkevich/ts/customer/pkg/postgres"
	v1 "github.com/rodkevich/ts/proto/customer/v1"
	serverGRPC "github.com/rodkevich/ts/server/grpc"
)

func main() {

	// configs:
	configPath := config.GetConfigPath(os.Getenv("config"))
	conf, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("unable to parse cfg path %v", err)

	}

	// pg:
	pgxConn, err := postgres.NewPgxConn(conf)
	if err != nil {
		log.Fatal("", err.Error())
	}
	defer pgxConn.Close()
	// serverGRPC:
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	builder := serverGRPC.GrpcServerBuilder{}
	builder.EnableReflection(true)

	s := builder.Build()
	s.RegisterService(customerServiceRegister)

	err = s.Start("0.0.0.0:50051")

	if err != nil {
		log.Fatalf("%v", err)
	}
	s.AwaitTermination(func() {
		log.Print("Shutting down the serverGRPC")
	})
}

func customerServiceRegister(sv *grpc.Server) {
	server := transport.CustomerService{}
	v1.RegisterCustomerServiceServer(sv, &server)
}
