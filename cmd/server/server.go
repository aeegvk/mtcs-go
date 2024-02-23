package server

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

var globalCounter uint16 = 0

func main() {
	ln, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Server failed to listen: ", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Server failed to accept the connection", err)
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	var clientCounter uint16 = 0
	reader := bufio.NewReader(conn)

	for {
		// Read the request
		reqType, _ := reader.ReadByte()
		if reqType != 0 {
			fmt.Println("Expected request type 0, got ", reqType)
			return
		}

		lenBuf := make([]byte, 4)
		reader.Read(lenBuf)
		reqLen := binary.BigEndian.Uint32(lenBuf)

		idBuf := make([]byte, 2)
		reader.Read(idBuf)
		reqID := binary.BigEndian.Uint16(idBuf)

		reader.Discard(4) // Discard the global and client counters

		msgBuf := make([]byte, reqLen-11)
		reader.Read(msgBuf)
		reqMsg := string(msgBuf)

		resMsg := strings.Map(func(r rune) rune {
			switch {
			case 'a' <= r && r <= 'z':
				return r - 'a' + 'A'
			case 'A' <= r && r <= 'Z':
				return r - 'A' + 'a'
			default:
				return r
			}
		}, reqMsg)

		// Update counters
		globalCounter++
		clientCounter++

		// Prepare the response
		resp := make([]byte, 11+len(resMsg))
		resp[0] = 1
		binary.BigEndian.PutUint32(resp[1:5], uint32(len(resMsg)+11))
		binary.BigEndian.PutUint16(resp[5:7], reqID)
		binary.BigEndian.PutUint16(resp[7:9], globalCounter)
		binary.BigEndian.PutUint16(resp[9:11], clientCounter)
		copy(resp[11:], resMsg)

		// Send the response
		conn.Write(resp)
	}
}
