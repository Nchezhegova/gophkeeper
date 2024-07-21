package commands

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)

func NewStoreCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [key] [type] [data]",
		Short: "Store data",
		Args:  cobra.ExactArgs(3),
		Run:   storeData,
	}
	cmd.Flags().StringP("token", "T", "", "Authorization token")
	cmd.MarkFlagRequired("token")
	return cmd
}

func storeData(cmd *cobra.Command, args []string) {
	key := args[0]
	dataType := args[1]
	data := args[2]
	token, _ := cmd.Flags().GetString("token")

	creds, err := credentials.NewClientTLSFromFile("../cert/server-cert.pem", "")
	if err != nil {
		log.Fatalf("could not load TLS certificate: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewGophKeeperClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)
	identifier := getIdentifier(dataType, data)

	reqGet := &proto.GetDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
	}
	resGet, err := client.GetData(ctx, reqGet)
	if err == nil && len(resGet.Data) > 0 {
		log.Fatalf("Already exists")
	}
	req := &proto.StoreDataRequest{
		Key:  key,
		Type: dataType,
		Data: []byte(data),
	}
	_, err = client.StoreData(ctx, req)
	if err != nil {
		log.Fatalf("could not store data: %v", err)
	}
	log.Println("Data stored successfully")
}
