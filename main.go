package main

import (
	"sync"
)

func main() {
	srv := MakeHandler()

	defer srv.db.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.HTTPServeMain()
	}()
	wg.Wait()
}
