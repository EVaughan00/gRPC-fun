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
import modules "grpc/golangServer/modules"

/////////////////////////
// Global Declarations //
/////////////////////////
var (
	port string = os.Getenv("GRPC_PORT")
	serviceName string = os.Getenv("GRPC_SERVICE_NAME")
	selectedModules []string = strings.Split(os.Getenv("GRPC_MODULES"), ",")
	successfulConfiguration bool
)

type server struct {
	pb.UnimplementedServiceInitConfigurationServer
}

//////////////////
// gRPC Methods //
//////////////////
func (s *server) RequestStatus(ctx context.Context, in *pb.StatusRequest) (*pb.StatusReply, error) {
	log.Printf("Sending Status Update")
	return &pb.StatusReply{ServiceName: serviceName, IsReady: successfulConfiguration}, nil
}

func (s *server) ConfigureAllModules(ctx context.Context, in *pb.ConfigurationInfo) (*pb.IngestConfirmation, error) {
	log.Printf("Configuring All Modules")
	log.Printf("Configuration File Path: %v", in.GetFilePath())
	log.Printf("Configuration Namespace: %v", in.GetNamespace())

	var wasSuccessful bool
	if len(selectedModules) > 0 {
		wasSuccessful = routeModuleConfiguration(in.GetFilePath(), in.GetNamespace())
	}
	successfulConfiguration = wasSuccessful

	errorMessage := ""

	if (!successfulConfiguration) {
		errorMessage = modules.GetErrorMessage()
	}

	return &pb.IngestConfirmation{ ServiceName: serviceName, WasSuccessful: wasSuccessful, ErrorMessage: errorMessage}, nil
}

//////////////////////////////////
// Module Configuration Routing //
//////////////////////////////////
func routeModuleConfiguration(path, namespace string) bool {

	var wasSuccessful bool

	var modulesList = modules.GetModulesList()

	for _, module := range modulesList {
		for _, selectedMod := range selectedModules {
			if module.GetName() == selectedMod {
				wasSuccessful = module.Configure(path, namespace)
			} else {
				wasSuccessful = true
			}		
			if !wasSuccessful {
				break
			}
		}
		if !wasSuccessful {
			break
		}
	}
	return wasSuccessful
}

//////////
// Main //
//////////
func main() {

	lis, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceInitConfigurationServer(s, &server{})

	log.Printf("Modules Loaded: %v", selectedModules)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
