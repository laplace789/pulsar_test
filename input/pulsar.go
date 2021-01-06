package input

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/laplace789/pulsar_test/config"
)

//PulsarGo is the struct for Pulsar inputer
type PulsarGo struct {
	cfg      *config.PulasrCfg
	consumer pulsar.Consumer
	msgChan  chan pulsar.ConsumerMessage
	earliest bool
	mux      sync.Mutex
}

//NewPulsarGo will return a empty PulsarGo
func NewPulsarGo() *PulsarGo {
	return &PulsarGo{}
}

//Init will init a pulsar inputer
func (p *PulsarGo) Init(cfg *config.ServiceCfg) error {
	p.cfg = &cfg.Pulsar
	port := strconv.Itoa(cfg.Pulsar.Port)
	connectStr := fmt.Sprint("pulsar://" + cfg.Pulsar.Server + ":" + port)
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: connectStr,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	opts := pulsar.ConsumerOptions{
		Topic:            cfg.Task.Topic,
		SubscriptionName: cfg.Task.SubscriptionName,
	}
	p.msgChan = opts.MessageChannel
	consumer, err := client.Subscribe(opts)
	if err != nil {
		log.Fatal(err)
		return err
	}
	p.consumer = consumer
	// if all the config sucess return nul
	return nil
}

/*
Run will pull data from to msgChan
can't use reader since pulsar.reader can't commit message
*/
func (p *PulsarGo) Run(ctx context.Context) {
	for cm := range p.msgChan {
		msg := cm.Message
		fmt.Printf("Received message  msgId: %v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		p.consumer.Ack(msg)
	}
}

//Stop will stop pulsar.reader and commit the rest data to pulsar
func (p *PulsarGo) Stop() error {
	if p.consumer != nil {
		p.consumer.Close()
	}
	return nil
}

//CommitMessages will commit the data has benn read from pulsar
func (p *PulsarGo) CommitMessages(ctx context.Context) error {
	//TODO use AckID
	return nil
}

//PrintStatus will print corrent config
func (p *PulsarGo) PrintStatus() {
	fmt.Println(p.consumer)
}
