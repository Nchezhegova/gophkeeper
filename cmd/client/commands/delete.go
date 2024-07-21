package commands

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/pkg/grpcclient"
	"github.com/spf13/cobra"
	"log"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete data",
		Run:   handleDeleteDataCommand,
	}
	cmd.Flags().StringP(flagKey, "k", "", "Key to delete")
	cmd.Flags().StringP(flagType, "t", "", "Type of data to delete")
	cmd.Flags().StringP(flagIdentifier, "i", "", "Identifier of the data to delete")
	cmd.Flags().StringP(flagToken, "T", "", "Authorization token")
	cmd.MarkFlagRequired(flagKey)
	cmd.MarkFlagRequired(flagType)
	cmd.MarkFlagRequired(flagToken)
	return cmd
}

func handleDeleteDataCommand(cmd *cobra.Command, args []string) {
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

	if _, err := client.GetData(ctx, key, dataType, identifier, token); err != nil {
		log.Fatalf("could not get data: %v", err)
	}

	if err := client.DeleteData(ctx, key, dataType, identifier, token); err != nil {
		log.Fatalf("could not delete data: %v", err)
	}
	log.Println("Data deleted successfully")
}
