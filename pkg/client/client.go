package client

import (
	"context"
	"fmt"
	"github.com/amjadjibon/grpc_todo/pkg/api/customer"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GrpcClient()  {
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := customer.NewCustomerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Create(c, ctx)
}

func Create(c customer.CustomerServiceClient, ctx context.Context)  {
	t := time.Now().In(time.UTC)
	time, _ := ptypes.TimestampProto(t)
	// Call Create
	req1 := customer.CreateRequest{
		Customer: &customer.Customer{
			FirstName: "Amjad",
			LastName: "Jibon",
			PhoneNumber: "01676805267",
			Email: "amjad@gmail.com",
			Password: "1234",
			CreatedAt: time,
			UpdatedAt: time,
		},
	}
	fmt.Println(req1)
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Printf("Create failed: %v\n", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)
}
