# Indigo

[![CircleCI](https://img.shields.io/circleci/project/osamingo/indigo/master.svg)](https://circleci.com/gh/osamingo/indigo)
[![codecov](https://codecov.io/gh/osamingo/indigo/branch/master/graph/badge.svg)](https://codecov.io/gh/osamingo/indigo)
[![Go Report Card](https://goreportcard.com/badge/osamingo/indigo)](https://goreportcard.com/report/osamingo/indigo)
[![codebeat badge](https://codebeat.co/badges/3885a5d8-7db0-4162-970a-577a1bf54199)](https://codebeat.co/projects/github-com-osamingo-indigo)
[![GoDoc](https://godoc.org/github.com/osamingo/indigo?status.svg)](https://godoc.org/github.com/osamingo/indigo)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/osamingo/indigo/master/LICENSE)

## About

A distributed unique ID generator of using Sonyflake and encoded by Base58.  
Base58 logic is optimized unsigned int64.

- ID max length is 11 characters by unsigned int64 max value.
- Default characters: `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`

## Install

```bash
$ go get -u github.com/osamingo/indigo
```

## Usage

```go
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
```

## Benchmark

```
# Machine: MacBook Pro (Retina, 15-inch, Mid 2015)
# CPU    : 2.8 GHz Intel Core i7
# Memory : 16 GB 1600 MHz DDR3

BenchmarkEncodeBase58-8    20000000      104 ns/op    46 B/op    1 allocs/op
BenchmarkDecodeBase58-8    10000000      238 ns/op     0 B/op    0 allocs/op
BenchmarkNextID-8             50000    38943 ns/op     8 B/op    1 allocs/op
```

## Bibliography

- [Sonyflake](https://github.com/sony/sonyflake) - A distributed unique ID generator inspired by Twitter's Snowflake.
- [Base58](https://en.wikipedia.org/wiki/Base58) - Base58 is a group of binary-to-text encoding schemes used to represent large integers as alphanumeric text.

## License

Released under the [MIT License](https://github.com/osamingo/indigo/blob/master/LICENSE).
