package main

import (
	"datafetch/pkg/alphavantage"
	"datafetch/pkg/exporter"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vrischmann/envconfig"
	"log"
	"net/http"
	"time"
)

type cfg struct {
	AlphaVantage struct {
		ApiKey string
	}
}

func main() {
	cfg := new(cfg)
	if err := envconfig.Init(cfg); err != nil {
		log.Fatalln(err)
	}

	symbols := []string{"GME", "NOK", "BB"}
	s := alphavantage.New(cfg.AlphaVantage.ApiKey)
	stockExporter := exporter.NewStocksExporter(s, symbols)
	prometheus.MustRegister(stockExporter)

	if err := createHttpServer().ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func createHttpServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf(":%s", "8080"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
