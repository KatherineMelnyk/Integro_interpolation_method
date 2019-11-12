package main

import (
	"fmt"
	"math"
)

const A = 1.
const B = 3.
const ALPHA1, ALPHA2 = 1., 2.
const Q1, Q2, Q3 = 3., 2., 1.
const K1, K2, K3 = 3., 3., 2.
const P1, P2, P3 = 1., 1., 1.
const M1, M2, M3 = 3., 2., 0.

func main() {
	N := 10
	X, h := sequenceOfX(N)
	U := make([]float64, len(X))
	fmt.Printf("The step : %.4f \n", h)
	for i := 0; i < len(X); i++ {
		U[i] = u(X[i])
	}
	Y := solution(X, h)
	for i := 0; i < len(X); i++ {
		fmt.Printf("X[i]: %.4f \t Y[i]: %.4f \t U[i]: %.4f \t delta[i]: %.4f \t", X[i], Y[i], U[i], math.Abs(Y[i]-U[i]))
		fmt.Print("\n")
	}
}
