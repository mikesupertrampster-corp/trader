package main

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/mikesupertrampster/trader/apis/services/alphavantage"
	"reflect"
	"strings"
	"time"
)

func (p *pusher) PushOverview(symbol string) error {
	data, err := p.av.GetCompanyOverview(symbol)
	if err != nil {
		return err
	}

	return p.pushes(symbol, "overview", data)
}

func (p *pusher) PushIntra(symbol string, interval string, size string) error {
	data, err := p.av.GetIntra(symbol, interval, size)
	if err != nil {
		return err
	}

	return p.pushes(symbol, "price", data)
}

func (p *pusher) PushBalance(symbol string) error {
	data, err := p.av.GetBalanceSheet(symbol)
	if err != nil {
		return err
	}

	return p.pushes(symbol, "balance", data)
}

func (p *pusher) PushCash(symbol string) error {
	data, err := p.av.CashFlow(symbol)
	if err != nil {
		return err
	}

	return p.pushes(symbol, "cash", data)
}

func (p *pusher) PushEarnings(symbol string) error {
	data, err := p.av.GetEarnings(symbol)
	if err != nil {
		return err
	}

	return p.pushes(symbol, "earnings", data)
}

func (p *pusher) PushIncome(symbol string) error {
	data, err := p.av.GetIncomeStatement(symbol)
	if err != nil {
		return err
	}

	return p.pushes(symbol, "income", data)
}

func (p *pusher) pushes(symbol string, name string, data interface{}) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  p.database.name,
		Precision: p.database.precision,
	})
	if err != nil {
		return err
	}

	switch data.(type) {
	case alphavantage.Intra:
		for date, entry := range data.(alphavantage.Intra).TimeSeries {
			t, err := time.Parse("2006-01-02 15:04:05", date)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				name,
				map[string]string{
					"symbol": symbol,
				},
				toIf(*entry),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}
	case alphavantage.BalanceSheet:
		for _, report := range data.(alphavantage.BalanceSheet).AnnualReports {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_annual", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}

		for _, report := range data.(alphavantage.BalanceSheet).QuarterlyReports {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_quarterly", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}
	case alphavantage.CashFlow:
		for _, report := range data.(alphavantage.CashFlow).AnnualReports {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_annual", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}

		for _, report := range data.(alphavantage.CashFlow).QuarterlyReports {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_quarterly", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}
	case alphavantage.Earnings:
		for _, report := range data.(alphavantage.Earnings).AnnualEarnings {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_annual", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}

		for _, report := range data.(alphavantage.Earnings).QuarterlyEarnings {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_quarterly", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}
	case alphavantage.IncomeStatement:
		for _, report := range data.(alphavantage.IncomeStatement).AnnualReports {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_annual", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}

		for _, report := range data.(alphavantage.IncomeStatement).QuarterlyReports {
			t, err := time.Parse("2006-01-02", report.FiscalDateEnding)
			if err != nil {
				return err
			}

			pt, err := client.NewPoint(
				fmt.Sprintf("%s_quarterly", name),
				map[string]string{
					"symbol": symbol,
				},
				toIf(report),
				t)
			if err != nil {
				return err
			}

			bp.AddPoint(pt)
		}
	case alphavantage.CompanyOverview:
		overview := data.(alphavantage.CompanyOverview)
		t, err := time.Parse("2006-01-02", overview.LatestQuarter)
		if err != nil {
			return err
		}

		pt, err := client.NewPoint(
			name,
			map[string]string{
				"symbol": symbol,
			},
			toIf(overview),
			t)
		if err != nil {
			return err
		}

		bp.AddPoint(pt)
	}

	if err := p.client.Write(bp); err != nil {
		return err
	}

	return nil
}

func influxDBClient(addr string, username string, password string) (client.Client, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		return c, err
	}

	return c, nil
}

func toIf(i interface{}) map[string]interface{} {
	iface := make(map[string]interface{})

	v := reflect.ValueOf(i)
	for i := 0; i < v.NumField(); i++ {
		iface[strings.ToLower(v.Type().Field(i).Name)] = v.Field(i)
	}

	return iface
}
