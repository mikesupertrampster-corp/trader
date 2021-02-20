package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/mikesupertrampster/trader/apis/services/alphavantage"
	"github.com/vrischmann/envconfig"
	"log"
	"reflect"
	"strings"
	"time"
)

type pusher struct {
	client client.Client
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
		ApiKey string `envconfig:"default=KEY"`
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
	p := pusher{client: c}

	av := alphavantage.New("demo")
	daily, err := av.GetDaily("IBM")
	if err != nil {
		log.Fatal(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.InfluxDB.Database.Name,
		Precision: config.InfluxDB.Database.Precision,
	})
	if err != nil {
		log.Fatalf("errhere : %v",err)
	}

	p.push(daily, bp)
}

func (p *pusher) push(daily alphavantage.Daily, bp client.BatchPoints) {
	for date, entry := range daily.TimeSeries {
		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Fatal(err)
		}

		pt, err := client.NewPoint(
			"price",
			map[string]string{
				"symbol": daily.MetaData.Symbol,
			},
			toIf(*entry),
			t)
		if err != nil {
			log.Fatal(err)
		}

		bp.AddPoint(pt)
	}

	if err := p.client.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func influxDBClient(addr string, username string, password string) client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func toIf(entry alphavantage.Entry) map[string]interface{} {
	in := make(map[string]interface{})

	v := reflect.ValueOf(entry)
	for i := 0; i < v.NumField(); i++ {
		in[strings.ToLower(v.Type().Field(i).Name)] = v.Field(i)
	}

	return in
}