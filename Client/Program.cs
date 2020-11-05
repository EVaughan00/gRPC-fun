using Grpc.Core;
using System;
using Generalwrapper;
using System.Threading.Tasks;
using Client.Modules;

namespace Client
{
    class Program
    {
        const string target = "127.0.0.1:50050";
        // const string minimegaTarget = "127.0.0.1:50051";
        // const string parserTarget = "127.0.0.1:50052";
        const string localFilePath = "/home/evaughan/Documents/NREL/CEEP/Refactor/gRPC-fun/golangServer/Config/";
        const string demoFilePath = "/app/golangServer/Config/";
        const string demoNamespace = "testspace";
        static async Task Main(string[] args)
        {
            
            // Host
            Channel channel = new Channel(target, ChannelCredentials.Insecure);
            await channel.ConnectAsync().ContinueWith((task) =>
            {
                if (task.Status == TaskStatus.RanToCompletion)
                    Console.WriteLine("The client connected successfully");
            });

            var generalClient = new GeneralWrapperService.GeneralWrapperServiceClient(channel);

            var contextRequest = new ContextRequest {
                Message = "Requesting Module Context"
            };

            var context = generalClient.RequestModuleContext(contextRequest);

            var registry = new ModuleRegistry();
            registry.RegisterModules();

            foreach (Module module in registry.GetRegisteredModules()) {
                if (module.GetModuleContext().Reference == context.Reference) {
                    module.Activate(channel);
                }
            }
            channel.ShutdownAsync().Wait();

            Console.ReadKey();
        }

    }
}