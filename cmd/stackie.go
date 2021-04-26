package main

import (
	"github.com/bjerkio/stackie/pkg/stackie"
)

func main() {
	err := stackie.Setup()
	if err != nil {
		panic(err)
	}
}