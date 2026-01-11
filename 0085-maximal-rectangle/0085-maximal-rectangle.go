package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for _, row := range matrix {
		for i, val := range row {
			if val == '1' {
				heights[i]++
			} else {
				heights[i] = 0
			}
		}

		currentMax := largestRectangleArea(heights)
		if currentMax > maxArea {
			maxArea = currentMax
		}
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	for i := 0; i <= len(heights); i++ {
		var currentH int
		if i == len(heights) {
			currentH = 0
		} else {
			currentH = heights[i]
		}
		for len(stack) > 0 && currentH < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
            
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}
			
			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}