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

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [flags]",
		Short: "Get data",
		Run:   getData,
	}
	cmd.Flags().StringP("key", "k", "", "Key to retrieve")
	cmd.Flags().StringP("type", "t", "", "Type of data to retrieve")
	cmd.Flags().StringP("identifier", "i", "", "Identifier of the data to retrieve")
	cmd.Flags().StringP("token", "T", "", "Authorization token")
	cmd.MarkFlagRequired("token")
	return cmd
}

func getData(cmd *cobra.Command, args []string) {
	key, _ := cmd.Flags().GetString("key")
	dataType, _ := cmd.Flags().GetString("type")
	identifier, _ := cmd.Flags().GetString("identifier")
	token, _ := cmd.Flags().GetString("token")

	creds, err := credentials.NewClientTLSFromFile("../cert/server-cert.pem", "")
	if err != nil {
		log.Fatalf("could not load TLS certificate: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewGophKeeperClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	req := &proto.GetDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
	}
	res, err := client.GetData(ctx, req)
	if err != nil {
		log.Fatalf("could not get data: %v", err)
	}

	for _, d := range res.Data {
		log.Printf("Data ID: %d, Type: %s, Key: %s, Data: %s", d.Id, d.Type, d.Key, string(d.Data))
	}
}
