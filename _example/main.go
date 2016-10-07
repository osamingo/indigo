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

	g, err := indigo.New(indigo.Settings{
		StartTime: time.Unix(startedAt, 0),
	})
	if err != nil {
		log.Fatalln(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			id, err := g.NextID()
			if err != nil {
				log.Fatalln(err)
			} else {
				log.Println("id:", id)
			}
		}()
	}

	wg.Wait()
}
