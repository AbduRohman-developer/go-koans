package go_koans

func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

func findPrimeNumbers(channel chan int) {
	for i := 2; ; /* infinite loop */ i++ {
		if i == 2 || i%2 != 0 {
			channel <- i
		}
		assert(i < 100) // i is afraid of heights
	}
}

func aboutConcurrency() {
	ch := make(chan int)

	//assert(isPrimeNumber(5)) // concurrency can be almost trivial
	go findPrimeNumbers(ch)

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 9)
	assert(<-ch == 11)
}
