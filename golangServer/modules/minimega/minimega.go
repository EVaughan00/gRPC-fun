package minimega

import (
	"log"
	"context"
	"errors"
	"google.golang.org/grpc"
	"strings"
	"os"
	"os/exec"

	module 	"grpc/golangServer/modules"
	protobuf "grpc/golangServer/proto/minimega"
)

var (
	this MinimegaModule
)

type MinimegaModule struct {
	Context module.ModuleContext
	Configuration MinimegaConfig
	ProtobufServer MinimegaProtobufServer
	Active bool
}

type MinimegaConfig struct {
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

type MinimegaProtobufServer struct {
	protobuf.UnimplementedMinimegaServiceServer
}

func (minimega *MinimegaModule) RegisterAsGrpcService(server grpc.ServiceRegistrar) {
	this = *minimega
	protobuf.RegisterMinimegaServiceServer(server, &MinimegaProtobufServer{})
	minimega.SetAsActive()
}

func (minimega *MinimegaModule) GetContext() module.ModuleContext {
	return minimega.Context
}

func (minimega *MinimegaModule) SetAsActive() {
	minimega.Active = true
}

func (minimega *MinimegaModule) IsActive() bool {
	return minimega.Active
}

//////////////////
// gRPC Methods //
//////////////////
func (server *MinimegaProtobufServer) Configure(ctx context.Context, in *protobuf.Configuration) (*protobuf.Confirmation, error) {

	var errorMessage string
	var completed bool

	config, err := server.mapProtobufToConfig(in)

	if err != nil {
		errorMessage = err.Error()
		completed = false
		log.Printf("Error: %v", err)
	} else {
		this.Configuration = config

		log.Printf("Successfully Configured Minimega Module")
		completed = true
	}

	customConfig()

	return &protobuf.Confirmation{ Completed: completed, ErrorMessage: errorMessage }, nil
}

func (server *MinimegaProtobufServer) mapProtobufToConfig(in *protobuf.Configuration) (MinimegaConfig, error) {

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

	var config MinimegaConfig

	config.VlanSpecs = vlanSpecs
	config.TapSpecs = tapSpecs
	config.Orchestrations = orchestrations

	if (MinimegaConfig{}).VlanSpecs.ManagementVLAN == config.VlanSpecs.ManagementVLAN || (MinimegaConfig{}).TapSpecs == config.TapSpecs {
		return config, errors.New("Minimega Incorrectly Configured")
	}

	return config, nil
}

func   customConfig() bool {

	if (!buildTaps()) {
		return false
	}

	if (!buildOrchestrations()) {
		return false
	}

	if (!buildHil()) {
		return false
	}
	return true
}

////////////////////////////////////
// Perform Orchestration Commands //
////////////////////////////////////
func buildTaps() bool {
	
	var snifferVLANs []string
	for _, vlan := range this.Configuration.SnifferVLANs {
		snifferVLANs = append(snifferVLANs, vlan)
	}
	commaVLANs := strings.Join(snifferVLANs[:], ",")

	if (!runOSCommand("/opt/minimega/bin/minimega", "-e tap create 0 name " + this.Configuration.SnifferTapName)) {
		return false
	}
	if (!runOSCommand("ovs-vsctl", "-- --id=@p get port " + this.Configuration.SnifferTapName + " -- --id=@m create mirror name=m0 select-all=true select-vlan=" + commaVLANs + " output-port=@p -- set bridge mega_bridge mirrors=@m")) {
		return false
	}
	if (!runOSCommand("/opt/minimega/bin/minimega", "-e tap create " + "testspace//" + this.Configuration.ManagementVLAN + " ip " + this.Configuration.PublisherTapIP +" " + this.Configuration.PublisherTapName)) {
		return false
	}
	if (! runOSCommand("/opt/minimega/bin/minimega", "-e capture netflow bridge mega_bridge udp " + this.Configuration.NetflowTapIP + ":" + this.Configuration.NetflowTapPort)) {
		return false
	}

	return true
}

func buildOrchestrations() bool {

	if (!runOSCommand("/opt/minimega/bin/minimega", "-e read " + this.Configuration.Orchestrations.Location)) {
		return false
	}
	return true
}

func buildHil() bool {

	var hilPrefix string = "hil-port-"

	for _, vlan := range this.Configuration.HilVLANs {
		if (!runOSCommand("ovs-vsctl", "add-port mega_bridge " + hilPrefix + vlan)) {
			return false
		}
		if (!runOSCommand("ovs-vsctl", "set port " + hilPrefix + vlan + " tag=" + vlan)) {
			return false
		}
	}
	return true
}

func runOSCommand(command, args string) bool {

	cmd := exec.Command("echo", "[Running OS Command]" + command + " " + args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}


// //////////////////////////////////
// // Build Sniffer Netflow Mirror //
// // This Requires Dependencies	//
// //////////////////////////////////
// func buildSnifferNetflowMirror() bool {

// 	    // // Deploy the internal mirror that mirrors all sniffer-tap traffic onto an
// 		// // ovs-docker interface
// 		// // Depends on 1. ovs-docker spawned interface: sniffer-mirror-in
//         // 3. ovs-vsctl add-br emulators-mirror-bridge
//         // 4. ovs-vsctl -- --id=@m create mirror name=emulators-mirror -- add bridge emulators-mirror-bridge mirrors @m -- --id=@port get port sniffer-mirror-in -- set mirror emulators-mirror output-port=@port
//         // 5. ovs-vsctl set Mirror emulators-mirror select_src_port=`ovs-vsctl get port this.Configuration.SnifferTapName _uuid` select_dst_port=`ovs-vsctl get port this.Configuration.SnifferTapName _uuid`
// 	return true
// }