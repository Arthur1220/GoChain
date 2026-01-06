package monitor

import (
	"math"
	"math/big"
)

// BytesToFloat converte o dado cru do log (big.Int) para float64 usando os decimais corretos
func BytesToFloat(data []byte, decimals int) float64 {
	// 1. Converte bytes para BigInt (Inteiro gigante)
	amountInt := new(big.Int).SetBytes(data)

	// 2. Transforma em BigFloat para precisão
	amountFloat := new(big.Float).SetInt(amountInt)

	// 3. Cria o divisor (10 elevado a decimals)
	// Ex: USDC (6) -> 1,000,000
	// Ex: ETH/DAI (18) -> 1,000,000,000,000,000,000
	divisor := new(big.Float).SetFloat64(math.Pow(10, float64(decimals)))

	// 4. Divide
	res := new(big.Float).Quo(amountFloat, divisor)

	// 5. Retorna como float64 (compatível com JSON e SQL normal)
	finalVal, _ := res.Float64()
	return finalVal
}
