//RPCserver

package main

import (
	"DAT320/lab5/zaplab/zmap"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"sync"
	"time"
)

var nzmapstore = zmap.NewZmapStore()
var running bool

type Subscription struct {
	Name string
	RR   int
}

type Clientlist struct {
	lock *sync.Mutex
	Map  map[string]int
}

type ResultBack struct {
	Channels string
	Viewers  int
}

func (cl *Clientlist) StatsTop10(rate int, R *ResultBack) error {
	time.Sleep(time.Duration(rate) * time.Second)
	R.Channels = "Har ikke fått gjort dette enda! Antall klient på serveren:"
	R.Viewers = len(cl.Map)
	//Res := nzmapstore.ComputeTop10()
	fmt.Printf("Kan desverre ikke sende statistikken til klienten, Klient tilkoblet: %d\n\n", len(cl.Map))
	//fmt.Println("Type: ", reflect.TypeOf(ResultList)) // type slice of struct
	return nil
}

func (cl *Clientlist) Subscribe(input Subscription, result *bool) error {

	cl.lock.Lock()
	cl.Map[input.Name] = input.RR
	*result = true
	fmt.Printf("Added client: %sRefreshrate:%d\n", input.Name, input.RR)
	fmt.Printf("Subscribers at the moment:%d\n", len(cl.Map))
	cl.lock.Unlock()
	return nil
}

func (cl *Clientlist) Unsubscribe(input Subscription, result *bool) error {
	cl.lock.Lock()
	delete(cl.Map, input.Name)
	fmt.Printf("Client: %shas unsubscribed\n", input.Name)
	fmt.Printf("Subscribers at the moment:%d\n", len(cl.Map))
	cl.lock.Unlock()
	return nil
}

func main() {

	cl := &Clientlist{new(sync.Mutex), make(map[string]int)}
	rpc.Register(cl)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":55555")
	checkError(err)

	fmt.Printf("Server now runs on port: 55555\n\n")

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
