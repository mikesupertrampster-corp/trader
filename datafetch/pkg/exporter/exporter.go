package exporter

import (
	"datafetch/pkg/alphavantage"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type StocksExporter struct {
	client alphavantage.Client
	price  *prometheus.Desc
	symbols []string
}

func NewStocksExporter(client alphavantage.Client, symbols []string) *StocksExporter {
	return &StocksExporter{
		client: client,
		price: prometheus.NewDesc("price", "Stock Price", []string{"symbol"}, nil),
		symbols: symbols,
	}
}

func (e *StocksExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.price
}

func (e *StocksExporter) Collect(ch chan<- prometheus.Metric) {
	for _, symbol := range e.symbols {
		q, err := e.client.GetQuote(symbol)
		if err != nil {
			return
		}

		price, err := strconv.ParseFloat(q.GlobalQuote.Price, 64)
		if err != nil {
			return
		}

		ch <- prometheus.MustNewConstMetric(
			e.price, prometheus.GaugeValue, price, symbol,
		)
	}
}