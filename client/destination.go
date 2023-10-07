package client

import (
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

type Destination struct {
	writeStart bool
	address    string
	conn       *net.TCPConn
}

func NewDestination(address string) (*Destination, error) {
	d_tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}

	d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
	if err != nil {
		return nil, err
	}
	d_conn.SetKeepAlive(true)
	d_conn.SetKeepAlivePeriod(30 * time.Second)
	d_conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteBuffer(2048)
	return &Destination{
		conn:       d_conn,
		address:    address,
		writeStart: false,
	}, nil
}

func (d *Destination) WriteStart() {
	d.writeStart = true
}

func (d *Destination) CanRead() bool {
	return d.writeStart
}

func (d *Destination) Write(data []byte) (int, error) {
	n, err := d.conn.Write(data)
	return n, err
}

func (d *Destination) Read(data []byte) (int, error) {
	n, err := d.conn.Read(data)
	return n, err
}

func (d *Destination) Refresh() {
	if d.conn != nil {
		d.conn.Close()
	}
	d_tcpAddr, err := net.ResolveTCPAddr("tcp4", d.address)
	if err != nil {
		logrus.Errorf("refresh dest error: %v", err)
		return
	}

	d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
	if err != nil {
		logrus.Errorf("refresh dest error: %v", err)
		return
	}
	d_conn.SetKeepAlive(true)
	d_conn.SetKeepAlivePeriod(30 * time.Second)
	d_conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteBuffer(2048)

	d.conn = d_conn
}
