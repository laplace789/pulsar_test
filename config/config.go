package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//ServiceCfg represent whole service config
type ServiceCfg struct {
	Pulsar     PulasrCfg
	Clickhouse ClickhouseCfg
	Task       TaskCfg
}

//PulasrCfg represent Pulasr service config
type PulasrCfg struct {
	Server string
	Port   int
}

//ClickhouseCfg represent clickhouse  service config
type ClickhouseCfg struct {
	Server string
	Port   int
}

//TaskCfg represent task  config
type TaskCfg struct {
	Topic            string
	SubscriptionName string
	Earliest         bool
}

//Config will get the value from path
func Config(path string) *ServiceCfg {
	// Config
	viper.SetConfigName("service") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv() // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	// Set default value

	clickhouse := &ClickhouseCfg{}
	pulsar := &PulasrCfg{}
	task := &TaskCfg{}

	pulsar.Server = viper.GetString("pulsar.Server")
	pulsar.Port = viper.GetInt("pulsar.Port")

	clickhouse.Server = viper.GetString("clickhouse.Server")
	clickhouse.Port = viper.GetInt("clickhouse.Port")

	task.Topic = viper.GetString("task.Topic")
	task.SubscriptionName = viper.GetString("task.SubscriptionName")
	task.Earliest = viper.GetBool("task.Earliest")

	return &ServiceCfg{Pulsar: *pulsar, Clickhouse: *clickhouse, Task: *task}
}
