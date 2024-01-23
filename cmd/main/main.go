package main

import (
	"LeetcodeGO/internal/exercises/asteroidscollision"
	"LeetcodeGO/internal/exercises/equalpairs"
	"LeetcodeGO/internal/exercises/issubsequence"
	"LeetcodeGO/internal/exercises/kidswithcandies"
	"LeetcodeGO/internal/exercises/removestars"
	"fmt"
	"reflect"
)

type ExerciseFunction func(...interface{}) interface{}

type TestCase struct {
	Input          []interface{}
	ExpectedOutput interface{}
}

type Exercise struct {
	Function ExerciseFunction
	Name     string
}

func NewExercise(function ExerciseFunction, name string) *Exercise {
	return &Exercise{
		Function: function,
		Name:     name,
	}
}

func RunTests(exercise Exercise, testCases []TestCase) {
	fmt.Println(pad(fmt.Sprintf("-----------------------Running %s", exercise.Name)))
	for i, testCase := range testCases {
		result := exercise.Function(testCase.Input...)

		fmt.Printf("Test case %d: %v\n", i, testCase.Input)
		if reflect.DeepEqual(result, testCase.ExpectedOutput) {
			fmt.Printf("Passed!\n Got %v as result\n\n", result)
		} else {
			fmt.Printf("Failed!\n")
			fmt.Printf("  Expected Output: %v\n", testCase.ExpectedOutput)
			fmt.Printf("  Actual Output: %v\n", result)
		}
	}
}

func main() {

	testCases := []TestCase{
		{Input: []interface{}{[]int{2, 3, 5, 1, 3}, 3}, ExpectedOutput: []bool{true, true, true, false, true}},
	}

	exercise := NewExercise(kidswithcandies.KidsWithCandies, "KidsWithCandies")
	RunTests(*exercise, testCases)

	//===========================================================================================

	testCases = []TestCase{
		{Input: []interface{}{"abc", "ahbgdc"}, ExpectedOutput: true},
		{Input: []interface{}{"axc", "ahbgdc"}, ExpectedOutput: false},
		{Input: []interface{}{"ab", "baab"}, ExpectedOutput: true},
	}

	exercise = NewExercise(issubsequence.IsSubsequence, "IsSubsequence")
	RunTests(*exercise, testCases)

	//===========================================================================================

	testCases = []TestCase{
		{Input: []interface{}{[][]int{{3, 1, 2, 2}, {1, 4, 4, 4}, {2, 4, 2, 2}, {2, 5, 2, 2}}}, ExpectedOutput: 3},
		{Input: []interface{}{[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}}, ExpectedOutput: 1},
		{Input: []interface{}{[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}}, ExpectedOutput: 3},
		{Input: []interface{}{[][]int{{13, 13}, {13, 13}}}, ExpectedOutput: 4},
	}

	exercise = NewExercise(equalpairs.EqualPairs, "EqualPairs")
	RunTests(*exercise, testCases)

	//===========================================================================================

	testCases = []TestCase{
		{Input: []interface{}{"leet**cod*e"}, ExpectedOutput: "lecoe"},
		{Input: []interface{}{"erase*****"}, ExpectedOutput: ""},
	}

	exercise = NewExercise(removestars.RemoveStars, "RemoveStars")
	RunTests(*exercise, testCases)

	//===========================================================================================

	testCases = []TestCase{
		{Input: []interface{}{[]int{5, 10, -5}}, ExpectedOutput: []int{5, 10}},
		{Input: []interface{}{[]int{8, -8}}, ExpectedOutput: []int{}},
		{Input: []interface{}{[]int{10, 2, -5}}, ExpectedOutput: []int{10}},
		{Input: []interface{}{[]int{-2, -1, 1, 2}}, ExpectedOutput: []int{-2, -1, 1, 2}},
		{Input: []interface{}{[]int{1, -1, -2, -2}}, ExpectedOutput: []int{-2, -2}},
	}

	exercise = NewExercise(asteroidscollision.AsteroidCollision, "AsteroidCollision")
	RunTests(*exercise, testCases)
}

func pad(input string) string {
	a := 90 - len(input)
	var sb string
	for a > 0 {
		sb += "-"
		a--
	}
	return input + sb
}
