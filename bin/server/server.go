package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var UnixDomain = flag.Bool("unixdomain", false, "Use Unix domain sockets")
var MsgSize = flag.Int("msgsize", 128, "Message size in each ping")
var RspSize = flag.Int("rspsize", 128, "Response size in each ping")
var NumPings = flag.Int("n", 50000, "Number of pings to measure")

var TcpAddress = "127.0.0.1:13500"
var UnixAddress = "/tmp/benchmark.sock"

// domainAndAddress returns the domain,address pair for net functions to connect
// to, depending on the value of the UnixDomain flag.
func domainAndAddress() (string, string) {
	if *UnixDomain {
		return "unix", UnixAddress
	} else {
		return "tcp", TcpAddress
	}
}

func server() {
	if *RspSize > *MsgSize {
		panic("response size exceeds message size")
	}
	if *UnixDomain {
		if err := os.RemoveAll(UnixAddress); err != nil {
			panic(err)
		}
	}

	domain, address := domainAndAddress()
	l, err := net.Listen(domain, address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, *MsgSize)
	for n := 0; n < *NumPings; n++ {
		sumRead := 0
		for sumRead < *MsgSize {
			nread, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			sumRead += nread
		}
		nwrite, err := conn.Write(buf[:*RspSize])
		if err != nil {
			log.Fatal(err)
		}
		if nwrite != *RspSize {
			log.Fatalf("server: bad nwrite = %d", nwrite)
		}
	}

	time.Sleep(50 * time.Millisecond)
}

func main() {
	flag.Parse()

	server()
}