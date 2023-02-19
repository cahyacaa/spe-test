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

func PartyOutliers(arr []int) []int {
	evenArr := make([]int, 0, 0)
	oddArr := make([]int, 0, 0)
	for _, num := range arr {
		if num%2 == 0 {
			evenArr = append(evenArr, num)
		} else {
			oddArr = append(oddArr, num)
		}
	}

	//cannot false because go is need explisit return
	if len(evenArr) == 0 || len(oddArr) == 0 {
		return []int{}
	}

	if len(evenArr) == 1 {
		return evenArr
	} else {
		return oddArr
	}
}

func NeedleHaystack(arr []string, query string) int {
	for i, val := range arr {
		if val == query {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(NarsisticNumber(153))
	fmt.Println(PartyOutliers([]int{160, 3, 1719, 19, 11, 13, -21}))
	fmt.Println(NeedleHaystack([]string{"red", "blue", "yellow"}, "blue"))
}
