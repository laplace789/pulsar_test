package main

import (
	"fmt"

	"github.com/laplace789/pulsar_test/config"
)

func main() {
	cfg := config.Config("./conf/")
	fmt.Println(cfg)
}
