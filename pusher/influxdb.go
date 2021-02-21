package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/mikesupertrampster/trader/apis/services/alphavantage"
	"log"
	"reflect"
	"strings"
	"time"
)

func (p *pusher) PushIntra(symbol string, interval string, size string) {
	data, err := p.av.GetIntra(symbol, interval, size)
	if err != nil {
		log.Fatal(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  p.database.name,
		Precision: p.database.precision,
	})
	if err != nil {
		log.Fatal(err)
	}

	for date, entry := range data.TimeSeries {
		t, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			log.Fatal(err)
		}

		pt, err := client.NewPoint(
			"price",
			map[string]string{
				"symbol": data.MetaData.Symbol,
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
	iface := make(map[string]interface{})

	v := reflect.ValueOf(entry)
	for i := 0; i < v.NumField(); i++ {
		iface[strings.ToLower(v.Type().Field(i).Name)] = v.Field(i)
	}

	return iface
}