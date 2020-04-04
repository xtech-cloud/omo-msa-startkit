package omo.msa.startkit;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.28.0)",
    comments = "Source: proto/startkit/echo.proto")
public final class EchoGrpc {

  private EchoGrpc() {}

  public static final String SERVICE_NAME = "omo.msa.startkit.Echo";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<omo.msa.startkit.EchoOuterClass.Request,
      omo.msa.startkit.EchoOuterClass.Response> getCallMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Call",
      requestType = omo.msa.startkit.EchoOuterClass.Request.class,
      responseType = omo.msa.startkit.EchoOuterClass.Response.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<omo.msa.startkit.EchoOuterClass.Request,
      omo.msa.startkit.EchoOuterClass.Response> getCallMethod() {
    io.grpc.MethodDescriptor<omo.msa.startkit.EchoOuterClass.Request, omo.msa.startkit.EchoOuterClass.Response> getCallMethod;
    if ((getCallMethod = EchoGrpc.getCallMethod) == null) {
      synchronized (EchoGrpc.class) {
        if ((getCallMethod = EchoGrpc.getCallMethod) == null) {
          EchoGrpc.getCallMethod = getCallMethod =
              io.grpc.MethodDescriptor.<omo.msa.startkit.EchoOuterClass.Request, omo.msa.startkit.EchoOuterClass.Response>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Call"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  omo.msa.startkit.EchoOuterClass.Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  omo.msa.startkit.EchoOuterClass.Response.getDefaultInstance()))
              .setSchemaDescriptor(new EchoMethodDescriptorSupplier("Call"))
              .build();
        }
      }
    }
    return getCallMethod;
  }

  private static volatile io.grpc.MethodDescriptor<omo.msa.startkit.EchoOuterClass.Ping,
      omo.msa.startkit.EchoOuterClass.Pong> getPingPongMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PingPong",
      requestType = omo.msa.startkit.EchoOuterClass.Ping.class,
      responseType = omo.msa.startkit.EchoOuterClass.Pong.class,
      methodType = io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
  public static io.grpc.MethodDescriptor<omo.msa.startkit.EchoOuterClass.Ping,
      omo.msa.startkit.EchoOuterClass.Pong> getPingPongMethod() {
    io.grpc.MethodDescriptor<omo.msa.startkit.EchoOuterClass.Ping, omo.msa.startkit.EchoOuterClass.Pong> getPingPongMethod;
    if ((getPingPongMethod = EchoGrpc.getPingPongMethod) == null) {
      synchronized (EchoGrpc.class) {
        if ((getPingPongMethod = EchoGrpc.getPingPongMethod) == null) {
          EchoGrpc.getPingPongMethod = getPingPongMethod =
              io.grpc.MethodDescriptor.<omo.msa.startkit.EchoOuterClass.Ping, omo.msa.startkit.EchoOuterClass.Pong>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PingPong"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  omo.msa.startkit.EchoOuterClass.Ping.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  omo.msa.startkit.EchoOuterClass.Pong.getDefaultInstance()))
              .setSchemaDescriptor(new EchoMethodDescriptorSupplier("PingPong"))
              .build();
        }
      }
    }
    return getPingPongMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static EchoStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EchoStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EchoStub>() {
        @java.lang.Override
        public EchoStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EchoStub(channel, callOptions);
        }
      };
    return EchoStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static EchoBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EchoBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EchoBlockingStub>() {
        @java.lang.Override
        public EchoBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EchoBlockingStub(channel, callOptions);
        }
      };
    return EchoBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static EchoFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EchoFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EchoFutureStub>() {
        @java.lang.Override
        public EchoFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EchoFutureStub(channel, callOptions);
        }
      };
    return EchoFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class EchoImplBase implements io.grpc.BindableService {

    /**
     */
    public void call(omo.msa.startkit.EchoOuterClass.Request request,
        io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Response> responseObserver) {
      asyncUnimplementedUnaryCall(getCallMethod(), responseObserver);
    }

    /**
     */
    public io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Ping> pingPong(
        io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Pong> responseObserver) {
      return asyncUnimplementedStreamingCall(getPingPongMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getCallMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                omo.msa.startkit.EchoOuterClass.Request,
                omo.msa.startkit.EchoOuterClass.Response>(
                  this, METHODID_CALL)))
          .addMethod(
            getPingPongMethod(),
            asyncBidiStreamingCall(
              new MethodHandlers<
                omo.msa.startkit.EchoOuterClass.Ping,
                omo.msa.startkit.EchoOuterClass.Pong>(
                  this, METHODID_PING_PONG)))
          .build();
    }
  }

  /**
   */
  public static final class EchoStub extends io.grpc.stub.AbstractAsyncStub<EchoStub> {
    private EchoStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EchoStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EchoStub(channel, callOptions);
    }

    /**
     */
    public void call(omo.msa.startkit.EchoOuterClass.Request request,
        io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Response> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCallMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Ping> pingPong(
        io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Pong> responseObserver) {
      return asyncBidiStreamingCall(
          getChannel().newCall(getPingPongMethod(), getCallOptions()), responseObserver);
    }
  }

  /**
   */
  public static final class EchoBlockingStub extends io.grpc.stub.AbstractBlockingStub<EchoBlockingStub> {
    private EchoBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EchoBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EchoBlockingStub(channel, callOptions);
    }

    /**
     */
    public omo.msa.startkit.EchoOuterClass.Response call(omo.msa.startkit.EchoOuterClass.Request request) {
      return blockingUnaryCall(
          getChannel(), getCallMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class EchoFutureStub extends io.grpc.stub.AbstractFutureStub<EchoFutureStub> {
    private EchoFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EchoFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EchoFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<omo.msa.startkit.EchoOuterClass.Response> call(
        omo.msa.startkit.EchoOuterClass.Request request) {
      return futureUnaryCall(
          getChannel().newCall(getCallMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CALL = 0;
  private static final int METHODID_PING_PONG = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final EchoImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(EchoImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CALL:
          serviceImpl.call((omo.msa.startkit.EchoOuterClass.Request) request,
              (io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Response>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_PING_PONG:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.pingPong(
              (io.grpc.stub.StreamObserver<omo.msa.startkit.EchoOuterClass.Pong>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class EchoBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    EchoBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return omo.msa.startkit.EchoOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Echo");
    }
  }

  private static final class EchoFileDescriptorSupplier
      extends EchoBaseDescriptorSupplier {
    EchoFileDescriptorSupplier() {}
  }

  private static final class EchoMethodDescriptorSupplier
      extends EchoBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    EchoMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (EchoGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new EchoFileDescriptorSupplier())
              .addMethod(getCallMethod())
              .addMethod(getPingPongMethod())
              .build();
        }
      }
    }
    return result;
  }
}
