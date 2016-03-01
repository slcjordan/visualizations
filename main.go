package main

import (
	"fmt"
	// _ "github.com/mkevac/debugcharts"
	_ "github.com/rakyll/gom/http"
	"net/http"
	_ "net/http/pprof"
)

// filter out multiples of a single prime number.
func filter(prime int64, in <-chan int64) <-chan int64 {
	out := make(chan int64)

	go func() {
		defer close(out)

		for i := range in {
			if (i % prime) != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {

	// pprof requires an http server
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	in := make(chan int64)
	var out <-chan int64 = in

	// insert all numbers 2 to 5 million into start of pipe.
	go func() {
		defer close(in)

		for i := int64(2); i < 5000000; i++ {
			in <- i
		}
	}()

	// read numbers off from end of pipe.
	for {
		i, ok := <-out
		if !ok {
			break
		}
		fmt.Println(i)
		out = filter(i, out)
	}

}
