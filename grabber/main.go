package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mikesupertrampster/trader/grabber/pkg/exporter"
	"github.com/mikesupertrampster/trader/grabber/pkg/services/alphavantage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vrischmann/envconfig"
	"log"
	"net/http"
	"time"
)

type cfg struct {
	Port string `envconfig:"default=8080"`

	SymbolsAPI struct {
		Url string `envconfig:"default=http://localhost:8000"`
	}

	AlphaVantage struct {
		ApiKey string `envconfig:"default=KEY"`
	}

	Finnhub struct {
		ApiKey string `envconfig:"default=KEY"`
	}

	IEX struct {
		ApiKey string `envconfig:"default=KEY"`
	}

	Tiingo struct {
		ApiKey string `envconfig:"default=KEY"`
	}
}

type Symbols []string

func main() {
	config := new(cfg)
	if err := envconfig.Init(config); err != nil {
		log.Fatal(err)
	}

	symbols := symbols(config)

	collect(config, symbols)
}

func symbols(config *cfg) []string {
	req, err := http.NewRequest(http.MethodPost, config.SymbolsAPI.Url, bytes.NewBuffer([]byte(`{"target":"symbols"}`)))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer func(resp *http.Response) {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp)

	var symbols Symbols
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&symbols)
	if err != nil {
		log.Fatal(err)
	}

	return symbols
}

func collect(config *cfg, symbols []string) {
	av := alphavantage.New(config.AlphaVantage.ApiKey)

	stockExporter := exporter.NewStocksExporter(av, symbols)
	prometheus.MustRegister(stockExporter)

	if err := createHttpServer(config.Port).ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func createHttpServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
