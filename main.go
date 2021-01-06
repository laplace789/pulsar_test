package main

import (
	"flag"

	"github.com/laplace789/pulsar_test/config"
	"github.com/laplace789/pulsar_test/input"
)

var cfgDir = flag.String("conf", "", "config dir")

func init() {
	flag.Parse()
}

func main() {
	if *cfgDir == "" {
		panic("no config file")
	}
	cfg := config.Config(*cfgDir)
	in := input.NewInputer("pulsar")
	in.Init(cfg)
	in.PrintStatus()
}
