package grpc_server

import (
	"github.com/AntonioSun/grpc-go-production-ready/grpcutils"
	"github.com/AntonioSun/grpc-go-production-ready/tlscert"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildGrpcServer(t *testing.T) {
	builder := &GrpcServerBuilder{}
	builder.SetTlsCert(&tlscert.Cert)
	builder.DisableDefaultHealthCheck(true)
	builder.EnableReflection(true)
	builder.SetStreamInterceptors(grpcutils.GetDefaultStreamServerInterceptors())
	builder.SetUnaryInterceptors(grpcutils.GetDefaultUnaryServerInterceptors())
	server := builder.Build()
	assert.NotNil(t, server)
}
