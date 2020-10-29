package modules

import (
	"log"
	"os"
	"strings"
	"github.com/BurntSushi/toml"
)

const (
	hilPrefix string = "hil-port-"
	pathExt string = "minimega/"
	fileName string = pathExt + "minimega.conf"
	orchestrationsExt string = pathExt + "orchestrations/"
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
	thisMinimega.Module.ConfigPath = path
	thisMinimega.Module.Namespace = namespace
	
	if (!parseConfigurationFile()) {
		return false
	}

	if (!validConfig(thisMinimega.MinimegaConfig)) {
		return false
	}

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

func parseConfigurationFile() bool {

	var configfile = thisMinimega.Module.ConfigPath + fileName

	_, err := os.Stat(configfile)
	if err != nil {
		errorMessage := "Config File is Missing: " + configfile
		log.Printf(errorMessage)
		SetErrorMessage(errorMessage)
		return false
	}

	var config MinimegaConfig
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		errorMessage := "Config File Incorrectly Formatted: " + configfile
		log.Printf(errorMessage)
		SetErrorMessage(errorMessage)
		return false
	}

	thisMinimega.MinimegaConfig = config
	return true
}

func validConfig(config MinimegaConfig) bool {

	if len(config.Orchestrations) > 0 {
		return true
	}
	return false
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

	if (!RunOSCommand("/opt/minimega/bin/minimega", "-e tap create 0 name " + thisMinimega.MinimegaConfig.SnifferTapName)) {
		return false
	}
	if (!RunOSCommand("ovs-vsctl", "-- --id=@p get port " + thisMinimega.MinimegaConfig.SnifferTapName + " -- --id=@m create mirror name=m0 select-all=true select-vlan=" + commaVLANs + " output-port=@p -- set bridge mega_bridge mirrors=@m")) {
		return false
	}
	if (!RunOSCommand("/opt/minimega/bin/minimega", "-e tap create " + thisMinimega.Module.Namespace + "//" + thisMinimega.MinimegaConfig.ManagementVLAN[0] + " ip " + thisMinimega.MinimegaConfig.PublisherTapIP +" " + thisMinimega.MinimegaConfig.PublisherTapName)) {
		return false
	}
	if (! RunOSCommand("/opt/minimega/bin/minimega", "-e capture netflow bridge mega_bridge udp " + thisMinimega.MinimegaConfig.NetflowTapIP + ":" + thisMinimega.MinimegaConfig.NetflowTapPort)) {
		return false
	}

	return true
}

func buildOrchestrations() bool {

	for _, orchestration := range thisMinimega.MinimegaConfig.Orchestrations {
		if (!RunOSCommand("/opt/minimega/bin/minimega", "-e read " + thisMinimega.Module.ConfigPath + orchestrationsExt + orchestration)) {
			return false
		}
	}
	return true
}

func buildHil() bool {

	for _, vlan := range thisMinimega.MinimegaConfig.HilVLANs {
		if (!RunOSCommand("ovs-vsctl", "add-port mega_bridge " + hilPrefix + vlan)) {
			return false
		}
		if (!RunOSCommand("ovs-vsctl", "set port " + hilPrefix + vlan + " tag=" + vlan)) {
			return false
		}
	}
	return true
}