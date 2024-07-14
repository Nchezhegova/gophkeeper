package commands

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func NewRegisterCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "register [username] [password]",
		Short: "Register a new user",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			creds, err := credentials.NewClientTLSFromFile("../cert/server-cert.pem", "")
			if err != nil {
				log.Fatalf("could not load TLS certificate: %v", err)
			}

			conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			client := proto.NewGophKeeperClient(conn)

			ctx := context.Background()

			_, err = client.Register(ctx, &proto.RegisterRequest{
				Username: args[0],
				Password: args[1],
			})
			if err != nil {
				log.Fatalf("could not register: %v", err)
			}
			log.Println("User registered successfully")
		},
	}
}
