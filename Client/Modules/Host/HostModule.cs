using System;
using Grpc.Core;
using Hostmodule;

namespace Client.Modules.Host
{
    public class HostModule : Module
    {
        private HostContext _context { get; set; }
        private HostConfiguration _configuration { get; set; }
        public HostModule() {
            _context = new HostContext();
            _configuration = new HostConfiguration(); // Should use dependency injection
        }

        public override void Activate(ChannelBase channel)
        {
            Console.WriteLine("Activated Host Module");

            var client = RegisterClient(channel);

            var configuration = _configuration.ConvertToProtobuf();

            var confirmation = client.Configure(configuration);

            if (confirmation.Completed) {
                Console.WriteLine("Host Module Configuration Successful");
            } else {
                Console.WriteLine("Host Module Configuration Unsuccessful: " + confirmation.ErrorMessage);
            }
        }

        public HostService.HostServiceClient RegisterClient(ChannelBase channel) {
            return new HostService.HostServiceClient(channel);
        }

        public override ModuleContext GetModuleContext() {
            return _context;
        }
    }
}