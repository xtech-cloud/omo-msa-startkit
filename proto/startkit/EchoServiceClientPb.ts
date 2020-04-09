/**
 * @fileoverview gRPC-Web generated client stub for omo.msa.startkit
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  Ping,
  Pong,
  Request,
  Response} from './echo_pb';

export class EchoClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCall = new grpcWeb.AbstractClientBase.MethodInfo(
    Response,
    (request: Request) => {
      return request.serializeBinary();
    },
    Response.deserializeBinary
  );

  call(
    request: Request,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Response) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/omo.msa.startkit.Echo/Call',
      request,
      metadata || {},
      this.methodInfoCall,
      callback);
  }

}

