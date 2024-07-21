package commands

import (
	"context"
	"fmt"
	"github.com/Nchezhegova/gophkeeper/pkg/grpcclient"
	"github.com/spf13/cobra"
	"log"
)

func NewLoginCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "login [username] [password]",
		Short: "Login a user",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := grpcclient.NewClient()
			if err != nil {
				log.Fatalf("could not connect to server: %v", err)
			}
			defer client.Close()

			ctx := context.Background()

			resp, err := client.Login(ctx, args[0], args[1])
			if err != nil {
				log.Fatalf("could not login: %v", err)
			}
			fmt.Printf("User logged in successfully, token: %s\n", resp.Token)
		},
	}
}
