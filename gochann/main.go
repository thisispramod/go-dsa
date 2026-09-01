package main

import (
	"fmt"
	"time"
)

func fetchApi(ch chan string) {
	ch <- "newonew"
	time.Sleep(time.Second)

}

func main() {

	/*
		unbufferedch := make(chan int)

		go func() {
			fmt.Println("sending to unbuffered channel..")
			unbufferedch <- 100
			fmt.Println("Sending 100 to unbuffered channel")
		}()
		time.Sleep(1 * time.Second)
		fmt.Println("Received from unbuffered : ", unbufferedch)

		bufferedch := make(chan string, 2)

		bufferedch <- "Data 1"
		bufferedch <- "Data 2"
		fmt.Println("sending two message to buffered channel without blocking..")

		fmt.Println("these data from buffered channel:", <-bufferedch)
		fmt.Println("these data from buffered channel:", <-bufferedch)
	*/
	/*
		ch := make(chan string)
		go fetchApi(ch)
		select {
		case res := <-ch:
			fmt.Println("Rechieved ", res)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout ! API Took too long")
		}
	*/
	/*
		email := EmailService{}

		NotifyUser(email, "interview is sechudled at 4am")
	*/

	acc := Account{Balance: 100}
	acc.DepositValue(50)
	fmt.Println("after DepositValue:", acc.Balance)

	// pointer receiver test

	acc.DepositPointer(50)
	fmt.Println("After deposit pointer", acc.Balance)
}

type notifier interface {
	Send(message string)
}

type EmailService struct{}

func (e EmailService) Send(message string) {
	fmt.Printf("Sending Mail %s\n", message)
}

func NotifyUser(n notifier, msg string) {
	n.Send(msg)
}

type Account struct {
	Balance int
}

func (a Account) DepositValue(amount int) {
	a.Balance += amount
}

func (a *Account) DepositPointer(amount int) {
	a.Balance += amount
}
