
using System;
using Grpc.Core;

namespace Client.Modules
{
    public abstract class Module
    {
        public abstract ModuleContext GetModuleContext();
        public abstract void Activate(ChannelBase channel);
    }
    public abstract class ModuleContext{
        public string Reference { get; protected set; }
    }
    public interface IConfiguration<T>{
         T ConvertToProtobuf();
    }
}