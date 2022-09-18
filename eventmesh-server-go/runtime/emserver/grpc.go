// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package emserver

import (
	"github.com/apache/incubator-eventmesh/eventmesh-server-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

// GRPCServer represents an instance of a eventmesh server.
type GRPCServer struct {
	// httpOption option for current http server
	grpcOption *config.GRPCOption

	// grpcServer grpc eventmesh server
	grpcServer *grpc.Server

	// lis server listen
	lis net.Listener
}

// NewGRPCServer returns a new GRPC server
func NewGRPCServer(opt *config.GRPCOption) (GracefulServer, error) {
	var (
		srv *grpc.Server
		err error
	)
	lis, err := net.Listen("tcp", opt.Port)
	if err != nil {
		return nil, err
	}
	if opt.TLSOption.EnableInsecure {
		creds, err := credentials.NewServerTLSFromFile(opt.Certfile, opt.Keyfile)
		if err != nil {
			return nil, err
		}
		srv = grpc.NewServer(grpc.Creds(creds))
	} else {
		srv = grpc.NewServer()
	}

	return &GRPCServer{
		grpcOption: opt,
		grpcServer: srv,
		lis:        lis,
	}, nil
}

func (g *GRPCServer) Serve() error {
	// TODO register eventmesh grpc server handler
	return g.grpcServer.Serve(g.lis)
}

func (g *GRPCServer) Stop() error {
	g.grpcServer.GracefulStop()
	return nil
}