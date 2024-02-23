package server

import (
	"bufio"
	"encoding/binary"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("Test server with valid message", func(t *testing.T) {
		server, client := net.Pipe()

		go handleRequest(server)

		writer := bufio.NewWriter(client)

		// Prepare the request
		msg := "Hello, World!\n"
		req := make([]byte, 11+len(msg))
		req[0] = 0
		binary.BigEndian.PutUint32(req[1:5], uint32(len(msg)+11))
		binary.BigEndian.PutUint16(req[5:7], 0)
		binary.BigEndian.PutUint16(req[7:9], 0)
		binary.BigEndian.PutUint16(req[9:11], 0)
		copy(req[11:], msg)

		// Send the request
		writer.Write(req)
		writer.Flush()

		// Read the response
		reader := bufio.NewReader(client)
		respType, _ := reader.ReadByte()
		if respType != 1 {
			t.Errorf("Expected response type 1, got %d", respType)
		}

		lenBuf := make([]byte, 4)
		reader.Read(lenBuf)
		respLen := binary.BigEndian.Uint32(lenBuf)
		if respLen != uint32(len(msg)+11) {
			t.Errorf("Expected response length %d, got %d", len(msg), respLen)
		}

		idBuf := make([]byte, 2)
		reader.Read(idBuf)
		respID := binary.BigEndian.Uint16(idBuf)
		if respID != 0 {
			t.Errorf("Expected response ID 0, got %d", respID)
		}

		globalBuf := make([]byte, 2)
		reader.Read(globalBuf)
		respGlobal := binary.BigEndian.Uint16(globalBuf)
		if respGlobal != 1 {
			t.Errorf("Expected response global 1, got %d", respGlobal)
		}

		clientBuf := make([]byte, 2)
		reader.Read(clientBuf)
		respClient := binary.BigEndian.Uint16(clientBuf)
		if respClient != 1 {
			t.Errorf("Expected response client 1, got %d", respClient)
		}

		msgBuf := make([]byte, respLen-11)
		reader.Read(msgBuf)
		respMsg := string(msgBuf)
		expectedMsg := "hELLO, wORLD!\n"
		if respMsg != expectedMsg {
			t.Errorf("Expected response message %q, got %q", expectedMsg, respMsg)
		}
	})
}
