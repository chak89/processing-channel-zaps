// UDP_server

package main

import (
	"DAT320/lab5/zaplab/Store_Zap_Events"
	"DAT320/lab5/zaplab/zmap"
	"DAT320/lab5/zaplab/ztorage"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime/pprof"
	"strings"
	"sync"
	"time"
)

var mm sync.Mutex
var nzapstore = ztorage.NewZapStore() //For using SLICE as storage
var nzmapstore = zmap.NewZmapStore()  //For using MAP as storage

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "224.0.1.130:10000")
	checkError(err)

	conn, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkError(err)

	for {
		handleClient(conn)

		//Enable this to get the number of NRK1 viewers, using SLICE
		/*go NumberofViewers("NRK1") */

		//Enable this to get the number of TV2 Norge viewers, using SLICE
		/*go NumberofViewers("TV2 Norge") */

		//Enable this to get the number of entries in the storage, using SLICE
		/*go EntriesInStorage(nzapstore) */

		//Enable this to activate the memory profiler
		/* memoryprofile() */

		//Enable this to get the number of NRK1 viewers, using MAP
		/* go NumberofViewersMap("NRK1") */

		//Enable this to get the number of entries in the storage, using MAP
		/* go EntriesInStorageMap(nzmapstore) */

		//Enable this to get the top-10 channels list, using MAP
		go DisplayTop10()

	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	m, _, err := conn.ReadFromUDP(buf[0:])
	checkError(err)

	if len(strings.Split(string(buf[:m]), ", ")) == 5 {
		SSS := *Store_Zap_Events.NewStoreZapEvents(string(buf[:m]))
		//fmt.Printf("%s\n", SSS)

		//Enable this, if intended to use SLICE as storage
		/* nzapstore.StoreZap(SSS) */

		//Enable this, if intended to use MAP as storage
		nzmapstore.StoreZmap(SSS)
	}
}

//Function for getting number of viewers, using SLICE
func NumberofViewers(chName string) {
	mm.Lock()
	for {
		time.Sleep(1 * time.Second)
		chviewers := nzapstore.ComputeViewers(chName)
		fmt.Printf("\nNumber of %s viewers: %d\n\n", chName, chviewers)
	}
	mm.Unlock()
}

//Function for getting number of viewers, using MAP
func NumberofViewersMap(chName string) {
	mm.Lock()
	for {
		time.Sleep(1 * time.Second)
		chviewers := nzmapstore.ComputeViewers(chName)
		fmt.Printf("\nNumber of %s viewers: %d\n\n", chName, chviewers)
	}
	mm.Unlock()
}

//Function for checking the SLICE length aka entries in storage
func EntriesInStorage(N_entries *ztorage.Zaps) {
	mm.Lock()
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("Number of entries in the storage using slice: %d \n", len(*N_entries))
	}
	mm.Unlock()

}

//Function for checking the MAP length aka entries in storage
func EntriesInStorageMap(N_entries *zmap.Zmap) {
	mm.Lock()
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("Number of entries in the storage using map: %d \n", len(N_entries.V))
	}
	mm.Unlock()
}

//Function for the top-10 list, using MAP
func DisplayTop10() {
	mm.Lock()
	for {
		time.Sleep(2 * time.Second)
		slices := nzmapstore.ComputeTop10()
		for i := range slices {
			fmt.Println()
			fmt.Printf("%d.place:", i+1)
			fmt.Println(slices[i])
		}
		fmt.Println("\n\n")
	}
	mm.Unlock()
}

//Function for the memory profiler
func memoryprofile() {
	flag.Parse()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
