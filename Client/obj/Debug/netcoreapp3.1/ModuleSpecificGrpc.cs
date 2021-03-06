// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: proto/Minimega/ModuleSpecific.proto
// </auto-generated>
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Modulespecific {
  public static partial class ModuleSpecficService
  {
    static readonly string __ServiceName = "modulespecific.ModuleSpecficService";

    static readonly grpc::Marshaller<global::Modulespecific.Configuration> __Marshaller_modulespecific_Configuration = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Modulespecific.Configuration.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Modulespecific.Confirmation> __Marshaller_modulespecific_Confirmation = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Modulespecific.Confirmation.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Modulespecific.StatusRequest> __Marshaller_modulespecific_StatusRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Modulespecific.StatusRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Modulespecific.DetailedEvaluation> __Marshaller_modulespecific_DetailedEvaluation = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Modulespecific.DetailedEvaluation.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Modulespecific.ConfigurationUpdate> __Marshaller_modulespecific_ConfigurationUpdate = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Modulespecific.ConfigurationUpdate.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Modulespecific.CustomCommand> __Marshaller_modulespecific_CustomCommand = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Modulespecific.CustomCommand.Parser.ParseFrom);

    static readonly grpc::Method<global::Modulespecific.Configuration, global::Modulespecific.Confirmation> __Method_Configure = new grpc::Method<global::Modulespecific.Configuration, global::Modulespecific.Confirmation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Configure",
        __Marshaller_modulespecific_Configuration,
        __Marshaller_modulespecific_Confirmation);

    static readonly grpc::Method<global::Modulespecific.StatusRequest, global::Modulespecific.DetailedEvaluation> __Method_RequestDetailedHealthCheck = new grpc::Method<global::Modulespecific.StatusRequest, global::Modulespecific.DetailedEvaluation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "RequestDetailedHealthCheck",
        __Marshaller_modulespecific_StatusRequest,
        __Marshaller_modulespecific_DetailedEvaluation);

    static readonly grpc::Method<global::Modulespecific.ConfigurationUpdate, global::Modulespecific.Confirmation> __Method_UpdateConfiguration = new grpc::Method<global::Modulespecific.ConfigurationUpdate, global::Modulespecific.Confirmation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UpdateConfiguration",
        __Marshaller_modulespecific_ConfigurationUpdate,
        __Marshaller_modulespecific_Confirmation);

    static readonly grpc::Method<global::Modulespecific.CustomCommand, global::Modulespecific.Confirmation> __Method_RunCustomCommand = new grpc::Method<global::Modulespecific.CustomCommand, global::Modulespecific.Confirmation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "RunCustomCommand",
        __Marshaller_modulespecific_CustomCommand,
        __Marshaller_modulespecific_Confirmation);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Modulespecific.ModuleSpecificReflection.Descriptor.Services[0]; }
    }

    /// <summary>Client for ModuleSpecficService</summary>
    public partial class ModuleSpecficServiceClient : grpc::ClientBase<ModuleSpecficServiceClient>
    {
      /// <summary>Creates a new client for ModuleSpecficService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public ModuleSpecficServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for ModuleSpecficService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public ModuleSpecficServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected ModuleSpecficServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected ModuleSpecficServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::Modulespecific.Confirmation Configure(global::Modulespecific.Configuration request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Configure(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Modulespecific.Confirmation Configure(global::Modulespecific.Configuration request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Configure, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.Confirmation> ConfigureAsync(global::Modulespecific.Configuration request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ConfigureAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.Confirmation> ConfigureAsync(global::Modulespecific.Configuration request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Configure, null, options, request);
      }
      public virtual global::Modulespecific.DetailedEvaluation RequestDetailedHealthCheck(global::Modulespecific.StatusRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RequestDetailedHealthCheck(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Modulespecific.DetailedEvaluation RequestDetailedHealthCheck(global::Modulespecific.StatusRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_RequestDetailedHealthCheck, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.DetailedEvaluation> RequestDetailedHealthCheckAsync(global::Modulespecific.StatusRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RequestDetailedHealthCheckAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.DetailedEvaluation> RequestDetailedHealthCheckAsync(global::Modulespecific.StatusRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_RequestDetailedHealthCheck, null, options, request);
      }
      public virtual global::Modulespecific.Confirmation UpdateConfiguration(global::Modulespecific.ConfigurationUpdate request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateConfiguration(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Modulespecific.Confirmation UpdateConfiguration(global::Modulespecific.ConfigurationUpdate request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UpdateConfiguration, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.Confirmation> UpdateConfigurationAsync(global::Modulespecific.ConfigurationUpdate request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateConfigurationAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.Confirmation> UpdateConfigurationAsync(global::Modulespecific.ConfigurationUpdate request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UpdateConfiguration, null, options, request);
      }
      public virtual global::Modulespecific.Confirmation RunCustomCommand(global::Modulespecific.CustomCommand request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RunCustomCommand(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Modulespecific.Confirmation RunCustomCommand(global::Modulespecific.CustomCommand request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_RunCustomCommand, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.Confirmation> RunCustomCommandAsync(global::Modulespecific.CustomCommand request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RunCustomCommandAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Modulespecific.Confirmation> RunCustomCommandAsync(global::Modulespecific.CustomCommand request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_RunCustomCommand, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override ModuleSpecficServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new ModuleSpecficServiceClient(configuration);
      }
    }

  }
}
#endregion
