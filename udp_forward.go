package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

type UDPForward struct {
	s_host string
	s_port string
	d_host string
	d_port string
}

func main() {
	args_arr := &UDPForward{
		d_host: os.Args[1],
		d_port: os.Args[2],
		s_host: os.Args[3],
		s_port: os.Args[4],
	}

	packageConn, conn := createConnections(*args_arr)

	src_package := make([]byte, 64)
	dst_package := make([]byte, 64)

	defer packageConn.Close()
	defer conn.Close()

	for {
		_, addr, err := packageConn.ReadFrom(src_package)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = conn.Write(src_package)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = conn.Read(dst_package)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = packageConn.WriteTo(dst_package, addr)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func createConnections(args UDPForward) (net.PacketConn, net.Conn) {
	packageConn, err := net.ListenPacket("udp", fmt.Sprintf("%s:%s", args.d_host, args.d_port))
	if err != nil {
		log.Fatalln("Listing Error:", err)
	}
	fmt.Printf("Listen %s:%s...", args.d_host, args.d_port)
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", args.s_host, args.s_port))

	if err != nil {
		log.Fatalln("Google connection: ", err)
	}
	return packageConn, conn
}
