package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:porttttt")
	if err != nil {
		fmt.Println("Error al resolver la dirección del servidor:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error al conectarse al servidor UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()
}
