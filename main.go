package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := [7]float64{0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8}
	pLoc := 0
	var weight float64
	var trials int
	var n int
	var bob int
	var alice int
	var bobwins int
	var alicewins int

	fmt.Println("Enter the number of coin tosses: ")
	fmt.Scanln(&n)
	fmt.Println("Enter the number of trials: ")
	fmt.Scanln(&trials)

	for i := 0; i < len(p); i++ {
		weight = p[pLoc]
		pLoc++

		trial(p, pLoc, trials, n, bob, alice, bobwins, alicewins, weight)
	}

}

func coinToss(n int, weight float64) (int, int) {
	temp := 0.0
	bob := 0
	alice := 0
	b := rand.New(rand.NewSource(time.Now().UnixNano()))
	sleep := time.Duration(1)
	time.Sleep(sleep)
	a := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Bob's coin tosses
	for j := 0; j < n + 1; j++ {
		temp = b.Float64()
		if temp < weight {
			bob++
		}
	}

	// Alice's coin tosses
	for k := 0; k < n; k++ {
		temp = a.Float64()
		if temp < weight {
			alice++
		}
	}

	return bob, alice
}

func trial(p [7]float64, pLoc int, trials int, n int, bob int, alice int, bobwins int, alicewins int, weight float64) {
	bobwins = 0
	alicewins = 0
	for j := 0; j < trials; j++ {
		bob, alice = coinToss(n, weight)
		if bob > alice {
			bobwins++
		} else if alice > bob {
			alicewins++
		}
	}

	relFreq := float64(bobwins) / float64(trials)

	// refactor print statements into single statement using fstring
	fmt.Printf("---------------------------\nWeight: %v\nTotal trials: %v\nBob wins: %v\nAlice wins: %v\nRelative frequency: %v\n",
		weight, trials, bobwins, alicewins, relFreq)
}
