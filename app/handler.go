package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	res "github.com/codecrafters-io/redis-starter-go/response"
	"github.com/codecrafters-io/redis-starter-go/symbols"
)

func handleConn(conn net.Conn) error {
	var buf = make([]byte, 1048)
	_, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("failed to read connection data: %w", err)
	}

	buffer := bytes.NewBuffer(buf)
	cmdCount := strings.Count(buffer.String(), "ping")
	for i := 0; i <= cmdCount; i++ {
		var writeBuffer bytes.Buffer
		writeBuffer.WriteString(symbols.Plus.String())
		writeBuffer.WriteString("PONG")
		writeBuffer.WriteString("\r\n")
		response := res.New(writeBuffer.Bytes())
		err = response.Write(conn)
		if err != nil {
			return err
		}
	}

	return nil
}
