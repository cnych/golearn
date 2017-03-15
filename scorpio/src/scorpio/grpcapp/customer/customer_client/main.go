package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	pb "scorpio/grpcapp/customer"
	"stardust/logx"
)

const (
	address = "localhost:50051"
)

func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		logx.Fatalf("Could not create Customer: %v", err)
	} else {
		if resp.Success {
			logx.Printf("A new Customer has been added with id: %d", resp.Id)
		}
	}
}

func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		logx.Fatalf("Error on get customers: %v", err)
	} else {
		for {
			// Receiving the stream of data
			customer, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				logx.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
			}
			logx.Printf("Customer: %v", customer)
		}
	}
}

func main() {
	// Setup a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logx.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.CustomerRequest_Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)

	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)

	// Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: ""}
	getCustomers(client, filter)
}
