package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func randomais(numerorandom int) [4]string {
	var pape [4]string

	switch numerorandom {
	case 0:
		pape = [4]string{"_", "B", "C", "D"}
	case 1:
		pape = [4]string{"A", "_", "C", "D"}
	case 2:
		pape = [4]string{"A", "B", "_", "D"}
	case 3:
		pape = [4]string{"A", "B", "C", "_"}
	default:
		pape = [4]string{"A", "B", "C", "D"}
	}
	return pape
}

func main() {

	rand.Seed(time.Now().UnixNano())
	numeroAleatorio := rand.Intn(4)
	numeroAleatorio2 := rand.Intn(4)
	tablero1 := randomais(numeroAleatorio)
	tablero2 := randomais(numeroAleatorio2)

	udp, err := net.ResolveUDPAddr("udp", ":8080")

	if err != nil {
		fmt.Println("Error en conexión UDP:", err)
		return
	}

	conexion, err := net.ListenUDP("udp", udp)
	if err != nil {
		fmt.Println("Error en conexión UDP:", err)
		return
	}
	defer conexion.Close()

	fmt.Println("Servidor UDP iniciado en el puerto 8080")

	buffer := make([]byte, 1024)
	n, udp, _ := conexion.ReadFromUDP(buffer)
	cIP := udp.IP.String()
	cIP2 := cIP + ":"
	cPort := "8080"
	xPort1 := ":" + cPort
	IPmasPuerto := cIP2 + cPort
	cMsj := string(buffer[:n])
	//fmt.Printf("\nCliente %v:%v ha enviado mensaje: %v\n", cIP, cPort, cMsj)
	if cMsj == "Iniciar" {

		rand.Seed(time.Now().UnixNano())
		numeroAleatorio := rand.Intn(4)
		tablero1 = randomais(numeroAleatorio)

		tablerocliente := fmt.Sprintf("%v%v%v%v", tablero1[0], tablero1[1], tablero1[2], tablero1[3])
		sResp := fmt.Sprintf("%q ha enviado mensaje: %s", IPmasPuerto, tablerocliente)

		_, err := conexion.WriteToUDP([]byte(sResp), udp)

		if err != nil {
			fmt.Println("Error al escribir a la conexión UDP: ", err)
		}

	}

	listener, err := net.Listen("tcp", xPort1)
	if err != nil {
		fmt.Println("Error al escuchar:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			return
		}

		fmt.Println("Cliente conectado desde", conn.RemoteAddr())

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer los datos:", err)

			return
		}

		// Imprimir el mensaje recibido del cliente
		mensaje := ""
		fmt.Println("Mensaje recibido: %v", string(buffer[:n]))
		fmt.Println(string(buffer))
		fmt.Println(tablero2[0])
		fmt.Println(tablero2[1])
		fmt.Println(tablero2[2])
		fmt.Println(tablero2[3])
		switch string(buffer[:n]) {
		case string(tablero2[0]):
			mensaje = "D"
		case string(tablero2[1]):
			mensaje = "A"
		case string(tablero2[2]):
			mensaje = "B"
		case string(tablero2[3]):
			mensaje = "C"
		default:
			break
		}
		fmt.Println(mensaje)
		_, err = conn.Write([]byte(mensaje))
		if err != nil {
			fmt.Println("Error al enviar mensaje:", err)
			return
		}
		// Cerrar la conexión

	}

	tab := [4]string{"1", "0", "0", "0"}

	for {

	}

}
