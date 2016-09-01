package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/osamingo/indigo"
)

func main() {

	indigo.New(time.Now(), randomMachineID, nil)

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

func randomMachineID() (uint16, error) {
	rand.Seed(time.Now().UnixNano())
	return uint16(rand.Intn(65535)), nil
}
