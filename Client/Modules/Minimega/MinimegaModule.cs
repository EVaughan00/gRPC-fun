using System;
using Grpc.Core;
using Minimegamodule;


namespace Client.Modules.Minimega
{
    public class MinimegaModule : Module
    {
        private MinimegaContext _context { get; set; }

        private MinimegaConfiguration _configuration { get; set; }
        public MinimegaModule() {
            _context = new MinimegaContext();
            _configuration = new MinimegaConfiguration(); // Should use dependency injection
        }

        public override void Activate(ChannelBase channel)
        {
            Console.WriteLine("Activated Minimega Module");

            var client = RegisterClient(channel);

            VlanSpecs vlanSpecs = new VlanSpecs{
                    ManagementVLAN = _configuration.ManagementVLAN,
            };

            TapSpecs tapSpecs = new TapSpecs{
                    NetflowTapPort = _configuration.NetflowTapPort,
                    NetflowTapIP = _configuration.NetflowTapIP,
                    PowerTapPort = _configuration.PowerTapPort,
                    SnifferTapName = _configuration.SnifferTapName,
                    PublisherTapName = _configuration.PublisherTapName,
                    PublisherTapIP = _configuration.PublisherTapIP
            };

            Orchestrations orchestrations = new Orchestrations{
                Location = _configuration.Location
            };

            var configuration = new Configuration {
                VlanSpecs = vlanSpecs,
                TapSpecs = tapSpecs,
                Orchestrations = orchestrations
            };

            configuration.VlanSpecs.SnifferVLANs.Add(_configuration.SnifferVLANs);
            configuration.VlanSpecs.HilVLANs.Add(_configuration.HilVLANs);

            var confirmation = client.Configure(configuration);

            if (confirmation.Completed) {
                Console.WriteLine("Minimega Module Configuration Successful");
            } else {
                Console.WriteLine("Minimega Module Configuration Unsuccessful: " + confirmation.ErrorMessage);
            }
        }

        public MinimegaService.MinimegaServiceClient RegisterClient(ChannelBase channel) {
            return new MinimegaService.MinimegaServiceClient(channel);
        }

        public override ModuleContext GetModuleContext() {
            return _context;
        }
    }
}