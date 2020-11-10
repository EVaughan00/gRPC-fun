package infrastructure

import (
	minimega "grpc/wrapper/modules/minimega"
	host "grpc/wrapper/modules/host"
	module "grpc/wrapper/modules"
)

type ModuleRegistry struct {
	modules []module.IModule
	activeModules []module.IModule
}

/////////////////////////
// Module Registration //
/////////////////////////
func (registry *ModuleRegistry) RegisterModules() {

	var minimegaConfig minimega.MinimegaConfig
	var minimegaProtobufServer minimega.MinimegaProtobufServer

	minimegaModule := &minimega.MinimegaModule{ 
		module.ModuleContext {
			Reference: "minimega",
		},
		minimegaConfig,
		minimegaProtobufServer,
		false,
	}

	registry.modules = append(registry.modules, minimegaModule)

	var hostConfig host.HostConfig
	var hostProtobufServer host.HostProtobufServer

	hostModule := &host.HostModule{ 
		module.ModuleContext {
			Reference: "host",
		},
		hostConfig,
		hostProtobufServer,
		false,
	}

	registry.modules = append(registry.modules, hostModule)


}

func (registry *ModuleRegistry) GetRegisteredModules() []module.IModule {
	return registry.modules
}

func (registry *ModuleRegistry) GetActiveModules() []module.IModule {

	var moduleIsUnique bool

	for _, currentModule := range registry.modules {

		moduleIsUnique = true

		if currentModule.IsActive() {
			
			for _, active := range registry.activeModules {
				if currentModule == active {
					moduleIsUnique = false
				}
			}

			if moduleIsUnique {
				registry.activeModules = append(registry.activeModules, currentModule)
			}
		}
	}

	return registry.activeModules
}