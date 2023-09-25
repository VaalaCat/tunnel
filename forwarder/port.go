package forwarder

import (
	"fmt"
	"io"
	"net"
	"os"
	"tunnel/protogen"
)

func ListenAndServe(localport string, cli protogen.TunnelServer_CallServer) (err error) {
	var dest_con net.Conn

	host := "0.0.0.0"
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, localport))
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}

	defer l.Close()

	for {
		dest_con, err = l.Accept()
		if err != nil {
			continue
		}

		go func() {
			recv_data := make([]byte, 1024)
			for {
				n, err := dest_con.Read(recv_data)
				if err != nil {
					fmt.Println(err, err.Error())
					dest_con.Close()
					break
				}

				if err := cli.Send(&protogen.Response{Payload: recv_data[:n]}); err != nil {
					fmt.Println(err, err.Error())
					dest_con.Close()
					break
				}
			}
		}()
		go func() {
			for {
				in, err := cli.Recv()
				if err == io.EOF {
					dest_con.Close()
					return
				}
				if err != nil {
					fmt.Println(err, err.Error())
					dest_con.Close()
					return
				}
				if _, err := dest_con.Write(in.Payload); err != nil {
					fmt.Println(err, err.Error())
					dest_con.Close()
					return
				}
			}
		}()
	}
}
