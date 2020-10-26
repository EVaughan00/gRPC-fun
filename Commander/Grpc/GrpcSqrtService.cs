using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Grpc.Core;
using Sqrt;
using static Sqrt.SqrtService;

namespace Commander.GrpcServices
{
    public class GrpcSqrtService : SqrtServiceBase
    {
        public override async Task<SqrtReponse> sqrt(SqrtRequest request, ServerCallContext context)
        {
            await Task.CompletedTask;

            int number = request.Number;

            if (number >= 0)
                return new SqrtReponse() { SquareRoot = Math.Sqrt(number) };
            else
                throw new RpcException(new Status(StatusCode.InvalidArgument, "number < 0"));
        }
    }
}