package utils

import "fmt"

func updateBarDL(msg string, percentage int, total int64, size int64) {
	length := 30

	numChars := percentage * length / 100

	fmt.Printf("\r%s [", msg)

	for i := 0; i < numChars; i++ {
		fmt.Printf("#")
	}

	for i := 0; i < length-numChars; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("] %d%% Done (%v / %v MB)", percentage, total/1000000, size/1000000)
}
