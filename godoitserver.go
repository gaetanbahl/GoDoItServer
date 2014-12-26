package main

import (
	"bufio"
	"fmt"
	"github.com/timosis/GoDoItServer/godoit"
	"net"
	"strings"
	"time"
)

const helpstring string = `
List of commands : 
HELP : get this list of commands 
QUIT : save changes and close connection 
RETR : retrieve everything 
DELE N : Delete entry N  
ADD TODO name : add something to do at some point in time 
ADD DOIT name : add something to do like right now 
ADD DONE name : add something that is already done (why the fuck would you want to do that) 
TODO N : mark N as TODO 
DOIT N : mark N as DOIT
DONE N : mark N as DONE 
`

var pass string = "test"

func main() {

	ln, err := net.Listen("tcp", ":1337")
	if err != nil {

		panic("euh quoi")
	}

	//infinite server loop

	for {

		conn, err := ln.Accept()
		if err != nil {
			panic("euh quoi, le retour")
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	conn.Write([]byte("Welcome to GoDoIt Server 1.0 ! What do you want ? (type HELP to see a list of available commands)"))

	time.Sleep(time.Second)

	auth := false

	//answer := make([]byte, 32)

	for {
		answer, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}

		s := strings.Trim(string(answer), " ")

		fmt.Println(s)

		if strings.HasPrefix(s, "HELP") {

			conn.Write([]byte(helpstring))

		} else if strings.HasPrefix(s, "QUIT") {
			conn.Write([]byte("KTHXBYE"))
			conn.Close()
			fmt.Println("Client ended connection with QUIT.")
			break
		} else if strings.HasPrefix(s, "RETR") {

			if auth {
				reponse := godoit.Retrieve() + "\n"
				fmt.Println(reponse)
				conn.Write([]byte(reponse))

			} else {
				conn.Write([]byte("You must type PASS your_password_here first !"))
			}

		} else if strings.HasPrefix(s, "PASS") {

			if strings.Trim(strings.Split(s, " ")[1], " \n"+string(0)) == pass {
				auth = true
				conn.Write([]byte("K Authentification Successful. You can do stuff now ! \n"))
			} else {
				conn.Write([]byte("NOPE Wrong Password !"))
			}
		} else if strings.HasPrefix(s, "DELE") {

		} else {

			conn.Write([]byte("This option is not implemented yet (or you typed random shit, whatever)"))
		}

	}

}

func stripAnswer([]byte) string {
	return "bite"
}
