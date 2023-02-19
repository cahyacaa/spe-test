package main

import (
	"fmt"
	"math"
)

func NarsisticNumber(num int) bool {
	totalDigit := countDigit(num)
	var sum float64
	remainNumber := num
	var isNarsistic bool

	for remainNumber > 0 {
		sum += math.Pow(float64(remainNumber%10), float64(totalDigit))
		remainNumber = remainNumber / 10
		if num == int(sum) {
			isNarsistic = true
			break
		}
	}
	return isNarsistic
}

func countDigit(n int) int {
	if n == 0 {
		return 0

	}
	return (1 + countDigit(n/10))
}

func main() {
	fmt.Println(NarsisticNumber(153))
}
