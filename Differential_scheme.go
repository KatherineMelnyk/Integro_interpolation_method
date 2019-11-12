package main

import (
	"fmt"

	"math"

	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/mat"
)

func countIntegral(low, up float64, F func(float64) float64) float64 {
	x := seqX(50, low, up)
	y := evaluatePoints(F, x)
	res := integrate.Simpsons(x, y)
	return res
}

func d0(x0, h float64) float64 {
	res := countIntegral(x0, x0+h/2, q)
	return (2 / h) * res
}

func dN(xN, h float64) float64 {
	res := countIntegral(xN-h/2, xN, q)
	return (2 / h) * res
}

func d_(xi, h float64) float64 {
	res := countIntegral(xi-h/2, xi+h/2, q)
	return (1 / h) * res
}

func phi0(x0, h float64) float64 {
	res := countIntegral(x0, x0+h/2, f)
	return (2 / h) * res
}

func phiN(xN, h float64) float64 {
	res := countIntegral(xN-h/2, xN, f)
	return (2 / h) * res
}

func k1(x float64) float64 {
	return 1 / k(x)
}

func a(xi, h float64) float64 {
	res := countIntegral(xi-h, xi, k1)
	return h / res // 1 / ((1 / h) * res)
}

//elements[i][i-1] = (-p(x[i]) / 2 * h) - a(x[i+1], h)/h_2
//elements[i][i] = (1/h_2)*(a(x[i+1], h)-a(x[i], h)) + d_(x[i], h)
//elements[i][i+1] = (p(x[i]) / 2 * h) - a(x[i+1], h)*(1/h_2)

func scheme(x []float64, h float64) *mat.Dense {
	N := len(x)
	h_2 := math.Pow(h, 2)
	coef := mat.NewDense(N, N, nil)
	coef.Set(0, 0, ALPHA1+(a(x[1], h)/h)+(d0(x[0], h)*(h/2)))
	coef.Set(0, 1, -a(x[1], h)/h)
	coef.Set(N-1, N-2, -a(x[N-1], h)/h)
	coef.Set(N-1, N-1, ALPHA2+(a(x[N-1], h)/h)+(dN(x[N-1], h)*(h/2)))
	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			if i == j {
				coef.Set(i, j-1, -a(x[i], h)/h_2)
				coef.Set(i, j, (a(x[i+1], h)+a(x[i], h))/h_2+d_(x[i], h))
				coef.Set(i, j+1, -a(x[i+1], h)/h_2)
			}
		}
	}
	return coef
}

func phi(x []float64, h float64) *mat.Dense {
	N := len(x)
	val := mat.NewDense(N, 1, nil)
	val.Set(0, 0, mu1(x[0])+(h/2)*phi0(x[0], h))
	val.Set(N-1, 0, mu2(x[N-1])+(h/2)*phiN(x[N-1], h))
	for i := 1; i < N-1; i++ {
		val.Set(i, 0, countIntegral(x[i]-h/2, x[i]+h/2, f)*(1/h))
	}
	return val
}

func solution(x []float64, h float64) []float64 {
	C := scheme(x, h)
	F := phi(x, h)
	var Res mat.Dense
	Res.Solve(C, F)

	fmt.Printf("Cond: %.5f\n", mat.Cond(C, 2))
	res := make([]float64, len(x))
	for i := 0; i < len(res); i++ {
		res[i] = Res.RawRowView(i)[0]
	}

	c := mat.NewVecDense(len(x), nil)
	c.MulVec(C, Res.ColView(0))
	//fmt.Print(c.RawVector().Data)
	fmt.Print("\n")
	for i := 0; i < len(x); i++ {
		fmt.Printf("%.3f \t", c.RawVector().Data[i]-F.RawRowView(i)[0])
	}
	fmt.Print("\n")
	return res
}
