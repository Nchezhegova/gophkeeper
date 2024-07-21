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

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete data",
		Run:   deleteData,
	}
	cmd.Flags().StringP("key", "k", "", "Key to delete")
	cmd.Flags().StringP("type", "t", "", "Type of data to delete")
	cmd.Flags().StringP("identifier", "i", "", "Identifier of the data to delete")
	cmd.Flags().StringP("token", "T", "", "Authorization token")
	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("token")
	return cmd
}

func deleteData(cmd *cobra.Command, args []string) {
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

	reqGet := &proto.GetDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
	}
	_, err = client.GetData(ctx, reqGet)
	if err != nil {
		log.Fatalf("%v", err)
	}
	reqDelete := &proto.DeleteDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
	}
	_, err = client.DeleteData(ctx, reqDelete)
	if err != nil {
		log.Fatalf("could not delete data: %v", err)
	}
	log.Println("Data deleted successfully")
}
