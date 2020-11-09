using Grpc.Core;
using System;
using Generalwrapper;
using System.Threading.Tasks;
using Client.Modules;

namespace Client
{
    class Program
    {
        private static ContextRequest _contextRequest = new ContextRequest {Message = "Requesting Module Context"};
        public static ModuleRegistry registry = new ModuleRegistry();

        static async Task Main(string[] args)
        {
            
            // Register Modules
            registry.RegisterModules();

            // Host Channel
            // var hostChannel = await GenerateNewChannelFromTarget("127.0.0.1:50050");
            // var hostClient = GenerateClientFromChannel(hostChannel);
            // var hostContext = hostClient.RequestModuleContext(_contextRequest);
            // FindAndActivateMatchingModuleByReference(hostContext.Reference, hostChannel);


            var minimegaChannel = await GenerateNewChannelFromTarget("127.0.0.1:50050");
            var minimegaClient = GenerateClientFromChannel(minimegaChannel);
            var minimegaContext = minimegaClient.RequestModuleContext(_contextRequest);
            FindAndActivateMatchingModuleByReference(minimegaContext.Reference, minimegaChannel);

            Console.ReadKey();
        }

        public static async Task<ChannelBase> GenerateNewChannelFromTarget(string target) {
            Channel channel = new Channel(target, ChannelCredentials.Insecure);
            await channel.ConnectAsync().ContinueWith((task) =>
            {
                if (task.Status == TaskStatus.RanToCompletion)
                    Console.WriteLine("The client connected successfully");
            });

            return channel;
        }   

        public static GeneralWrapperService.GeneralWrapperServiceClient GenerateClientFromChannel(ChannelBase channel) {
            return new GeneralWrapperService.GeneralWrapperServiceClient(channel);
        }

        public static void FindAndActivateMatchingModuleByReference(string reference, ChannelBase channel) {
            foreach (Module module in registry.GetRegisteredModules()) {
                if (module.GetModuleContext().Reference == reference) {
                    module.Activate(channel);
                }
            }
        }
    }
}