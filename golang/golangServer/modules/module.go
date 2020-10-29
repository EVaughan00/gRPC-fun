package modules

import (
	"os/exec"
	"os"
	"log"
)

var (
	errorMessage string
)

func GetModulesList() []IModule{

	var minimegaConfig MinimegaConfig

	minimega := &Minimega{
		Module{
			Name: 		"minimega",
			ConfigPath: "",
			Namespace:	"",
		},
		minimegaConfig,
	}

	opendss := &OpenDSS{
		Module{
			Name: 		"opendss",
			ConfigPath: "",
			Namespace:	"",
		},
	}

	parser := &Parser{
		Module{
			Name: 		"parser",
			ConfigPath: "",
			Namespace:	"",
		},
	}

	var list []IModule = []IModule{ minimega, opendss, parser }

	return list
}

type Module struct {
	Name 		string
	ConfigPath 	string
	Namespace 	string
}

type IModule interface {
	Configure(path, namespace string) bool
	GetName() string
}

func RunOSCommand(command, args string) bool {

	// cmd := exec.Command(command, args)
	cmd := exec.Command("echo", "Running Example OS Command: " + command + " " + args)
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

func SetErrorMessage(err string) {
	errorMessage = err
}

func GetErrorMessage() string {
	return errorMessage
}
