package tsFloat

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func Float64ToString(info float64) string {
	tmp := fmt.Sprintf("%f", info)
	return tmp
}

func Float64ToDecimal(info float64) decimal.Decimal {
	tmp := decimal.NewFromFloat(info)
	return tmp
}
