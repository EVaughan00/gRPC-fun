package modules

import (
	"log"
	"strings"
)

const (
	orchestrationsExt string = "orchestrations/"
)

var (
	thisMinimega *Minimega
)

type Minimega struct {
	Module
	MinimegaConfig
}

type MinimegaConfig struct {
	ManagementVLAN	[]string
	SnifferVLANs	map[string]string
	HilVLANs		map[string]string
	NetflowTapPort	string
	NetflowTapIP	string
	PowerTapPort	string
	SnifferTapName	string
	PublisherTapName	string
	PublisherTapIP	string
	Orchestrations	[]string	
}

func (minimega *Minimega) GetName() string {
	return minimega.Module.Name
}

///////////////////////
// Configure Minmega //
///////////////////////
func ConfigureMinimega(path, namespace string) bool {
	log.Printf("Configuring Minimega")
	return true
}


func (minimega *Minimega) Configure(path, namespace string) bool {
	log.Printf("Configuring Minimega")

	thisMinimega = minimega
	thisMinimega.Module.ConfigPath = path + thisMinimega.Module.ConfigExtension
	thisMinimega.Module.Namespace = namespace
	
	var configFile string = thisMinimega.Module.ConfigPath + thisMinimega.Module.ConfigFile

	parseSuccessful := parseConfigurationFile(configFile, IConfig(&thisMinimega.MinimegaConfig))

	if !parseSuccessful {
		return false
	}

	var config IConfig = IConfig(&thisMinimega.MinimegaConfig)

	if (!validConfig(config)) {
		return false
	}

	if (!performCustomConfig(thisMinimega)) {
		return false
	}

	return true
}

func (minimegaConfig *MinimegaConfig) isValidConfig() bool {
	if len(minimegaConfig.Orchestrations) > 0 {
		return true
	}
	return false
}

func  (minimega *Minimega) customConfig() bool {

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
	for _, vlan := range thisMinimega.MinimegaConfig.SnifferVLANs {
		snifferVLANs = append(snifferVLANs, vlan)
	}
	commaVLANs := strings.Join(snifferVLANs[:], ",")

	if (!runOSCommand("/opt/minimega/bin/minimega", "-e tap create 0 name " + thisMinimega.MinimegaConfig.SnifferTapName)) {
		return false
	}
	if (!runOSCommand("ovs-vsctl", "-- --id=@p get port " + thisMinimega.MinimegaConfig.SnifferTapName + " -- --id=@m create mirror name=m0 select-all=true select-vlan=" + commaVLANs + " output-port=@p -- set bridge mega_bridge mirrors=@m")) {
		return false
	}
	if (!runOSCommand("/opt/minimega/bin/minimega", "-e tap create " + thisMinimega.Module.Namespace + "//" + thisMinimega.MinimegaConfig.ManagementVLAN[0] + " ip " + thisMinimega.MinimegaConfig.PublisherTapIP +" " + thisMinimega.MinimegaConfig.PublisherTapName)) {
		return false
	}
	if (! runOSCommand("/opt/minimega/bin/minimega", "-e capture netflow bridge mega_bridge udp " + thisMinimega.MinimegaConfig.NetflowTapIP + ":" + thisMinimega.MinimegaConfig.NetflowTapPort)) {
		return false
	}

	return true
}

func buildOrchestrations() bool {

	for _, orchestration := range thisMinimega.MinimegaConfig.Orchestrations {
		if (!runOSCommand("/opt/minimega/bin/minimega", "-e read " + thisMinimega.Module.ConfigPath + orchestrationsExt + orchestration)) {
			return false
		}
	}
	return true
}

func buildHil() bool {

	var hilPrefix string = "hil-port-"

	for _, vlan := range thisMinimega.MinimegaConfig.HilVLANs {
		if (!runOSCommand("ovs-vsctl", "add-port mega_bridge " + hilPrefix + vlan)) {
			return false
		}
		if (!runOSCommand("ovs-vsctl", "set port " + hilPrefix + vlan + " tag=" + vlan)) {
			return false
		}
	}
	return true
}

//////////////////////////////////
// Build Sniffer Netflow Mirror //
// This Requires Dependencies	//
//////////////////////////////////
func buildSnifferNetflowMirror() bool {

	    // // Deploy the internal mirror that mirrors all sniffer-tap traffic onto an
		// // ovs-docker interface
		// // Depends on 1. ovs-docker spawned interface: sniffer-mirror-in
        // 3. ovs-vsctl add-br emulators-mirror-bridge
        // 4. ovs-vsctl -- --id=@m create mirror name=emulators-mirror -- add bridge emulators-mirror-bridge mirrors @m -- --id=@port get port sniffer-mirror-in -- set mirror emulators-mirror output-port=@port
        // 5. ovs-vsctl set Mirror emulators-mirror select_src_port=`ovs-vsctl get port thisMinimega.MinimegaConfig.SnifferTapName _uuid` select_dst_port=`ovs-vsctl get port thisMinimega.MinimegaConfig.SnifferTapName _uuid`
	return true
}