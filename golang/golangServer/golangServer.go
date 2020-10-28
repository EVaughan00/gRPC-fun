package main

import (
	"context"
	"log"
	"net"
	"os"
	"strings"
	"google.golang.org/grpc"
)

import pb "grpc/golangServer/proto"
import mods "grpc/golangServer/modules"

const (
	serviceName = "Emulator Service"
)
var (
	port string = os.Getenv("GRPC_PORT")
	modules []string = strings.Split(os.Getenv("GRPC_MODULES"), ",")
)

type server struct {
	pb.UnimplementedServiceInitConfigurationServer
}

func (s *server) RequestStatus(ctx context.Context, in *pb.StatusRequest) (*pb.StatusReply, error) {
	log.Printf("Received Message: %v", in.GetMessage())
	return &pb.StatusReply{ServiceName: serviceName, IsReady: true}, nil
}

func (s *server) IngestConfiguration(ctx context.Context, in *pb.ConfigurationInfo) (*pb.IngestConfirmation, error) {
	log.Printf("Configuration File Path: %v", in.GetFilePath())
	log.Printf("Configuration File Name: %v", in.GetConfigurationFile())

	performModuleConfiguration()

	return &pb.IngestConfirmation{ ServiceName: serviceName, WasSuccessful: true}, nil
}

func performModuleConfiguration() {

	for _, module := range modules {

		switch module {
		case "minimega":
			log.Printf("Loading minimega module")
			mods.ConfigureMinimega()

		case "opendss": 
			log.Printf("Loading opendss module")
			mods.ConfigureOpenDSS()
		}


	}

}

func main() {

	lis, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceInitConfigurationServer(s, &server{})

	log.Printf("Modules Loaded: %v", modules)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
