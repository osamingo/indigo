# Indigo

[![CircleCI](https://img.shields.io/circleci/project/osamingo/indigo/master.svg)](https://circleci.com/gh/osamingo/indigo)
[![codecov](https://codecov.io/gh/osamingo/indigo/branch/master/graph/badge.svg)](https://codecov.io/gh/osamingo/indigo)
[![Go Report Card](https://goreportcard.com/badge/osamingo/indigo)](https://goreportcard.com/report/osamingo/indigo)
[![codebeat badge](https://codebeat.co/badges/3885a5d8-7db0-4162-970a-577a1bf54199)](https://codebeat.co/projects/github-com-osamingo-indigo)
[![GoDoc](https://godoc.org/github.com/osamingo/indigo?status.svg)](https://godoc.org/github.com/osamingo/indigo)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/osamingo/indigo/master/LICENSE)

## About

A distributed unique ID generator of using Sonyflake and encoded by Base58.

## Install

```bash
$ go get -u github.com/osamingo/indigo
```

## Usage

```go
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/osamingo/indigo"
)

func randomMachineID() (uint16, error) {
	rand.Seed(time.Now().UnixNano())
	return uint16(rand.Intn(65535)), nil
}

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
```

## Bibliography

- [Sonyflake](https://github.com/sony/sonyflake) - A distributed unique ID generator inspired by Twitter's Snowflake.
- [Base58](https://en.wikipedia.org/wiki/Base58) - Base58 is a group of binary-to-text encoding schemes used to represent large integers as alphanumeric text.

## License

Released under the [MIT License](https://github.com/osamingo/indigo/blob/master/LICENSE).
