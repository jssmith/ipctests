package main

import (
	"fmt"
	"os"
	gs "github.com/dangerousHobo/go-semaphore"
	"strconv"
)

func main() {
	semName := "/testsem1"
	semNameR := "/testsem1r"
	mode := os.Args[1]
	count, _ := strconv.Atoi(os.Args[2])
	switch mode {
	case "server":
		fmt.Printf("running as server\n")
		var sem gs.Semaphore
		var semr gs.Semaphore
		if err := sem.Open(semName, 0644, 0); err != nil {
			panic(err)
		}
		if err := semr.Open(semNameR, 0644, 0); err != nil {
			panic(err)
		}
		for i := 0; i < count; i++ {
			if err := sem.Wait(); err != nil {
				panic(err)
			}
			if err := semr.Post(); err != nil {
				panic(err)
			}
			//if i % 1000 == 0 {
			//	v, _ := sem.GetValue()
			//	vr, _ := semr.GetValue()
			//	fmt.Printf("value is %+v %+v\n", v, vr)
			//}
		}

		if err := sem.Close(); err != nil {
			panic(err)
		}

		if err := sem.Unlink(); err != nil {
			panic(err)
		}

		if err := semr.Close(); err != nil {
			panic(err)
		}

		if err := semr.Unlink(); err != nil {
			panic(err)
		}

	case "client":
		fmt.Printf("running as client\n")
		var sem gs.Semaphore
		var semr gs.Semaphore
		if err := sem.Open(semName, 0644, 1); err != nil {
			panic(err)
		}

		if err := semr.Open(semNameR, 0644, 1); err != nil {
			panic(err)
		}

		for i := 0; i < count; i++ {
			if err := sem.Post(); err != nil {
				panic(err)
			}
			if err := semr.Wait(); err != nil {
				panic(err)
			}
			//fmt.Printf("client: %d", i)
		}

		if err := sem.Close(); err != nil {
			panic(err)
		}

		if err := semr.Close(); err != nil {
			panic(err)
		}

	default:
		fmt.Printf("ERROR - must specify server or client\n")
	}
}
