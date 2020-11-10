package main

import (
	"context"
	"log"
	"net"
	"os"
	"strings"
	"errors"
	"google.golang.org/grpc"
)

import protobuf "grpc/wrapper/proto"
import moduleInfrastruture "grpc/wrapper/modules/infrastructure"

/////////////////////////
// Global Declarations //
/////////////////////////
var (
	port string = os.Getenv("GRPC_PORT")
	selectedModules []string = strings.Split(os.Getenv("GRPC_MODULES"), ",")
	moduleRegistry = moduleInfrastruture.ModuleRegistry{}
	errorMessage string
)

type gernericServer struct {
	protobuf.UnimplementedGeneralWrapperServiceServer
}

////////////////////
// Error Handling //
////////////////////
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

func New(text string) error {
    return &errorString{text}
}


//////////////////
// gRPC Methods //
//////////////////
func (s *gernericServer) RequestModuleContext(ctx context.Context, in *protobuf.ContextRequest) (*protobuf.ModuleContext, error) {

	var reference string

	for _, module := range moduleRegistry.GetActiveModules() {
		log.Printf("Sending Loaded Module Context: %v", module.GetContext().Reference)
		reference = module.GetContext().Reference
	}

	return &protobuf.ModuleContext{ Reference: reference }, nil
}

//////////
// Main //
//////////
func main() {

	lis, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	server := grpc.NewServer()

	// Register general protobuf with the gRPC Server
	protobuf.RegisterGeneralWrapperServiceServer(server, &gernericServer{})

	if err := RegisterModulesAndBindServices(server); err != nil {
		log.Fatalf("Failed to Bind Module Services: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}

/////////////////////////
// Module Registration //
/////////////////////////
func RegisterModulesAndBindServices(server grpc.ServiceRegistrar) error {

	var noMatchingModulesFound bool = true

	log.Printf("Modules Loaded: %v", selectedModules)

	moduleRegistry.RegisterModules()
	registeredModulesList := moduleRegistry.GetRegisteredModules()

	if len(registeredModulesList) < 1 {
		return errors.New("Could Not Find Any Modules to Register")
	}

	for _, module := range registeredModulesList {

		for _, selectedModule := range selectedModules {

			if module.GetContext().Reference == selectedModule {
				noMatchingModulesFound = false

				log.Printf("Registering Module: %v", module.GetContext().Reference)

				module.RegisterAsGrpcService(server)
			}
		}
	}

	if noMatchingModulesFound {
		return errors.New("Could Find Matching Modules From Selected Modules")
	}

	return nil
}
