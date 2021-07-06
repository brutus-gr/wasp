package wasp

import (
	"bufio"
	"net"
	"strconv"
)

type Client struct {
	connection net.Conn
}

func (c *Client) write(buffer []byte) (int, error) {
	return c.connection.Write(buffer)
}

func (c *Client) ConnectTo(ip string, port int) error {
	connection, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))

	if err != nil {
		return err
	}

	c.connection = connection

	return nil
}

func (c *Client) Send(m []byte) error {

	writer := bufio.NewWriter(c.connection)
	_, err := writer.Write(append(m, 4))

	if err == nil {
		err = writer.Flush()
	}
	return err
}

func (c *Client) Listen(onMessage func([]byte)) error {
	reader := bufio.NewReader(c.connection)
	for {
		message, err := reader.ReadBytes(4)
		message = message[:len(message)-1]

		if err != nil {
			return err
		}

		go onMessage(message)
	}
}
