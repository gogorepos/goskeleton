package client

import (
	"encoding/json"
	"net"
)

type Client interface {
	Address() string
	Command() []byte
}

func Send(client Client) ([]byte, error) {
	conn, err := net.Dial("tcp", client.Address())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if _, err = conn.Write(client.Command()); err != nil {
		return nil, err
	}
	var result []byte
	for {
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			return nil, err
		}
		result = append(result, buf[:n]...)
		if json.Valid(result) {
			break
		}
	}
	return result, nil
}
