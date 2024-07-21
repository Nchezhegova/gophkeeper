package commands

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/pkg/grpcclient"
	"github.com/spf13/cobra"
	"log"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [flags]",
		Short: "Get data",
		Run:   handleGetDataCommand,
	}
	cmd.Flags().StringP(flagKey, "k", "", "Key to retrieve")
	cmd.Flags().StringP(flagType, "t", "", "Type of data to retrieve")
	cmd.Flags().StringP(flagIdentifier, "i", "", "Identifier of the data to retrieve")
	cmd.Flags().StringP(flagToken, "T", "", "Authorization token")
	cmd.MarkFlagRequired(flagToken)
	return cmd
}

func handleGetDataCommand(cmd *cobra.Command, args []string) {
	key, _ := cmd.Flags().GetString(flagKey)
	dataType, _ := cmd.Flags().GetString(flagType)
	identifier, _ := cmd.Flags().GetString(flagIdentifier)
	token, _ := cmd.Flags().GetString(flagToken)

	client, err := grpcclient.NewClient()
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	res, err := client.GetData(ctx, key, dataType, identifier, token)
	if err != nil {
		log.Fatalf("could not get data: %v", err)
	}

	for _, d := range res.Data {
		log.Printf("Data ID: %d, Type: %s, Key: %s, Data: %s", d.Id, d.Type, d.Key, string(d.Data))
	}
}
