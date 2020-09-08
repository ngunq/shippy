// shippy/shippy-cli-consignment/main.go
package main

import (
	"encoding/json"
	"github.com/micro/go-micro/v2"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/ngunq/shippy/shippy-service-consignment/proto/consignment"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("Did not connect: %v", err)
	//}
	//defer conn.Close()
	//client := pb.NewShippingServiceClient(conn)

	service := micro.NewService(micro.Name("shippy.service.consignment"))
	service.Init()
	client := pb.NewShippingService("shippy.service.consignment", service.Client())

	//Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	consignments, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range consignments.Consignments {
		log.Println(v)
	}
}