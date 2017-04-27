package hermes

import (
	"errors"

	zmq "github.com/pebbe/zmq4"
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	ID     string
	s      *zmq.Socket
	cfg    *Config
	closed bool
}

func NewClient(c *Config) *Client {
	return &Client{
		ID:     uuid.NewV4().String(),
		cfg:    c,
		closed: false,
	}
}

func (c *Client) Init() (chan *Message, error) {
	socket, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		return nil, err
	}
	c.s = socket
	c.s.SetIdentity(c.ID)

	if err = c.s.Connect(c.cfg.ServerEndpoint); err != nil {
		return nil, err
	}

	msgs := make(chan *Message)

	go func() {
		for !c.closed {
			msg, err := c.s.RecvMessage(0)
			if err != nil {
				continue
			}

			m, err := ParseMessage(msg)
			if err == nil {
				msgs <- m
			}
		}
	}()

	return msgs, nil
}

func (c *Client) Close() error {
	if c.s != nil {
		c.closed = true

		if err := c.s.Close(); err != nil {
			return err
		}

	}
	return nil
}

func (c *Client) SendMessage(msg *Message) error {
	if c.s == nil {
		return errors.New("Socket is not connected. Please call Connect() first.")
	}
	if msg == nil {
		return errors.New("Message cannot be nil.")
	}

	if _, err := c.s.SendMessage(msg.ToStringArray()); err != nil {
		return err
	}

	return nil
}
