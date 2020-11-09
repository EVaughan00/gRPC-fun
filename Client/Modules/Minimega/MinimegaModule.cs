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

            var configuration = _configuration.ConvertToProtobuf();

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