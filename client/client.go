package client

import (
	"context"
	"net"

	"github.com/anthdm/ggcache/proto"
)

type Options struct{}

type Client struct {
	conn net.Conn
}

func New(endpoint string, opts Options) (*Client, error) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Set(ctx context.Context, key []byte, value []byte, ttl int) (any, error) {
	cmd := &proto.CommandSet{
		Key:   key,
		Value: value,
		TTL:   ttl,
	}

	_, err := c.conn.Write(cmd.Bytes())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
