package main

import (
	"fmt"

	"math"

	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/mat"
)

const integralN = 50

func d0(x0, h float64) float64 {
	x := seqX(integralN, x0, x0+h/2)
	y := evaluatePoints(q, x)
	res := integrate.Simpsons(x, y)
	return (2 / h) * res
}

func dN(xN, h float64) float64 {
	x := seqX(integralN, xN-h/2, xN)
	y := evaluatePoints(q, x)
	res := integrate.Simpsons(x, y)
	return (2 / h) * res
}

func d_(xi, h float64) float64 {
	x := seqX(integralN, xi-h/2, xi+h/2)
	y := evaluatePoints(q, x)
	res := integrate.Simpsons(x, y)
	return (1 / h) * res
}

func phi0(x0, h float64) float64 {
	x := seqX(integralN, x0, x0+h/2)
	y := evaluatePoints(f, x)
	res := integrate.Simpsons(x, y)
	return (2 / h) * res
}

func phiN(xN, h float64) float64 {
	values := seqX(integralN, xN-h/2, xN)
	y := evaluatePoints(f, values)
	res := integrate.Simpsons(values, y)
	return (2 / h) * res
}

func k1(x float64) float64 {
	return 1 / k(x)
}

func a(xi, h float64) float64 {
	x := seqX(integralN, xi-h, xi)
	y := evaluatePoints(k1, x)
	res := integrate.Simpsons(x, y)
	return h / res // 1 / ((1 / h) * res)
}

//func scheme(x []float64, h float64) [][]float64 {
//	elements := matrix(len(x), len(x))
//	N := len(x)
//	h_2 := math.Pow(h, 2)
//	elements[0][0] = ALPHA1 + (a(x[1], h) / h) + (d0(x[0], h) * (h / 2))
//	elements[0][1] = -a(x[1], h) / h
//	elements[N-1][N-2] = -a(x[N-1], h) / h
//	elements[N-1][N-1] = ALPHA2 + (a(x[N-1], h) / h) - (dN(x[N-1], h) * (h / 2))
//	for i := 1; i < len(x)-1; i++ {
//		//elements[i][i-1] = (-p(x[i]) / 2 * h) - a(x[i+1], h)/h_2
//		//elements[i][i] = (1/h_2)*(a(x[i+1], h)-a(x[i], h)) + d_(x[i], h)
//		//elements[i][i+1] = (p(x[i]) / 2 * h) - a(x[i+1], h)*(1/h_2)
//
//		elements[i][i-1] = -a(x[i], h) / h_2
//		elements[i][i] = (1/h_2)*(a(x[i+1], h)+a(x[i], h)) + d_(x[i], h)
//		elements[i][i+1] = -a(x[i+1], h) / h_2
//	}
//	return elements
//}

func scheme(x []float64, h float64) *mat.Dense {
	N := len(x)
	h_2 := math.Pow(h, 2)
	coef := mat.NewDense(N, N, nil)
	coef.Set(0, 0, ALPHA1+(a(x[1], h)/h)+(d0(x[0], h)*(h/2)))
	coef.Set(0, 1, -a(x[1], h)/h)
	coef.Set(N-1, N-2, -a(x[N-1], h)/h)
	coef.Set(N-1, N-1, ALPHA2+(a(x[N-1], h)/h)-(dN(x[N-1], h)*(h/2)))
	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			if i == j {
				coef.Set(i, j-1, -a(x[i], h)/h_2)
				coef.Set(i, j, (1/h_2)*(a(x[i+1], h)+a(x[i], h))+d_(x[i], h))
				coef.Set(i, j+1, -a(x[i+1], h)/h_2)
			}
		}
	}
	return coef
}

func phi(x []float64, h float64) []float64 {
	values := make([]float64, len(x))
	N := len(x)
	values[0] = mu1(x[0]) + (h/2)*phi0(x[0], h)
	values[N-1] = mu2(x[N-1]) - (h/2)*phiN(x[N-1], h)
	for i := 1; i < len(x)-1; i++ {
		xlow := x[i] - h/2
		xup := x[i] + h/2
		val := seqX(20, xlow, xup)
		y := evaluatePoints(f, val)
		res := integrate.Simpsons(val, y)
		values[i] = res * (1 / h)
	}
	return values
}

func solution(x []float64, h float64) []float64 {
	coef := scheme(x, h)
	//C := FromMattoVec(coef)
	phi := phi(x, h)
	F := mat.NewDense(len(phi), 1, phi)
	//E := mat.NewDense(len(coef), len(coef[0]), C)
	var Res mat.Dense
	//Res.Solve(E, F)
	Res.Solve(coef, F)
	//fmt.Print(E.RawMatrix().Data)
	fmt.Printf("Cond: %.5f\n", mat.Cond(coef, 2))
	res := make([]float64, len(phi))
	for i := 0; i < len(res); i++ {
		res[i] = Res.RawRowView(i)[0]
	}
	return res
}
