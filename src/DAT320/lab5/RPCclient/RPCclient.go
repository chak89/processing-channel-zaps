// RPCclient

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/signal"
	"time"
)

var rr int
var name string
var running bool
var result bool
var Res ResultBack

type Subscription struct {
	Name string
	RR   int
}

type ResultBack struct {
	Channels string
	Viewers  int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	//Make a channel for interrupt = CTRL+ c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	client, err := rpc.Dial("tcp", service)
	checkError(err)
	// get the user name and refreshrate
	fmt.Print("Enter a name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter a refreshrate for the subscription: ")
	fmt.Scanf("%d", &rr)

	Subs := Subscription{name, rr}
	err = client.Call("Clientlist.Subscribe", Subs, &result)
	checkError(err)
	fmt.Printf("Subscribed to the server = %v \nPress CTRL+ c to unsubscribe and disconnect from the server\n\n", result)
	running = true
	time.Sleep(3 * time.Second)

	for {
		err = client.Call("Clientlist.StatsTop10", rr, &Res)
		fmt.Println(Res, "\n")

		go func() {
			<-c
			running = false
		}()

		if running == false {
			err = client.Call("Clientlist.Unsubscribe", Subs, &result)
			checkError(err)
			os.Exit(-1)
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("key", err)
		os.Exit(1)
	}
}
