package main

import (
	"fmt"
	"ticketApp/internal/ticket"
)

func main() {
	for {
		fmt.Println(ticket.CreateTicket())
	}

}
