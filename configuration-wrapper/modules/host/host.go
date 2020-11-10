package host

import (
	"log"
	"context"
	"errors"
	"google.golang.org/grpc"
	module "grpc/wrapper/modules"

	protobuf "grpc/wrapper/proto/host"
)

var (
	this HostModule
)

type HostModule struct {
	Context module.ModuleContext
	Configuration HostConfig
	ProtobufServer HostProtobufServer
	Active bool
}

type HostConfig struct {
	VlanSpecs
	TapSpecs
	Orchestrations	
}

type VlanSpecs struct {
	ManagementVLAN	string
	SnifferVLANs	map[string]string
	HilVLANs		map[string]string
}

type TapSpecs struct {
	NetflowTapPort	string
	NetflowTapIP	string
	PowerTapPort	string
	SnifferTapName	string
	PublisherTapName	string
	PublisherTapIP	string
}

type Orchestrations struct {
	Location string
}

type HostProtobufServer struct {
	protobuf.UnimplementedHostServiceServer
}

func (host *HostModule) RegisterAsGrpcService(server grpc.ServiceRegistrar) {
	this = *host
	protobuf.RegisterHostServiceServer(server, &HostProtobufServer{})
	host.SetAsActive()
}

func (host *HostModule) GetContext() module.ModuleContext {
	return host.Context
}

func (host *HostModule) SetAsActive() {
	host.Active = true
}

func (host *HostModule) IsActive() bool {
	return host.Active
}

//////////////////
// gRPC Methods //
//////////////////
func (server *HostProtobufServer) Configure(ctx context.Context, in *protobuf.Configuration) (*protobuf.Confirmation, error) {

	var errorMessage string
	var completed bool

	config, err := server.mapProtobufToConfig(in)

	if err != nil {
		errorMessage = err.Error()
		completed = false
		log.Printf("Error: %v", err)
	} else {
		this.Configuration = config

		log.Printf("Successfully Configured Host Module")
		completed = true
	}

	return &protobuf.Confirmation{ Completed: completed, ErrorMessage: errorMessage }, nil
}

func (server *HostProtobufServer) mapProtobufToConfig(in *protobuf.Configuration) (HostConfig, error) {

	vlanSpecs := VlanSpecs{
		in.VlanSpecs.ManagementVLAN,
		in.VlanSpecs.SnifferVLANs,
		in.VlanSpecs.HilVLANs,
	}

	tapSpecs := TapSpecs{
		in.TapSpecs.NetflowTapPort,
		in.TapSpecs.NetflowTapIP,
		in.TapSpecs.PowerTapPort,
		in.TapSpecs.SnifferTapName,
		in.TapSpecs.PublisherTapName,
		in.TapSpecs.PublisherTapIP,
	}

	orchestrations := Orchestrations{
		in.Orchestrations.Location,
	}

	var config HostConfig

	config.VlanSpecs = vlanSpecs
	config.TapSpecs = tapSpecs
	config.Orchestrations = orchestrations

	if (HostConfig{}).VlanSpecs.ManagementVLAN == config.VlanSpecs.ManagementVLAN || (HostConfig{}).TapSpecs == config.TapSpecs {
		return config, errors.New("Host Incorrectly Configured")
	}

	return config, nil
}
// func buildBridgesAndTaps() bool { 

// 	if !runOSCommand("ovs-vsctl", "add-br " + thisHost.HostConfig.HilBridge) {
// 		return false
// 	}

// 	if !runOSCommand("ovs-vsctl", "add-br " + thisHost.HostConfig.MirrorBridge) {
// 		return false
// 	}

// 	emulatorsDeployment, emulatorRetrieved := getOSCommandStdout("docker ps", "| grep -o 'k8s_POD_emulators." + thisHost.Module.Namespace + "*'")
// 	idsDeployment, idsRetrieved := getOSCommandStdout("docker ps", "| grep -o 'k8s_POD_ids." + thisHost.Module.Namespace + "*'")
// 	if !emulatorRetrieved || !idsRetrieved {
// 		return false
// 	}

// 	if !runOSCommand("ovs-docker", "add-port " + thisHost.HostConfig.MirrorBridge + " " + thisHost.HostConfig.SnifferMirrorIn + " " + emulatorsDeployment) {
// 		return false
// 	}
// 	if !runOSCommand("ovs-docker", "add-port " + thisHost.HostConfig.MirrorBridge + " " + thisHost.HostConfig.SnifferMirrorOut + " " + idsDeployment) {
// 		return false
// 	}
// 	return true
// }

// func buildMirrors() bool { 
	
// 	// SNIFFER_MIRROR_OUT = (ovs-docker)sniffer-mirror-out
// 	// SNIFFER_MIRROR_IN = (ovs-docker)sniffer-mirror-in

// 	// ovs-vsctl -- --id=@m create mirror name=hardware-mirror -- add bridge hardware-mirror-bridge mirrors @m -- --id=@port get port $SNIFFER_MIRROR_OUT -- set mirror emulators-mirror output-port=@port
// 	// ovs-vsctl set Mirror hardware-mirror select_src_port=`ovs-vsctl get port $SNIFFER_MIRROR_IN _uuid` select_dst_port=`ovs-vsctl get port $SNIFFER_MIRROR_IN _uuid`
// 	return true

// }