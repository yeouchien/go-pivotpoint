package pivotpoint

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
)

var (
	// This Comparer only operates on decimal.Decimal.
	// It uses decimal.Equal to compare
	cmpOpt = cmp.Comparer(func(x, y decimal.Decimal) bool {
		return x.Equal(y)
	})
)

func TestStandard(t *testing.T) {
	ohlc := OHLC{
		OpenTime: time.Date(
			2018, time.October, 15, 0, 0, 0, 0, time.UTC,
		),
		Open:  decimal.NewFromFloat(0.000546),
		High:  decimal.NewFromFloat(0.000688),
		Low:   decimal.NewFromFloat(0.000523),
		Close: decimal.NewFromFloat(0.000650),
	}

	expectedPP := PivotPoint{
		P:  decimal.NewFromFloat(0.000620),
		R3: decimal.NewFromFloat(0.000950),
		R2: decimal.NewFromFloat(0.000785),
		R1: decimal.NewFromFloat(0.000717),
		S1: decimal.NewFromFloat(0.000552),
		S2: decimal.NewFromFloat(0.000455),
		S3: decimal.NewFromFloat(0.000290),
	}

	resultPP := Standard(ohlc, 6)

	if !cmp.Equal(expectedPP, resultPP, cmpOpt) {
		t.Fatalf("expected pivot point to be %v but got %v", expectedPP, resultPP)
	}
}

func TestFibonacci(t *testing.T) {
	ohlc := OHLC{
		OpenTime: time.Date(
			2018, time.October, 15, 0, 0, 0, 0, time.UTC,
		),
		Open:  decimal.NewFromFloat(0.000546),
		High:  decimal.NewFromFloat(0.000688),
		Low:   decimal.NewFromFloat(0.000523),
		Close: decimal.NewFromFloat(0.000650),
	}

	expectedPP := PivotPoint{
		P:  decimal.NewFromFloat(0.000620),
		R3: decimal.NewFromFloat(0.000785),
		R2: decimal.NewFromFloat(0.000722),
		R1: decimal.NewFromFloat(0.000683),
		S1: decimal.NewFromFloat(0.000557),
		S2: decimal.NewFromFloat(0.000518),
		S3: decimal.NewFromFloat(0.000455),
	}

	resultPP := Fibonacci(ohlc, 6)

	if !cmp.Equal(expectedPP, resultPP, cmpOpt) {
		t.Fatalf("expected pivot point to be %v but got %v", expectedPP, resultPP)
	}
}
