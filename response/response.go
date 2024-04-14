package response

import (
	"fmt"
	"net"
)

type Response struct {
	data []byte
}

func New(data []byte) *Response {
	return &Response{
		data: data,
	}
}

func (r *Response) Write(conn net.Conn) error {
	_, err := conn.Write(r.data)
	if err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}
