package pivotpoint

import (
	"time"

	"github.com/shopspring/decimal"
)

type PivotPoint struct {
	P  decimal.Decimal
	R3 decimal.Decimal
	R2 decimal.Decimal
	R1 decimal.Decimal
	S1 decimal.Decimal
	S2 decimal.Decimal
	S3 decimal.Decimal
}

type OHLC struct {
	OpenTime time.Time
	Open     decimal.Decimal
	High     decimal.Decimal
	Low      decimal.Decimal
	Close    decimal.Decimal
}
