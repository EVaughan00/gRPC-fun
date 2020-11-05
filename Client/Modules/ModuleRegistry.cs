using System;
using System.Collections.Generic;
using Client.Modules.Minimega;

namespace Client.Modules
{
    public class ModuleRegistry
    {
        private List<Module> _registeredModules { get; set; }

        public void RegisterModules() {
            
            _registeredModules = new List<Module>();

            _registeredModules.Add(
                new MinimegaModule()
            );
        }

        internal List<Module> GetRegisteredModules()
        {
            return _registeredModules;
        }
    }
}