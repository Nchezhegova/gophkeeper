package grpcclient

import (
	"context"
	"github.com/Nchezhegova/gophkeeper/internal/interfaces/grpc/proto"
	"github.com/Nchezhegova/gophkeeper/internal/tlsconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	serverAddress = "localhost:50051"
)

type Client struct {
	conn   *grpc.ClientConn
	client proto.GophKeeperClient
}

func NewClient() (*Client, error) {
	creds, err := tlsconfig.LoadClientTLSCredentials()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	client := proto.NewGophKeeperClient(conn)
	return &Client{conn: conn, client: client}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Login(ctx context.Context, username, password string) (*proto.LoginResponse, error) {
	req := &proto.LoginRequest{
		Username: username,
		Password: password,
	}
	return c.client.Login(ctx, req)
}

func (c *Client) Register(ctx context.Context, username, password string) (*proto.RegisterResponse, error) {
	req := &proto.RegisterRequest{
		Username: username,
		Password: password,
	}
	return c.client.Register(ctx, req)
}

func (c *Client) DeleteData(ctx context.Context, key, dataType, identifier, token string) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	reqDelete := &proto.DeleteDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
	}
	_, err := c.client.DeleteData(ctx, reqDelete)
	return err
}

func (c *Client) GetData(ctx context.Context, key, dataType, identifier, token string) (*proto.GetDataResponse, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	req := &proto.GetDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
	}
	return c.client.GetData(ctx, req)
}

func (c *Client) StoreData(ctx context.Context, key, dataType, data, token string) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	req := &proto.StoreDataRequest{
		Key:  key,
		Type: dataType,
		Data: []byte(data),
	}
	_, err := c.client.StoreData(ctx, req)
	return err
}

func (c *Client) UpdateData(ctx context.Context, key, dataType, identifier, newData, token string) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	req := &proto.UpdateDataRequest{
		Key:        key,
		Type:       dataType,
		Identifier: identifier,
		NewData:    []byte(newData),
	}
	_, err := c.client.UpdateData(ctx, req)
	return err
}
