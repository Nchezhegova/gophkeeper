package commands

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/pkg/grpcclient"
	"github.com/spf13/cobra"
	"log"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update data",
		Run:   handleUpdateDataCommand,
	}
	cmd.Flags().StringP(flagKey, "k", "", "Key to update")
	cmd.Flags().StringP(flagType, "t", "", "Type of data to update")
	cmd.Flags().StringP(flagIdentifier, "i", "", "Identifier of the data to update")
	cmd.Flags().StringP(flagNewData, "d", "", "New data in JSON format")
	cmd.Flags().StringP(flagToken, "T", "", "Authorization token")
	cmd.MarkFlagRequired(flagKey)
	cmd.MarkFlagRequired(flagType)
	cmd.MarkFlagRequired(flagNewData)
	cmd.MarkFlagRequired(flagToken)
	return cmd
}

func handleUpdateDataCommand(cmd *cobra.Command, args []string) {
	key, _ := cmd.Flags().GetString(flagKey)
	dataType, _ := cmd.Flags().GetString(flagType)
	identifier, _ := cmd.Flags().GetString(flagIdentifier)
	newData, _ := cmd.Flags().GetString(flagNewData)
	token, _ := cmd.Flags().GetString(flagToken)

	client, err := grpcclient.NewClient()
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	newIdentifier, err := getIdentifier(dataType, newData)
	if err != nil {
		log.Fatalf("could not extract identifier: %v", err)
	}

	if _, err := client.GetData(ctx, key, dataType, identifier, token); err != nil {
		log.Fatalf("%v", err)
	}

	if newIdentifier != identifier {
		resNew, err := client.GetData(ctx, key, dataType, newIdentifier, token)
		if err == nil && len(resNew.Data) > 0 {
			log.Fatalf("Already exists")
		}
	}

	if err := client.UpdateData(ctx, key, dataType, identifier, newData, token); err != nil {
		log.Fatalf("could not update data: %v", err)
	}
	log.Println("Data updated successfully")
}
