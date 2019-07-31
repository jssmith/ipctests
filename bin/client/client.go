// Latency benchmark for comparing Unix sockets with TCP sockets.
//
// Idea: ping-pong 128-byte packets between a goroutine acting as a server and
// main acting as client. Measure how long it took to do 2*N ping-pongs and find
// the average latency.
//
// Eli Bendersky [http://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var UnixDomain = flag.Bool("unixdomain", false, "Use Unix domain sockets")
var MsgSize = flag.Int("msgsize", 128, "Message size in each ping")
var RspSize = flag.Int("rspsize", 128, "Response size in each ping")
var NumPings = flag.Int("n", 50000, "Number of pings to measure")
var TCPAddress = flag.String("tcpaddress", "127.0.0.1", "TCP address to bind to")

var UnixAddress = "/tmp/benchmark.sock"

// domainAndAddress returns the domain,address pair for net functions to connect
// to, depending on the value of the UnixDomain flag.
func domainAndAddress() (string, string) {
	if *UnixDomain {
		return "unix", UnixAddress
	} else {
		return "tcp", strings.Join([]string{*TCPAddress, "13500"}, ":")
	}
}

func main() {
	flag.Parse()

	if *RspSize > *MsgSize {
		panic("response size exceeds message size")
	}

	time.Sleep(50 * time.Millisecond)

	// This is the client code in the main goroutine.
	domain, address := domainAndAddress()
	conn, err := net.Dial(domain, address)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, *MsgSize)
	t1 := time.Now()
	for n := 0; n < *NumPings; n++ {
		nwrite, err := conn.Write(buf)
		if err != nil {
			log.Fatal(err)
		}
		if nwrite != *MsgSize {
			log.Fatalf("client: bad nwrite = %d", nwrite)
		}
		sumRead := 0
		for sumRead < *RspSize {
			nread, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			sumRead += nread
		}
	}
	elapsed := time.Since(t1)

	//totalpings := int64(*NumPings * 2)
	fmt.Println("Client done")
	fmt.Printf("%d pingpongs took %d ns\navg. rt latency %d ns\navg. thoughput %f GB/sec\n",
		*NumPings, elapsed.Nanoseconds(),
		elapsed.Nanoseconds()/int64(*NumPings),
		float64(*MsgSize * *NumPings) / elapsed.Seconds() / (1024 * 1024 * 1024),
	)

	time.Sleep(50 * time.Millisecond)
}
