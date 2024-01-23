package asteroidscollision

import (
	"LeetcodeGO/internal/utils/ds"
	"math"
)

func AsteroidCollision(args ...interface{}) interface{} {
	var asteroids = args[0].([]int)

	return asteroidCollision(asteroids)
}

func asteroidCollision(asteroids []int) []int {

	if len(asteroids) <= 1 || allSameDirection(asteroids) || divergent(asteroids) {
		return asteroids
	}

	stack := ds.NewStack[int]()
	stack.Push(asteroids[0])
	for i := 1; i < len(asteroids); i++ {
		ast := asteroids[i]
		if stack.Length() == 0 {
			stack.Push(ast)
			continue
		}
		topElement := stack.Peek()
		if topElement*ast > 0 { //inserting same direction asteroids, no collision
			stack.Push(ast)
		} else { //different direction!
			if topElement > 0 && ast < 0 { //converging, address collision
				if math.Abs(float64(ast)) > math.Abs(float64(topElement)) {
					stack.Pop()
					stack.Push(ast)
				} else if math.Abs(float64(ast)) == math.Abs(float64(topElement)) {
					stack.Pop()
				} else if math.Abs(float64(ast)) < math.Abs(float64(topElement)) {
					//do nothing
				}
			} else if topElement < 0 && ast > 0 { //diverging, no collision
				stack.Push(ast)
			}

		}
	}

	ans := make([]int, stack.Length())

	for i := stack.Length() - 1; i >= 0; i-- {
		ans[i] = stack.Pop()
	}

	return asteroidCollision(ans)
}

func allSameDirection(asteroids []int) bool {
	sign := float64(asteroids[0]) / math.Abs(float64(asteroids[0]))

	for _, num := range asteroids {
		if num == 0 || float64(num)/math.Abs(float64(num)) != sign {
			return false
		}
	}

	return true
}

func divergent(asteroids []int) bool {
	sign := float64(asteroids[0]) / math.Abs(float64(asteroids[0]))
	if sign == 1 { //all elements to the left of the array must be directed left (negative)
		return false
	}
	var divergencyPoint int

	for i, num := range asteroids {
		if num == 0 || float64(num)/math.Abs(float64(num)) != sign {
			divergencyPoint = i
			break
		}
	}
	sign = sign * -1 //invert
	for i := divergencyPoint + 1; i < len(asteroids); i++ {
		if asteroids[i] == 0 || float64(asteroids[i])/math.Abs(float64(asteroids[i])) != sign {
			return false
		}
	}

	return true
}
