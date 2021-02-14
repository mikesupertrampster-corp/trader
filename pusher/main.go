package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/vrischmann/envconfig"
	"log"
)

type cfg struct {
	Gateway struct {
		Url string `envconfig:"default=http://localhost:9091"`
	}
}

func main() {
	config := new(cfg)
	if err := envconfig.Init(config); err != nil {
		log.Fatal(err)
	}

	single(config.Gateway.Url)
}

func single(url string) {
	price := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "price",
		Help: "Stock Price",
	})
	price.SetToCurrentTime()

	if err := push.New(url, "push_prices").
		Collector(price).
		Push(); err != nil {
		log.Fatal(err)
	}
}