# Indigo

[![CI](https://github.com/osamingo/indigo/actions/workflows/actions.yml/badge.svg)](https://github.com/osamingo/indigo/actions/workflows/actions.yml)
[![codecov](https://codecov.io/gh/osamingo/indigo/branch/master/graph/badge.svg)](https://codecov.io/gh/osamingo/indigo)
[![Go Report Card](https://goreportcard.com/badge/osamingo/indigo)](https://goreportcard.com/report/osamingo/indigo)
[![codebeat badge](https://codebeat.co/badges/3885a5d8-7db0-4162-970a-577a1bf54199)](https://codebeat.co/projects/github-com-osamingo-indigo)
[![Maintainability](https://api.codeclimate.com/v1/badges/44865a174db0fad61812/maintainability)](https://codeclimate.com/github/osamingo/indigo/maintainability)
[![GoDoc](https://godoc.org/github.com/osamingo/indigo?status.svg)](https://godoc.org/github.com/osamingo/indigo)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/osamingo/indigo/master/LICENSE)

## About

- A distributed unique ID generator of using Sonyflake and encoded by Base58.
- ID max length is 11 characters by unsigned int64 max value.
- An Encoder can change your original encoder ;)

## Install

```shell
$ go get github.com/osamingo/indigo@latest
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

var g *indigo.Generator

func init() {
	t := time.Unix(1257894000, 0) // 2009-11-10 23:00:00 UTC
	g = indigo.New(nil, indigo.StartTime(t))
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
```

## Benchmark

```
# Machine: MacBook Pro (14-inch, 2021)
# CPU    : Apple M1 Pro
# Memory : 32 GB

goos: darwin
goarch: arm64
pkg: github.com/osamingo/indigo
BenchmarkGenerator_NextID-10        30079       39191 ns/op        7 B/op     1 allocs/op
PASS
```

## Bibliography

- [Sonyflake](https://github.com/sony/sonyflake) - A distributed unique ID generator inspired by Twitter's Snowflake.
- [Base58](https://en.wikipedia.org/wiki/Base58) - Base58 is a group of binary-to-text encoding schemes used to represent large integers as alphanumeric text.

## License

Released under the [MIT License](https://github.com/osamingo/indigo/blob/master/LICENSE).
