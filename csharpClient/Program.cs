using Grpc.Core;
using Helloworld;
using Statuscheck;
using System;
using System.Threading.Tasks;

namespace Client
{
    class Program
    {
        const string hostTarget = "127.0.0.1:50050";
        const string minimegaTarget = "127.0.0.1:50051";
        const string parserTarget = "127.0.0.1:50052";
        const string localFilePath = "/home/evaughan/Documents/NREL/CEEP/Refactor/gRPC-fun/golang/golangServer/Config/";
        const string demoFilePath = "/app/golangServer/Config/";
        const string demoNamespace = "testspace";
        static async Task Main(string[] args)
        {
            bool hostSuccess = false;
            bool emulatorsSuccess = false;

            // Host
            Channel hostChannel = new Channel(hostTarget, ChannelCredentials.Insecure);
            await hostChannel.ConnectAsync().ContinueWith((task) =>
            {
                if (task.Status == TaskStatus.RanToCompletion)
                    Console.WriteLine("The client connected successfully");
            });
            var hostClient = new ServiceInitConfiguration.ServiceInitConfigurationClient(hostChannel);

            // GRPC Configuration Message
            var configurationInfo = new ConfigurationInfo {
                FilePath = demoFilePath,
                Namespace = demoNamespace
            };

            var hostConfigurationResponse = hostClient.ConfigureAllModules(configurationInfo);
            if (hostConfigurationResponse.WasSuccessful) {
                hostSuccess = true;

                Console.WriteLine("Service: " + hostConfigurationResponse.ServiceName + " successfully configured");
                Console.WriteLine("Advancing to Next Service and Shutting Down Channel");

            } else {
                Console.WriteLine("Service: " + hostConfigurationResponse.ServiceName + " was not configured");
                Console.WriteLine("Error: " + hostConfigurationResponse.ErrorMessage);
            }

            hostChannel.ShutdownAsync().Wait();

            if (hostSuccess) {

                // Minimega
                Channel minimegaChannel = new Channel(minimegaTarget, ChannelCredentials.Insecure);
                await minimegaChannel.ConnectAsync().ContinueWith((task) =>
                {
                    if (task.Status == TaskStatus.RanToCompletion)
                        Console.WriteLine("The client connected successfully");
                });
                var minimegaClient = new ServiceInitConfiguration.ServiceInitConfigurationClient(minimegaChannel);

                var minimegaConfigurationResponse = minimegaClient.ConfigureAllModules(configurationInfo);

                if (minimegaConfigurationResponse.WasSuccessful) {
                    emulatorsSuccess = true;
                    
                    Console.WriteLine("Service: " + minimegaConfigurationResponse.ServiceName + " successfully configured");
                    Console.WriteLine("Advancing to Next Service and Shutting Down Channel");

                } else {
                    Console.WriteLine("Service: " + minimegaConfigurationResponse.ServiceName + " was not configured");
                    Console.WriteLine("Error: " + minimegaConfigurationResponse.ErrorMessage);
                }
                minimegaChannel.ShutdownAsync().Wait();
            }

            
            if (hostSuccess && emulatorsSuccess) {

                // Parser
                Channel parserChannel = new Channel(parserTarget, ChannelCredentials.Insecure);
                await parserChannel.ConnectAsync().ContinueWith((task) =>
                {
                    if (task.Status == TaskStatus.RanToCompletion)
                        Console.WriteLine("The client connected successfully");
                });
                var parserClient = new ServiceInitConfiguration.ServiceInitConfigurationClient(parserChannel);

                var parserConfigurationResponse = parserClient.ConfigureAllModules(configurationInfo);

                if (parserConfigurationResponse.WasSuccessful) {

                    Console.WriteLine("Service: " + parserConfigurationResponse.ServiceName + " successfully configured");
                    Console.WriteLine("Advancing to Next Service and Shutting Down Channel");

                } else {
                    Console.WriteLine("Service: " + parserConfigurationResponse.ServiceName + " was not configured");
                    Console.WriteLine("Error: " + parserConfigurationResponse.ErrorMessage);
                }
                parserChannel.ShutdownAsync().Wait();
            }

            // var statusRequest = new StatusRequest()
            // {
            //     Message = "Requesting Service Status"
            // };

            // var statusResponse = client.RequestStatus(statusRequest);

            // if (statusResponse.IsReady) {
            //     Console.WriteLine("Service: " + statusResponse.ServiceName + " is ready");
            // } else {
            //     Console.WriteLine("Service: " + statusResponse.ServiceName + " is not ready");
            // }

            Console.ReadKey();
        }
    }
}