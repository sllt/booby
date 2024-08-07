package quic

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/quic-go/quic-go"
)

// Listener wraps quick.Listener to net.Listener
type Listener struct {
	*quic.Listener
}

// Accept waits for and returns the next connection to the listener.
func (ln *Listener) Accept() (net.Conn, error) {
	session, err := ln.Listener.Accept(context.Background())
	if err != nil {
		return nil, err
	}

	stream, err := session.AcceptStream(context.Background())
	if err != nil {
		return nil, err
	}

	return &Conn{session, stream}, err
}

// Conn wraps quick.Session to net.Conn
type Conn struct {
	quic.Connection
	quic.Stream
}

// ListenAddr Listen wraps quic listen
func ListenAddr(addr string, config *tls.Config) (*Listener, error) {
	ln, err := quic.ListenAddr(addr, config, nil)
	if err != nil {
		return nil, err
	}
	return &Listener{ln}, err
}

// Listen wraps quic listen
func Listen(conn net.PacketConn, tlsConf *tls.Config, quicConf *quic.Config) (*Listener, error) {
	ln, err := quic.Listen(conn, tlsConf, quicConf)
	if err != nil {
		return nil, err
	}
	return &Listener{ln}, err
}

// Dial wraps quic dial
func Dial(addr string, tlsConf *tls.Config, quicConf *quic.Config, timeout time.Duration) (*Conn, error) {
	var (
		ctx    = context.Background()
		cancel func()
	)
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
	}

	session, err := quic.DialAddr(ctx, addr, tlsConf, quicConf)
	if err != nil {
		return nil, err
	}

	stream, err := session.OpenStreamSync(ctx)
	if err != nil {
		return nil, err
	}

	return &Conn{session, stream}, err
}
