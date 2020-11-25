package main

import (
	"fmt"
	"os"

	"github.com/bounoable/dragoman/internal/cli"
)

func main() {
	if err := cli.New().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
