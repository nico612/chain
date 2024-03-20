package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
)

// ParseEther parse ether, amount is string, float64, int64, decimal.Decimal, *decimal.Decimal
func ParseEther(amount interface{}) *big.Int {
	return ParseUnits(amount, 18)
}

// ParseUnits decimals to wei
func ParseUnits(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case int:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	// 10^decimals
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))

	// amount * 10^decimals
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// FormatEther format ether, amount is string, *big.Int
func FormatEther(amount interface{}) decimal.Decimal {
	return FormatUnits(amount, 18)
}

// FormatUnits wei to decimals
// ivalue type: string, *big.Int
func FormatUnits(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}
	// 10^decimals
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())

	// value / 10^decimals
	result := num.Div(mul)

	return result
}
