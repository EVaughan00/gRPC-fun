using Grpc.Core;
using Helloworld;
using Statuscheck;
using System;
using System.Threading.Tasks;

namespace Client
{
    class Program
    {
        const string target = "127.0.0.1:50051";

        static async Task Main(string[] args)
        {
            Channel channel = new Channel(target, ChannelCredentials.Insecure);

            await channel.ConnectAsync().ContinueWith((task) =>
            {
                if (task.Status == TaskStatus.RanToCompletion)
                    Console.WriteLine("The client connected successfully");
            });

            var client = new ServiceInitConfiguration.ServiceInitConfigurationClient(channel);

            var configurationInfo = new ConfigurationInfo {
                FilePath = "/opt/minimega/images",
                ConfigurationFile = "EmulatorsInitConfiguration.txt"
            };

            var configurationResponse = client.IngestConfiguration(configurationInfo);

            if (configurationResponse.WasSuccessful) {
                Console.WriteLine("Service: " + configurationResponse.ServiceName + " successfully configured");
            } else {
                Console.WriteLine("Service: " + configurationResponse.ServiceName + " was not configured");
            }

            var statusRequest = new StatusRequest()
            {
                Message = "Requesting Service Status"
            };

            var statusResponse = client.RequestStatus(statusRequest);

            if (statusResponse.IsReady) {
                Console.WriteLine("Service: " + statusResponse.ServiceName + " is ready");
            } else {
                Console.WriteLine("Service: " + statusResponse.ServiceName + " is not ready");
            }

            channel.ShutdownAsync().Wait();
            Console.ReadKey();
        }
    }
}