package influx

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/mikesupertrampster/trader/apis/services/alphavantage"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type Influx struct {
	BatchPoints *client.BatchPoints
	Client       client.Client
}

func New(logger *logrus.Logger, addr string, username string, password string, database string, precision string) (Influx, error) {
	var influx Influx
	log = logger

	httpClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Error("Could not create new http client: ", err)
		return influx, err
	}
	influx.Client = httpClient

	defer func(c client.Client) {
		err := c.Close()
		if err != nil {
			log.Error("Failed to close influxdb client: ", err)
		}
	}(httpClient)

	batchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: precision,
	})
	if err != nil {
		log.Fatal("Failed to create new batch-points: ",err)
		return influx, err
	}
	influx.BatchPoints = &batchPoints

	return influx, err
}

func (c *Influx) AddToWriteBuffer(symbol string, series alphavantage.Series) error {
	for _, datapoint := range series {
		pt, err := client.NewPoint(
			datapoint.Name,
			map[string]string{
				"symbol": symbol,
			},
			datapoint.Data,
			datapoint.Timestamp)
		if err != nil {
			log.Fatal("Failed to add data to batch point", err)
			return err
		}

		(*c.BatchPoints).AddPoint(pt)
	}

	return nil
}

func (c *Influx) WriteToDB() error {
	return c.Client.Write(*c.BatchPoints)
}
