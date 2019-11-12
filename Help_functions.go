package main

import (
	"fmt"
)

func evaluatePoints(g func(float64) float64, x []float64) []float64 {
	y := make([]float64, len(x))
	for i, element := range x {
		y[i] = g(element)
	}
	return y
}

func matrix(c, r int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < r; i++ {
		matrix = append(matrix, make([]float64, c))
	}
	return matrix
}

func FromMattoVec(matrix [][]float64) []float64 {
	Size := len(matrix) * len(matrix[0])
	vector := make([]float64, Size)
	for i := 0; i < len(matrix); i++ {
		copy(vector[i*len(matrix[i]):(i+1)*len(matrix[i])], matrix[i])
	}
	return vector
}

func sequenceOfX(N int) ([]float64, float64) {
	sequence := make([]float64, N+1)
	h := (B - A) / float64(N)
	for i := 0; i < len(sequence); i++ {
		sequence[i] = A + h*float64(i)
	}
	return sequence, h
}

func printValues(v []func(float64, int) float64, x float64) {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%.8f \t", v[i](x, i))
	}
}

func printMatrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Printf("%.4f \t", m[i][j])
		}
		fmt.Print("\n")
	}
}

func printVector(v []float64) {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%.4f \t", v[i])
	}
	fmt.Print("\n")
}

func seqX(N int, a, b float64) []float64 {
	sequence := make([]float64, N+1)
	h := (b - a) / float64(N)
	for i := 0; i < len(sequence); i++ {
		sequence[i] = a + h*float64(i)
	}
	sequence[N] = b
	return sequence
}
