package modules

import (
	// "os/exec"
	// "os"
	// "log"
	// "github.com/BurntSushi/toml"
	// "google.golang.org/grpc"
)

type ModuleRegistry struct {
	modules []IModule
	activeModules []IModule
}

/////////////////////////
// Module Registration //
/////////////////////////
func (registry *ModuleRegistry) RegisterModules() {

	var minimegaConfig MinimegaConfig
	var minimegaProtobufServer MinimegaProtobufServer

	minimega := &MinimegaModule{ 
		ModuleContext {
			Reference: "minimega",
		},
		minimegaConfig,
		minimegaProtobufServer,
		false,
	}

	registry.modules = append(registry.modules, minimega)
}

func (registry *ModuleRegistry) GetRegisteredModules() []IModule {
	return registry.modules
}

func (registry *ModuleRegistry) GetActiveModules() []IModule {

	for _, module := range registry.modules {
		if module.isActive() {
			registry.activeModules = append(registry.activeModules, module)
		}
	}
	
	return registry.activeModules
}