package main

import (
	"github.com/laplace789/pulsar_test/config"
	"github.com/laplace789/pulsar_test/input"
)

func main() {
	cfg := config.Config("./conf/")
	in := input.NewInputer("pulsar")
	in.Init(cfg)
	in.PrintStatus()
}
