package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/wwwstephen/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(r context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	id := uuid.New()
	return &invoicer.CreateResponse{
		Id:       id.String(),
		Sender:   req.GetSender(),
		Receiver: req.GetReceiver(),
		Amount:   req.GetAmount(),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatal("cannot create listener: %s", err)
	}

	ServiceRegistrar := grpc.NewServer()
	service := &myInvoiceServer{}

	invoicer.RegisterInvoicerServer(ServiceRegistrar, service)
	err = ServiceRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}

}
