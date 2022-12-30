package main

import (
	"net"
	"os"

	"github.com/rahadiangg/belajar-grpc/config"
	"github.com/rahadiangg/belajar-grpc/model"
	"github.com/rahadiangg/belajar-grpc/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const port = ":50050"

func init() {
	log.SetOutput(os.Stdout)

}

func main() {

	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Err can't listen net: ", err)
	}

	grpcServer := grpc.NewServer()

	db := config.ConnectDb()

	productService := service.ProductService{DB: db}
	model.RegisterProductServiceServer(grpcServer, &productService)

	log.Println("Server started at ", netListen.Addr())

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatal("Failed to serve GRPC: ", err)
	}
}
