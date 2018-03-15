package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Coupon collector problem. Generating random integers as in the previous exercise,
// run experiments to validate the hypothesis that the number of integers generated
// before all possible values are generated is ~N HN.

func main() {
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}

		if n < 1 {
			panic("non positive n")
		}
		experiment(n)
	} else {
		upper := int(math.Pow(2, 18))

		for i := 4; i <= upper; i *= 2 {
			experiment(i)
		}
	}
}

func experiment(n int) {
	nhn := float64(n) * harmonicNumbersSum(n)
	//natural logarithm
	// mln := ln(float64(n))
	// fmt.Println(mln)

	m := make(map[int64]bool)
	count := 0

	rand.Seed(time.Now().Unix())
	for {
		num := rand.Int63n(int64(n - 1))
		count++
		if _, ok := m[num]; !ok {
			m[num] = true
		}
		if len(m) == n-1 {
			fmt.Printf("n: %10d | Result: %10d | Expected: %10.2f | Accuracy: %10.2f%% \n", n, count, nhn, (float64(count) / nhn) *100)
			return
		}
	}
}

func harmonicNumbersSum(n int) float64 {
	sum := 0.0
	nrange := float64(n)
	for i := 1.0; i <= nrange; i++ {
		sum += 1 / i
	}
	return sum
}

// https://gist.github.com/thinkphp/ae5024dbd0ea6b83b6308a028ea22323
func f(x, a float64) float64 {
	return math.Exp(x) - a
}

func ln(n float64) float64 {
	var lo, hi, m float64
	if n <= 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	EPS := 0.00001
	lo = 0
	hi = n
	for math.Abs(lo-hi) >= EPS {
		m = float64((lo + hi) / 2.0)
		if f(m, n) < 0 {
			lo = m
		} else {
			hi = m
		}
	}
	return float64((lo + hi) / 2.0)
}
