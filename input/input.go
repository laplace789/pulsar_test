package input

import (
	"context"
	"log"

	"github.com/laplace789/pulsar_test/config"
)

//TypePulsar represent apache pulsar reader
const (
	TypePulsar = "pulsar"
)

//Inputer is a interface for multi-input
type Inputer interface {
	Init(cfg *config.ServiceCfg) error
	Run(ctx context.Context)
	Stop() error
	CommitMessages(ctx context.Context) error
	PrintStatus()
}

//NewInputer will retrun a inputer base on typ
func NewInputer(typ string) Inputer {
	switch typ {
	case TypePulsar:
		return NewPulsarGo()
	default:
		log.Fatalf("%s is not a supported input type", typ)
		return nil
	}
}
