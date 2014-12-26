package main

import (
	"bufio"
	"fmt"
	"github.com/timosis/GoDoItServer/godoit"
	"net"
	"strconv"
	"strings"
	"time"
)

const helpstring string = `
List of commands : 
HELP : get this list of commands 
QUIT : save changes and close connection
PASS <password> : get access to the useful commands below
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

	conn.Write([]byte("Welcome to GoDoIt Server 1.0 ! What do you want ? (type HELP to see a list of available commands)\n"))

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
			conn.Write([]byte("KTHXBYE\n"))
			conn.Close()
			fmt.Println("Client ended connection with QUIT.\n")
			break
		} else if strings.HasPrefix(s, "RETR") {

			if auth {
				reponse := godoit.Retrieve() + "\n"
				fmt.Println(reponse)
				conn.Write([]byte(reponse))

			} else {
				conn.Write([]byte("You must type PASS your_password_here first !\n"))
			}

		} else if strings.HasPrefix(s, "PASS") {

			stuff = strings.Split(s, " ")
			if len(stuff) == 2 {

				if strings.Trim(stuff[1], " \n"+string(0)) == pass {
					auth = true
					conn.Write([]byte("K Authentification Successful. You can do stuff now ! \n"))
				} else {
					conn.Write([]byte("NOPE Wrong Password !\n"))
				}
			} else {

				conn.Write([]byte("usage : PASS <password> \n"))
			}

		} else if strings.HasPrefix(s, "DELE") {

			stuff = strings.Split(s, " ")
			if len(stuff) == 2 {
				godoit.DeleteItem(strconv.Atoi(stuff[1]))
				conn.Write([]byte("OK Item Successfully Deleted\n"))
			} else {
				conn.Write([]byte("NOPE Usage : DELE N where N is an integer\n"))
			}
		} else if strings.HasPrefix(s, "") == "ADD" {
			stuff = strings.Split(s, " ")

			if len(stuff) > 2 {

				if stuff[1] == "TODO" {
					godoit.CreateItem("TODO", strings.Join(stuff[2:len(stuff)]))
				} else if stuff[1] == "DOIT" {
					godoit.CreateItem("DOIT", strings.Join(stuff[2:len(stuff)]))
				} else if stuff[1] == "DONE" {
					godoit.CreateItem("DONE", strings.Join(stuff[2:len(stuff)]))
				} else {
					conn.Write([]byte("NOPE this command doesn't exist ... \n"))
				}

			} else {
				conn.Write([]byte("NOPE Usage : ADD <state> <name> \n"))
			}
		} else {

			conn.Write([]byte("NOPE :This option is not implemented yet (or you typed random shit, whatever) \n"))
		}

	}

}
