package main

import byoredisgo "github.com/ahmadnaufal/byo-redis-go"

func main() {
	if err := byoredisgo.StartServer(); err != nil {
		panic(err)
	}
}
