package commands

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/pkg/grpcclient"
	"github.com/spf13/cobra"
	"log"
)

func NewRegisterCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "register [username] [password]",
		Short: "Register a new user",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := grpcclient.NewClient()
			if err != nil {
				log.Fatalf("could not connect to server: %v", err)
			}
			defer client.Close()

			ctx := context.Background()

			_, err = client.Register(ctx, args[0], args[1])
			if err != nil {
				log.Fatalf("could not register: %v", err)
			}
			log.Println("User registered successfully")
		},
	}
}
