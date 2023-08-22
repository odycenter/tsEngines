package tsdecimal

import "github.com/shopspring/decimal"

// 轉換 小數點轉浮點數
func DecimalToFloat64(info decimal.Decimal) float64 {
	tmp, _ := info.Float64()
	return tmp
}
