package modules

import (
	"log"
)

type Host struct {
	Module
	HostConfig
}

var (
	thisHost *Host
)

type HostConfig struct {
	HilBridge		string
	MirrorBridge	string
	SnifferMirrorIn	string
	SnifferMirrorOut 	string
}

func (host *Host) GetName() string {
	return host.Module.Name
}

func ConfigureHost(path string) bool {
	log.Printf("Running Host Module")
	return true
}

func (host *Host) Configure(path, namespace string) bool {
	log.Printf("Running Host Module")

	thisHost = host
	thisHost.Module.ConfigPath = path + thisHost.Module.ConfigExtension
	thisHost.Module.Namespace = namespace
	
	var configFile string = thisHost.Module.ConfigPath + thisHost.Module.ConfigFile

	parseSuccessful := parseConfigurationFile(configFile, IConfig(&thisHost.HostConfig))

	if !parseSuccessful {
		return false
	}

	var config IConfig = IConfig(&thisHost.HostConfig)

	if (!validConfig(config)) {
		return false
	}

	if (!performCustomConfig(thisHost)) {
		return false
	}

	return true
}

func (hostConfig *HostConfig) isValidConfig() bool {
	return true
}

func  (host *Host) customConfig() bool {

	if (!buildBridgesAndTaps()) {
		return false
	}

	if (!buildMirrors()) {
		return false
	}

	return true
}

func buildBridgesAndTaps() bool { 

	if !runOSCommand("ovs-vsctl", "add-br " + thisHost.HostConfig.HilBridge) {
		return false
	}

	if !runOSCommand("ovs-vsctl", "add-br " + thisHost.HostConfig.MirrorBridge) {
		return false
	}

	emulatorsDeployment, emulatorRetrieved := getOSCommandStdout("docker ps", "| grep -o 'k8s_POD_emulators." + thisHost.Module.Namespace + "*'")
	idsDeployment, idsRetrieved := getOSCommandStdout("docker ps", "| grep -o 'k8s_POD_ids." + thisHost.Module.Namespace + "*'")
	if !emulatorRetrieved || !idsRetrieved {
		return false
	}

	if !runOSCommand("ovs-docker", "add-port " + thisHost.HostConfig.MirrorBridge + " " + thisHost.HostConfig.SnifferMirrorIn + " " + emulatorsDeployment) {
		return false
	}
	if !runOSCommand("ovs-docker", "add-port " + thisHost.HostConfig.MirrorBridge + " " + thisHost.HostConfig.SnifferMirrorOut + " " + idsDeployment) {
		return false
	}
	return true
}

func buildMirrors() bool { 
	
	// SNIFFER_MIRROR_OUT = (ovs-docker)sniffer-mirror-out
	// SNIFFER_MIRROR_IN = (ovs-docker)sniffer-mirror-in

	// ovs-vsctl -- --id=@m create mirror name=hardware-mirror -- add bridge hardware-mirror-bridge mirrors @m -- --id=@port get port $SNIFFER_MIRROR_OUT -- set mirror emulators-mirror output-port=@port
	// ovs-vsctl set Mirror hardware-mirror select_src_port=`ovs-vsctl get port $SNIFFER_MIRROR_IN _uuid` select_dst_port=`ovs-vsctl get port $SNIFFER_MIRROR_IN _uuid`
	return true

}
