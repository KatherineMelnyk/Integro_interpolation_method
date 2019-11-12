package main

import "math"

func u(x float64) float64 {
	return M1*math.Sin(M2*x) + M3
}

func du(x float64) float64 {
	return M1 * M2 * math.Cos(M2*x)
}

func d2u(x float64) float64 {
	return M1 * math.Pow(M2, 2) * (-1) * math.Sin(x*M2)
}

func mu1(x float64) float64 {
	return -k(x)*du(x) + ALPHA1*u(x)
}

func mu2(x float64) float64 {
	return k(x)*du(x) + ALPHA2*u(x)
}

func k(x float64) float64 {
	return K1*math.Pow(x, K2) + K3
}

func dk(x float64) float64 {
	return K1 * K2 * (math.Pow(x, K2-1))
}

func p(x float64) float64 {
	return P1*math.Pow(x, P2) + P3
}

func q(x float64) float64 {
	return Q1*math.Pow(x, Q2) + Q3
}

//func f(x float64) float64 {
//	return (-1)*(dk(x)*du(x)+k(x)*d2u(x)) + p(x)*du(x) + q(x)*u(x)
//}

func f(x float64) float64 {
	return (-1)*(dk(x)*du(x)+k(x)*d2u(x)) + q(x)*u(x)
}
