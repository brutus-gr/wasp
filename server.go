package wasp

import (
	"bufio"
	"net"
	"strconv"
)

type Server struct {
}

func (s *Server) handleNewConnection(conn net.Conn, onMessage func([]byte, net.Conn)) error {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadBytes(4)

		if err != nil {
			conn.Close()
			return err
		}

		message = message[:len(message)-1]

		go onMessage(message, conn)
	}
}

func (s *Server) Send(conn net.Conn, m []byte) error {
	writer := bufio.NewWriter(conn)

	_, err := writer.Write(append(m, 4))
	if err == nil {
		err = writer.Flush()
	}
	return err
}

func (s *Server) Listen(ip string, port int, onMessage func([]byte, net.Conn)) error {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	// Close the listener when the application closes.
	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		// Handle new connection.
		err = s.handleNewConnection(conn, onMessage)
		if err != nil {
			return err
		}
	}
}
