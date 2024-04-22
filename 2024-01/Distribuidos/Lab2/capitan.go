package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveTCPAddr("tcp", "localhost:9799")
	if err != nil {
		fmt.Println("Error al resolver la dirección del servidor:", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error al conectarse al servidor UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("¡Hola, servidor!"))
	if err != nil {
		fmt.Println("Error al enviar datos al servidor:", err)
		return
	}

	max := 10
	min := 0
	var botines [6]int
	counter := 0
	for {
		if counter == 6 {
			break
		}
		botines[counter] = rand.Intn(max - min)
		counter += 1
	}

	

	conn.Close()
}
