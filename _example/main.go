package main

import (
	"log"
	"sync"
	"time"

	"github.com/osamingo/indigo"
)

// 2009-11-10 23:00:00 UTC
const startedAt = 1257894000

func main() {

	indigo.New(time.Unix(startedAt, 0), nil, nil)

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			id, err := indigo.NextID()
			if err != nil {
				log.Fatalln(err)
			} else {
				log.Println("id:", id)
			}
		}()
	}

	wg.Wait()
}
