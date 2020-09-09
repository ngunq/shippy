// shippy/shippy-service-consignment/main.go

package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	pb "github.com/ngunq/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/ngunq/shippy/shippy-service-vessel/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consigmentCollection := client.Database("shippy").Collection("consignment")

	repository := &MongoRepository{collection: consigmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.client", service.Client())
	h := &handler{repository, vesselClient}

	// Register service
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
