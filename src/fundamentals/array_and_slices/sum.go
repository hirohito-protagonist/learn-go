package main

func Sum(n [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += n[i]
	}
	return sum
}
