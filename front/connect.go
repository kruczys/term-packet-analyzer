package front

import (
	"fmt"
	"log"
	"os"
	"net"
)

// IT MIGHT BE ADDED SOMEWHERE AS ENVIRONMENTAL VARIABLE, MAYBE WE COULD DOCKERIZE OUR WHOLE APP?
const SOCKET_PATH = "/tmp/fiber_orcas.sock"

func connectToSocket() {
	if err := os.Remove(SOCKET_PATH); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	listener, err := net.Listen("unix", SOCKET_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// TO BE DELETED, USED FOR DEVELOPMENT PURPOSES, DEBUG MODE CAN BE ADDED LATER
	fmt.Println("Listening on ", SOCKET_PATH)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Println("Reading error bruv", err)
			}
			return
		}

		data := buffer[:n]
		// TO BE CHANGED
		// CHAIN OF PARSE -> OUTPUT -> SMILE -> HAVE FUN -> REPEAT?, WILL BE ADDED
		fmt.Println(string(data))

		// TO BE CHANGED ADDED JUST FOR FUNSIES WE DONT NEED TO SEND OK TO BACKEND AS IT DOESNT CARE ABOUT THAT
		conn.Write([]byte("OK"))
	}
}