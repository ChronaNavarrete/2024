package main

import (
	"fmt"
	"net"
)

func main() {
	server := "localhost:8080"

	conn, err := net.Dial("udp", server)
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	fmt.Println("éxito")
	defer conn.Close()

	_, err = conn.Write([]byte("Iniciar"))
	if err != nil {
		fmt.Println("Error al enviar datos al servidor:", err)
		return
	}
	fmt.Println("éxito2")
	// Suponiendo que MensajeUDP es una estructura que has definido en otro lugar
	var tablero [1024]byte // Tamaño del buffer a ajustar según necesidad
	n, err := conn.Read(tablero[:])
	if err != nil {
		fmt.Println("Error al recibir respuesta del servidor:", err)
		return
	}
	fmt.Println("éxito3")
	// Procesa los datos recibidos aquí
	fmt.Println("Recibido:", string(tablero[:n]))
	fmt.Println("Recibido:", string(tablero[31:]))
	localm := "[::1]" + string(tablero[4:9])
	conn2, err2 := net.Dial("tcp", localm)

	if err2 != nil {
		fmt.Println("Error al conectar al servidor:", err)
	}
	defer conn2.Close()

	var smigh string
	tableromio := tablero[31:]

	for {
		fmt.Print("Ingresa la casilla donde quiera atacar: ")
		fmt.Scan(&smigh)
		mensaje := smigh
		_, err = conn2.Write([]byte(mensaje))
		if err != nil {
			fmt.Println("Error al enviar mensaje:", err)
			return
		}
		buffer := make([]byte, 1024)
		n, err := conn2.Read(buffer)
		fmt.Println("HOLAAAAA")
		if err != nil {
			fmt.Println("Error al leer los datos:", err)

			return
		}
		switch string(buffer[:n]) {
		case string(tableromio[0]):
		case string(tableromio[1]):
		case string(tableromio[2]):
		case string(tableromio[3]):
		default:
			break
		}

	}

}
