package main

import (
	"log"
	"os"

	"github.com/ac2393921/hexagonal-architecture-with-go/internal/adapters/app/api"
	gRPC "github.com/ac2393921/hexagonal-architecture-with-go/internal/adapters/framework/left/grpc"
	"github.com/ac2393921/hexagonal-architecture-with-go/internal/adapters/framework/right/db"
	"github.com/ac2393921/hexagonal-architecture-with-go/internal/ports"
)

func main() {
	var err error

	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
