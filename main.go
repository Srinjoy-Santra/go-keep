package main

import (
	"go-keep/initialize"
)

func main() {

	err := initialize.Run()
	if err != nil {
		panic(err)
	}
}
