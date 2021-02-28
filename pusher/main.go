package main

import (
	"github.com/mikesupertrampster/trader/apis/services/alphavantage"
	"github.com/mikesupertrampster/trader/pusher/pkg/influx"
	"github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
	"os"
)

type cfg struct {
	InfluxDB struct {
		Address string `envconfig:"default=http://localhost:8086"`
		User    struct {
			Name     string `envconfig:"default=admin"`
			Password string `envconfig:"default=test"`
		}
		Database struct {
			Name      string `envconfig:"default=trader"`
			Precision string `envconfig:"default=s"`
		}
	}

	AlphaVantage struct {
		ApiKey string `envconfig:"default=Demo"`
	}

	Symbols []string `envconfig:"default=IBM"`
}

var logger = logrus.New()

func init() {
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
}

func main() {
	config := new(cfg)
	if err := envconfig.Init(config); err != nil {
		logger.Fatal("Failed to read configs: ", err)
	}

	influxClient, err := influx.New(
		logger,
		config.InfluxDB.Address,
		config.InfluxDB.User.Name,
		config.InfluxDB.User.Password,
		config.InfluxDB.Database.Name,
		config.InfluxDB.Database.Precision,
	)
	if err != nil {
		logger.Fatal("Failed to initialise influxdb client: ", err)
	}

	av := alphavantage.New(logger, config.AlphaVantage.ApiKey)
	for _, symbol := range config.Symbols {
		PushOverview(symbol, av, influxClient)
		PushIntra(symbol, "5min", "full", av, influxClient)
		PushBalanceSheets(symbol, av, influxClient)
		PushCashFlows(symbol, av, influxClient)
		PushEarnings(symbol, av, influxClient)
		PushIncomeStatements(symbol, av, influxClient)
	}

	err = influxClient.WriteToDB()
	if err != nil {
		logger.Fatal("Failed to write to DB: ", err)
	}
}

func PushOverview(symbol string, av alphavantage.Client, influxClient influx.Influx) {
	data, err := av.GetCompanyOverview(symbol)
	if err != nil {
		logger.Fatal("Failed to get company overview: ", err)
	}

	err = influxClient.AddToWriteBuffer(symbol, data)
	if err != nil {
		logger.Fatal("Failed to add company data to buffer: ", err)
	}
}

func PushIntra(symbol string, interval string, size string, av alphavantage.Client, influxClient influx.Influx) {
	data, err := av.GetIntra(symbol, interval, size)
	if err != nil {
		logger.Fatal("Failed to get intra day prices: ", err)
	}

	err = influxClient.AddToWriteBuffer(symbol, data)
	if err != nil {
		logger.Fatal("Failed to add intra day prices to buffer: ", err)
	}
}

func PushBalanceSheets(symbol string, av alphavantage.Client, influxClient influx.Influx) {
	data, err := av.GetBalanceSheet(symbol)
	if err != nil {
		logger.Fatal("Failed to get balance sheets: ", err)
	}

	err = influxClient.AddToWriteBuffer(symbol, data)
	if err != nil {
		logger.Fatal("Failed to add balance sheets to buffer: ", err)
	}
}

func PushCashFlows(symbol string, av alphavantage.Client, influxClient influx.Influx) {
	data, err := av.GetCashFlow(symbol)
	if err != nil {
		logger.Fatal("Failed to get cash flows: ", err)
	}

	err = influxClient.AddToWriteBuffer(symbol, data)
	if err != nil {
		logger.Fatal("Failed to add cash flows to buffer: ", err)
	}
}

func PushEarnings(symbol string, av alphavantage.Client, influxClient influx.Influx) {
	data, err := av.GetEarnings(symbol)
	if err != nil {
		logger.Fatal("Failed to get earnings: ", err)
	}

	err = influxClient.AddToWriteBuffer(symbol, data)
	if err != nil {
		logger.Fatal("Failed to add earnings to buffer: ", err)
	}
}

func PushIncomeStatements(symbol string, av alphavantage.Client, influxClient influx.Influx) {
	data, err := av.GetIncomeStatement(symbol)
	if err != nil {
		logger.Fatal("Failed to get income statements: ", err)
	}

	err = influxClient.AddToWriteBuffer(symbol, data)
	if err != nil {
		logger.Fatal("Failed to add income statements to buffer: ", err)
	}
}