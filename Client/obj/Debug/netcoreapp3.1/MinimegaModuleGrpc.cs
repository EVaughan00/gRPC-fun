// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: proto/Minimega/MinimegaModule.proto
// </auto-generated>
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Minimegamodule {
  public static partial class MinimegaService
  {
    static readonly string __ServiceName = "minimegamodule.MinimegaService";

    static readonly grpc::Marshaller<global::Minimegamodule.Configuration> __Marshaller_minimegamodule_Configuration = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Minimegamodule.Configuration.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Minimegamodule.Confirmation> __Marshaller_minimegamodule_Confirmation = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Minimegamodule.Confirmation.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Minimegamodule.StatusRequest> __Marshaller_minimegamodule_StatusRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Minimegamodule.StatusRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Minimegamodule.DetailedEvaluation> __Marshaller_minimegamodule_DetailedEvaluation = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Minimegamodule.DetailedEvaluation.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Minimegamodule.ConfigurationUpdate> __Marshaller_minimegamodule_ConfigurationUpdate = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Minimegamodule.ConfigurationUpdate.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Minimegamodule.CustomCommand> __Marshaller_minimegamodule_CustomCommand = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Minimegamodule.CustomCommand.Parser.ParseFrom);

    static readonly grpc::Method<global::Minimegamodule.Configuration, global::Minimegamodule.Confirmation> __Method_Configure = new grpc::Method<global::Minimegamodule.Configuration, global::Minimegamodule.Confirmation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Configure",
        __Marshaller_minimegamodule_Configuration,
        __Marshaller_minimegamodule_Confirmation);

    static readonly grpc::Method<global::Minimegamodule.StatusRequest, global::Minimegamodule.DetailedEvaluation> __Method_RequestDetailedHealthCheck = new grpc::Method<global::Minimegamodule.StatusRequest, global::Minimegamodule.DetailedEvaluation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "RequestDetailedHealthCheck",
        __Marshaller_minimegamodule_StatusRequest,
        __Marshaller_minimegamodule_DetailedEvaluation);

    static readonly grpc::Method<global::Minimegamodule.ConfigurationUpdate, global::Minimegamodule.Confirmation> __Method_UpdateConfiguration = new grpc::Method<global::Minimegamodule.ConfigurationUpdate, global::Minimegamodule.Confirmation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UpdateConfiguration",
        __Marshaller_minimegamodule_ConfigurationUpdate,
        __Marshaller_minimegamodule_Confirmation);

    static readonly grpc::Method<global::Minimegamodule.CustomCommand, global::Minimegamodule.Confirmation> __Method_RunCustomCommand = new grpc::Method<global::Minimegamodule.CustomCommand, global::Minimegamodule.Confirmation>(
        grpc::MethodType.Unary,
        __ServiceName,
        "RunCustomCommand",
        __Marshaller_minimegamodule_CustomCommand,
        __Marshaller_minimegamodule_Confirmation);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Minimegamodule.MinimegaModuleReflection.Descriptor.Services[0]; }
    }

    /// <summary>Client for MinimegaService</summary>
    public partial class MinimegaServiceClient : grpc::ClientBase<MinimegaServiceClient>
    {
      /// <summary>Creates a new client for MinimegaService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public MinimegaServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for MinimegaService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public MinimegaServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected MinimegaServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected MinimegaServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::Minimegamodule.Confirmation Configure(global::Minimegamodule.Configuration request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Configure(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Minimegamodule.Confirmation Configure(global::Minimegamodule.Configuration request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Configure, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.Confirmation> ConfigureAsync(global::Minimegamodule.Configuration request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ConfigureAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.Confirmation> ConfigureAsync(global::Minimegamodule.Configuration request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Configure, null, options, request);
      }
      public virtual global::Minimegamodule.DetailedEvaluation RequestDetailedHealthCheck(global::Minimegamodule.StatusRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RequestDetailedHealthCheck(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Minimegamodule.DetailedEvaluation RequestDetailedHealthCheck(global::Minimegamodule.StatusRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_RequestDetailedHealthCheck, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.DetailedEvaluation> RequestDetailedHealthCheckAsync(global::Minimegamodule.StatusRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RequestDetailedHealthCheckAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.DetailedEvaluation> RequestDetailedHealthCheckAsync(global::Minimegamodule.StatusRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_RequestDetailedHealthCheck, null, options, request);
      }
      public virtual global::Minimegamodule.Confirmation UpdateConfiguration(global::Minimegamodule.ConfigurationUpdate request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateConfiguration(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Minimegamodule.Confirmation UpdateConfiguration(global::Minimegamodule.ConfigurationUpdate request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UpdateConfiguration, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.Confirmation> UpdateConfigurationAsync(global::Minimegamodule.ConfigurationUpdate request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateConfigurationAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.Confirmation> UpdateConfigurationAsync(global::Minimegamodule.ConfigurationUpdate request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UpdateConfiguration, null, options, request);
      }
      public virtual global::Minimegamodule.Confirmation RunCustomCommand(global::Minimegamodule.CustomCommand request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RunCustomCommand(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Minimegamodule.Confirmation RunCustomCommand(global::Minimegamodule.CustomCommand request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_RunCustomCommand, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.Confirmation> RunCustomCommandAsync(global::Minimegamodule.CustomCommand request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RunCustomCommandAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Minimegamodule.Confirmation> RunCustomCommandAsync(global::Minimegamodule.CustomCommand request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_RunCustomCommand, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override MinimegaServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new MinimegaServiceClient(configuration);
      }
    }

  }
}
#endregion
