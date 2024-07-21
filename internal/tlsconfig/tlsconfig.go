package tlsconfig

import (
	"google.golang.org/grpc/credentials"
	"log"
)

// Путь к сертификатам
const (
	clientCertFile = "../cert/server-cert.pem"
	serverCertFile = "../cert/server-cert.pem"
	serverKeyFile  = "../cert/server-key.pem"
)

func LoadServerTLSCredentials() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewServerTLSFromFile(serverCertFile, serverKeyFile)
	if err != nil {
		log.Fatalf("failed to load server TLS credentials: %v", err)
	}
	return creds, err
}

func LoadClientTLSCredentials() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewClientTLSFromFile(clientCertFile, "")
	if err != nil {
		log.Fatalf("failed to load client TLS credentials: %v", err)
	}
	return creds, err
}
