package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	From    string
	Payload string
}
type server struct {
	MsgCH  chan Message
	ExitCH chan int
	InternalWG sync.WaitGroup
	WG *sync.WaitGroup
}

func (s *server) Up() {
	defer s.WG.Done()
	for {
		select {
		case msg := <-s.MsgCH:
			s.InternalWG.Add(1)
			go s.Operation(msg)
		case <-s.ExitCH:
			close(s.MsgCH)
			close(s.ExitCH)
			s.InternalWG.Wait()
			return
		}
	}
}

func (s *server) Operation(msg Message) {
	time.Sleep(4 * time.Second)
	log.Printf("server response for this message : %s --- %s\n", msg.From, msg.Payload)
	s.InternalWG.Done()
}

func SendMessageToServer(msg Message, msgCH chan Message) {
	msgCH <- msg
}

func SendExitSignalToServer(exitCH chan int) {
	exitCH <- 1
}

func main() {
	wg := sync.WaitGroup{}
	server := &server{
		MsgCH:  make(chan Message),
		ExitCH: make(chan int),
		InternalWG: sync.WaitGroup{},
		WG: &wg,
	}
	wg.Add(1)
	go server.Up()

	for {
		fmt.Println("enter option ( 1 for operation 2 for exit ) :")
		var option string
		fmt.Scanln(&option)
		if option == "1" {
			var from, payload string
			fmt.Scanln(&from)
			fmt.Scanln(&payload)
			msg := Message{
				Payload: payload,
				From:    from,
			}
			SendMessageToServer(msg, server.MsgCH)
		} else {
			SendExitSignalToServer(server.ExitCH)
			break
		}
	}
	wg.Wait()
}
