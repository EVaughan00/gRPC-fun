using Grpc.Core;
using Helloworld;
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

            var client = new Greeter.GreeterClient(channel);

            var request = new HelloRequest()
            {
                Name = "Clement"
            };


            var response = client.SayHello(request);

            Console.WriteLine(response.Message);

            channel.ShutdownAsync().Wait();
            Console.ReadKey();
        }
    }
}