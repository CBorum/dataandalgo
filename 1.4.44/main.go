package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Birthday problem. Write a program that takes an integer N from the command line and uses StdRandom.uniform()
// to generate a random sequence of integers between 0 and N – 1. Run experiments to validate the hypothesis
// that the number of integers generated before the first repeated value is found is ~√pi * N/2.

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
		upper := int(math.Pow(2, 25))

		for i := 4; i <= upper; i *= 2 {
			experiment(i)
		}
	}
}

func experiment(n int) {
	rootPIN2 := math.Sqrt(math.Pi * float64(n) / 2.0)
	m := make(map[int64]bool)

	rand.Seed(time.Now().Unix())
	for {
		num := rand.Int63n(int64(n - 1))
		if _, ok := m[num]; ok {
			fmt.Printf("n: %10d | Result: %10d | Expected: %10.2f | Accuracy: %10.2f%% \n", n, len(m), rootPIN2, (float64(len(m)) / rootPIN2) *100)
			return
		}
		m[num] = true
	}
}