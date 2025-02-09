package main

import (
	"bufio"
	"github.com/kruczys/fiberOrca/models"
	"log"
	"net"
	"os"
)

// IT MIGHT BE ADDED SOMEWHERE AS ENVIRONMENTAL VARIABLE
const SOCKET_PATH = "/tmp/fiber_orcas.sock"

// THIS FUNCTION CONNECTS TO CERTAIN UNIX SOCKET AND ACCEPTS ANY INCOMING MESSAGES
// THEN IT CALLBACKS OUR HANDLER FUNCTION THAT READS PACKETS FROM TEMPORARY BUFFER AND DOES SOMETHING WITH THEM
// 1. CONNECT -> 2. READ -> 3. PARSE -> 4. OUTPUT -> REPEAT 2
func connectToSocket(s *models.Session, ch chan string) {
	if err := os.Remove(SOCKET_PATH); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	listener, err := net.Listen("unix", SOCKET_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConnection(conn, s, ch)
	}
}

// FUNCTION THAT GETS PACKETS FROM SOCKET AND DOES SOMETHING WITH THEM
// NO RESEARCH WAS DONE WHETHER IT IS THE MOST MEMORY EFFICIENT IMPLEMENTATION
// NO RESEARCH WAS DONE WHETHER ANY PACKETS ARE LOST
// IT CAUGHT IMPORTANT ONES DURING VERY EXTENSIVE :rofl: TESTING
// IT JUST WORKS
func handleConnection(conn net.Conn, s *models.Session, ch chan string) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		select {
		case <-ch:
			os.Exit(0)
		default:
			packet, err := parsePacket(scanner.Bytes())
			if err != nil {
				log.Println("Parsing error: ", err)
				continue
			}
			if packet != nil {
				s.Update(packet)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Scanner error:", err)
	}
}
