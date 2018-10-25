package pivotpoint

import (
	"github.com/shopspring/decimal"
)

// Pivot Point (P) = (High + Low + Close)/3
// Resistance 1 (R1) = (P x 2) - Low
// Resistance 2 (R2) = P + (High  -  Low)
// Resistance 3 (R3) = R2 + (High  -  Low)
// Support 1 (S1) = (P x 2) - High
// Support 2 (S2) = P  -  (High  -  Low)
// Support 3 (S3) = S2  -  (High  -  Low)
func Standard(ohlc OHLC, scale int32) PivotPoint {
	p := (ohlc.High.Add(ohlc.Low).Add(ohlc.Close)).Div(decimal.New(3, 0)).Round(scale)
	r1 := p.Mul(decimal.New(2, 0)).Sub(ohlc.Low).Round(scale)
	r2 := p.Add(ohlc.High.Sub(ohlc.Low)).Round(scale)
	r3 := r2.Add(ohlc.High.Sub(ohlc.Low)).Round(scale)
	s1 := p.Mul(decimal.New(2, 0)).Sub(ohlc.High).Round(scale)
	s2 := p.Sub(ohlc.High.Sub(ohlc.Low)).Round(scale)
	s3 := s2.Sub(ohlc.High.Sub(ohlc.Low)).Round(scale)

	return PivotPoint{p, r3, r2, r1, s1, s2, s3}
}

// Pivot Point (P) = (High + Low + Close)/3
// Resistance 1 (R1) = P + {.382 * (High  -  Low)}
// Resistance 2 (R2) = P + {.618 * (High  -  Low)}
// Resistance 3 (R3) = P + {1 * (High  -  Low)}
// Support 1 (S1) = P - {.382 * (High  -  Low)}
// Support 2 (S2) = P - {.618 * (High  -  Low)}
// Support 3 (S3) = P - {1 * (High  -  Low)}
func Fibonacci(ohlc OHLC, scale int32) PivotPoint {
	p := (ohlc.High.Add(ohlc.Low).Add(ohlc.Close)).Div(decimal.New(3, 0)).Round(scale)
	r1 := p.Add(decimal.NewFromFloat(0.382).Mul(ohlc.High.Sub(ohlc.Low))).Round(scale)
	r2 := p.Add(decimal.NewFromFloat(0.618).Mul(ohlc.High.Sub(ohlc.Low))).Round(scale)
	r3 := p.Add(decimal.New(1, 0).Mul(ohlc.High.Sub(ohlc.Low))).Round(scale)
	s1 := p.Sub((decimal.NewFromFloat(0.382).Mul(ohlc.High.Sub(ohlc.Low)))).Round(scale)
	s2 := p.Sub((decimal.NewFromFloat(0.618).Mul(ohlc.High.Sub(ohlc.Low)))).Round(scale)
	s3 := p.Sub((decimal.New(1, 0).Mul(ohlc.High.Sub(ohlc.Low)))).Round(scale)

	return PivotPoint{p, r3, r2, r1, s1, s2, s3}
}
