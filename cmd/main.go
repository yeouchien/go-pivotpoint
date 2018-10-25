package main

import (
	"context"
	"errors"
	"log"
	"time"

	binance "github.com/adshao/go-binance"
	"github.com/shopspring/decimal"
	pivotpoint "github.com/yeouchien/go-pivotpoint"
)

func main() {
	ohlc, err := dayOHLC()
	if err != nil {
		log.Printf("error getting day OHLC: %v", err)
		return
	}

	pp := pivotpoint.Standard(ohlc, 6)
	log.Printf("pivot point: %v", pp)
}

func dayOHLC() (pivotpoint.OHLC, error) {
	var ohlc pivotpoint.OHLC
	now := time.Now().UTC().Unix() * 1000
	client := binance.NewClient("", "")
	res, err := client.NewKlinesService().
		Symbol("QTUMBTC").
		Interval("1d").
		EndTime(now).
		Limit(1).
		Do(context.Background())
	if err != nil {
		return ohlc, err
	}

	if len(res) == 0 {
		return ohlc, errors.New("no ohlc")
	}

	dayOHLC := res[0]

	ohlc.OpenTime = time.Unix(dayOHLC.OpenTime/1000, 0).UTC()

	ohlc.Open, err = decimal.NewFromString(dayOHLC.Open)
	if err != nil {
		log.Printf("error parsing open price: %v", err)
		return ohlc, err
	}

	ohlc.High, err = decimal.NewFromString(dayOHLC.High)
	if err != nil {
		log.Printf("error parsing high price: %v", err)
		return ohlc, err
	}

	ohlc.Low, err = decimal.NewFromString(dayOHLC.Low)
	if err != nil {
		log.Printf("error parsing low price: %v", err)
		return ohlc, err
	}

	ohlc.Close, err = decimal.NewFromString(dayOHLC.Close)
	if err != nil {
		log.Printf("error parsing close price: %v", err)
		return ohlc, err
	}

	return ohlc, nil
}
