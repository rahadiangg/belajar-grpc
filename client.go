package main

import (
	"context"
	"os"

	"github.com/rahadiangg/belajar-grpc/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	// var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50050", grpc.WithInsecure())
	if err != nil {
		log.Error("Can't start GRPC: ", err.Error())
	}

	defer conn.Close()

	client := model.NewProductServiceClient(conn)

	var page int64 = 2
	request := &model.Page{
		Page: &page,
	}

	response, err := client.GetProducts(context.Background(), request)
	if err != nil {
		log.Error("Err when GetProducts: ", err.Error())

	}

	log.Printf("Response from server: %v\n", response)
}
