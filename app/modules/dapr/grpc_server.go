package dapr

import (
	"context"
	"fmt"
	pb "goke/proto"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

type stateServer struct {
	pb.UnimplementedStateStoreServer
}

func (s *stateServer) Store(ctx context.Context, state *pb.State) (*pb.Confirm, error) {
	log.Println("Function: Store()")
	if err := daprClient.SaveState(ctx, STORE_NAME, state.GetKey(), []byte(state.GetValue()), nil); err != nil {
		return &pb.Confirm{
			Ok: false,
		}, err
	}
	return &pb.Confirm{
		Ok: true,
	}, nil
}

func (s *stateServer) Retrieve(ctx context.Context, state *pb.State) (*pb.State, error) {
	log.Println("Function: Retrieve()")
	value, err := daprClient.GetState(ctx, STORE_NAME, state.GetKey(), nil); 
	if err != nil {
		return nil, err
	}
	return &pb.State{
		Key: state.Key,
		Value: string(value.Value),
	}, nil
}

func newStateServer() *stateServer {
	s := &stateServer{}
	return s
}

func runStateServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterStateStoreServer(grpcServer, newStateServer())
	grpcServer.Serve(lis)
}

func callGRPC() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:" + os.Getenv("DAPR_GRPC_PORT")), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("error!")
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewStateStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", os.Getenv("APP_ID"))

	ok, err := client.Store(ctx, &pb.State{Key: "grpc", Value: "hello world"})

	if !ok.GetOk() {
		log.Fatal("Something went wrong." + err.Error())
	}
	state, err := client.Retrieve(ctx, &pb.State{Key: "grpc"})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(state.GetValue())
}