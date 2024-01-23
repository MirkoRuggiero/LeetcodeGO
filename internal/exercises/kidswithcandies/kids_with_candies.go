package kidswithcandies

import (
	"math"
)

func KidsWithCandies(args ...interface{}) interface{} {

	var kids = args[0].([]int)
	var candies = args[1].(int)

	return kidsWithCandiesInternal(kids, candies)
}

func kidsWithCandiesInternal(candies []int, extraCandies int) []bool {
	var result []bool
	maximum := findMax(candies)
	for _, candy := range candies {
		result = append(result, calculateIfRicher(candy, extraCandies, maximum))
	}
	return result
}

func findMax(candies []int) int {
	maximum := -1
	for _, candy := range candies {
		maximum = int(math.Max(float64(maximum), float64(candy)))
	}
	return maximum
}

func calculateIfRicher(candy, extraCandies, max int) bool {
	return candy+extraCandies >= max
}
