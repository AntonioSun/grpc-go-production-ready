package main

import (
	grpcserver "github.com/AntonioSun/grpc-go-production-ready/server"
)

func main() {
	grpcserver.ServerInitialization()
	//grpcserver.ServerInitializationWithTLS()
}
