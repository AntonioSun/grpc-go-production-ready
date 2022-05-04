package main

import (
	grpcclient "github.com/AntonioSun/grpc-go-production-ready/client"
)

func main() {
	grpcclient.TimeoutLogExample()
	//grpcclient.TLSConnExample()
}
