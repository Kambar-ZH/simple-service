package main

import (
	"github.com/Kambar-ZH/simple-service/internal/core"
)

func main() {
	if err := core.Run(); err != nil {
		panic(err)
	}
}
