package client

import (
	"net"
	"time"
)

type Destination struct {
	writeFinish bool
	address     string
	conn        *net.TCPConn
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
		conn:        d_conn,
		address:     address,
		writeFinish: false,
	}, nil
}

func (d *Destination) WriteFinish() {
	d.writeFinish = true
}

func (d *Destination) CanRead() bool {
	return d.conn != nil
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
		return
	}

	d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
	if err != nil {
		return
	}
	d_conn.SetKeepAlive(true)
	d_conn.SetKeepAlivePeriod(30 * time.Second)
	d_conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteBuffer(2048)

	d.conn = d_conn
}
