package main

import (
	"log"
	"sync"
	"time"

	"github.com/osamingo/indigo"
)

var g *indigo.Generator

func init() {
	g = indigo.New(indigo.Settings{
		// 2009-11-10 23:00:00 UTC
		StartTime: time.Unix(1257894000, 0),
	})
	_, err := g.NextID()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			id, err := g.NextID()
			if err != nil {
				log.Fatalln(err)
			} else {
				log.Println("ID:", id)
			}
		}()
	}

	wg.Wait()
}
