package commands

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/pkg/grpcclient"
	"github.com/spf13/cobra"
	"log"
)

func NewStoreCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [key] [type] [data]",
		Short: "Store data",
		Args:  cobra.ExactArgs(3),
		Run:   handleStoreDataCommand,
	}
	cmd.Flags().StringP(flagToken, "T", "", "Authorization token")
	cmd.MarkFlagRequired(flagToken)
	return cmd
}

func handleStoreDataCommand(cmd *cobra.Command, args []string) {
	key := args[0]
	dataType := args[1]
	data := args[2]
	token, _ := cmd.Flags().GetString(flagToken)

	client, err := grpcclient.NewClient()
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	identifier, err := getIdentifier(dataType, data)
	if err != nil {
		log.Fatalf("could not extract identifier: %v", err)
	}

	resGet, err := client.GetData(ctx, key, dataType, identifier, token)
	if err == nil && len(resGet.Data) > 0 {
		log.Fatalf("Already exists")
	}

	if err := client.StoreData(ctx, key, dataType, data, token); err != nil {
		log.Fatalf("could not store data: %v", err)
	}
	log.Println("Data stored successfully")
}
