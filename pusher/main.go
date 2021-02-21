package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/mikesupertrampster/trader/apis/services/alphavantage"
	"github.com/vrischmann/envconfig"
	"log"
)

type pusher struct {
	av       alphavantage.Client
	client   client.Client
	database database
}

type database struct {
	name      string
	precision string
}

type cfg struct {
	InfluxDB struct {
		Address string `envconfig:"default=http://localhost:8086"`
		User    struct {
			Name     string `envconfig:"default=admin"`
			Password string `envconfig:"default=test"`
		}
		Database struct {
			Name      string `envconfig:"default=ticker"`
			Precision string `envconfig:"default=s"`
		}
	}

	AlphaVantage struct {
		ApiKey string `envconfig:"default=demo"`
	}
}

func main() {
	config := new(cfg)
	if err := envconfig.Init(config); err != nil {
		log.Fatal(err)
	}

	c := influxDBClient(config.InfluxDB.Address, config.InfluxDB.User.Name, config.InfluxDB.User.Password)
	defer func(c client.Client) {
		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(c)

	av := alphavantage.New(config.AlphaVantage.ApiKey)
	p := pusher{
		av:     av,
		client: c,
		database: database{
			name:      config.InfluxDB.Database.Name,
			precision: config.InfluxDB.Database.Precision,
		},
	}

	for _, symbol := range []string{"IBM"} {
		p.PushIntra(symbol, "5min", "full")
	}
}
