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

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update data",
		Run:   updateData,
	}
	cmd.Flags().StringP("key", "k", "", "Key to update")
	cmd.Flags().StringP("type", "t", "", "Type of data to update")
	cmd.Flags().StringP("identifier", "i", "", "Identifier of the data to update")
	cmd.Flags().StringP("new_data", "d", "", "New data in JSON format")
	cmd.Flags().StringP("token", "T", "", "Authorization token")
	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("new_data")
	cmd.MarkFlagRequired("token")
	return cmd
}

func updateData(cmd *cobra.Command, args []string) {
	key, _ := cmd.Flags().GetString("key")
	dataType, _ := cmd.Flags().GetString("type")
	identifier, _ := cmd.Flags().GetString("identifier")
	newData, _ := cmd.Flags().GetString("new_data")
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

	newIdentifier := getIdentifier(dataType, newData)

	reqGet := &proto.GetDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
	}
	_, err = client.GetData(ctx, reqGet)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if newIdentifier != identifier {
		reqGetNew := &proto.GetDataRequest{
			Key:        key,
			Type:       dataType,
			Identifier: newIdentifier,
		}
		resNew, err := client.GetData(ctx, reqGetNew)
		if err == nil && len(resNew.Data) > 0 {
			log.Fatalf("Already exists")
		}
	}

	reqUpdate := &proto.UpdateDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
		NewData:    []byte(newData),
	}
	_, err = client.UpdateData(ctx, reqUpdate)
	if err != nil {
		log.Fatalf("could not update data: %v", err)
	}
	log.Println("Data updated successfully")
}
