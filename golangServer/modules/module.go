package modules

import (
	"os/exec"
	"os"
	"log"
	"github.com/BurntSushi/toml"
)

var (
	errorMessage string
)

/////////////////////////
// Module Registration //
/////////////////////////
func GetModulesList() []IModule{

	var minimegaConfig MinimegaConfig
	var hostConfig HostConfig

	minimega := &Minimega{
		Module{
			Name: 		"minimega",
			ConfigPath: "",
			ConfigExtension: "minimega/",
			ConfigFile: "minimega.conf",
			Namespace:	"",
		},
		minimegaConfig,
	}

	opendss := &OpenDSS{
		Module{
			Name: 		"opendss",
			ConfigPath: "",
			ConfigExtension: "opendss/",
			ConfigFile: "opendss.conf",
			Namespace:	"",
		},
	}

	parser := &Parser{
		Module{
			Name: 		"parser",
			ConfigPath: "",
			ConfigExtension: "parser/",
			ConfigFile: "parser.conf",
			Namespace:	"",
		},
	}

	host := &Host{
		Module{
			Name: 		"host",
			ConfigPath: "",
			ConfigExtension: "host/",
			ConfigFile: "host.conf",
			Namespace:	"",
		},
		hostConfig,
	}

	var list []IModule = []IModule{ minimega, opendss, parser, host }

	return list
}

type Module struct {
	Name 		string
	ConfigPath 	string
	ConfigExtension	string
	ConfigFile	string
	Namespace 	string
}

type IModule interface {
	Configure(path, namespace string) bool
	customConfig() bool
	GetName() string
}

type IConfig interface {
	isValidConfig() bool
}

func runOSCommand(command, args string) bool {
	// cmd := exec.Command(command, args)
	cmd := exec.Command("echo", "[Running OS Command]: " + command + " " + args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		errorMessage = "Command Failed: " + command + " " + args
		log.Printf(errorMessage)
		return false
	}
	return true
}

func getOSCommandStdout(command, args string) (string, bool) {
	// out, err := exec.Command(command, args).Output()
    out, err := exec.Command("echo", "[OS Stdout Command]: " + command + " " + args).Output()
    if err != nil {
		errorMessage = "Command Failed: " + command + " " + args
		log.Printf(errorMessage)
		return "", false    
	}
	return string(out), true
}

func parseConfigurationFile(configfile string, config IConfig) bool {

	_, err := os.Stat(configfile)
	if err != nil {
		errorMessage := "Config File is Missing: " + configfile
		log.Printf(errorMessage)
		return false
	}

	if _, err := toml.DecodeFile(configfile, config); err != nil {
		errorMessage := "Config File Incorrectly Formatted: " + configfile
		log.Printf(errorMessage)
		return false
	}

	return true
}

func validConfig(config IConfig) bool {
	if config.isValidConfig() {
		return true
	}
	return false
}

func performCustomConfig(module IModule) bool {
	if module.customConfig() {
		return true
	}
	return false
}

func GetErrorMessage() string {
	return errorMessage
}
