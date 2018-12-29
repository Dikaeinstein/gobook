package main

type client chan string // an outgoing message channel

type userClient struct {
	name string
	c    client
}

var (
	entering = make(chan userClient)
	leaving  = make(chan userClient)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[userClient]bool)
	for {
		select {
		case cli := <-entering:
			for c := range clients {
				cli.c <- c.name + " is online"
			}
			clients[cli] = true
		case uCli := <-leaving:
			delete(clients, uCli)
			close(uCli.c)
		// Broadcast incoming message to all
		// clients' outgoing message channels.
		case msg := <-messages:
			for uCli := range clients {
				select {
				case uCli.c <- msg:
				default:
					// Discard message if client writer is not ready to accept it
				}

			}
		}
	}
}
