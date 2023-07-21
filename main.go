package main

import "fmt"

func main() {
	var confrenceName = "Go Conference"
	const confrenceTickets = 50
	var remainningTickets = 50

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickers and %v are still available.\n", confrenceTickets, remainningTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)


}
