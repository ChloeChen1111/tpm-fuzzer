package fuzz

import (
	"github.com/ChloeChen1111/tpm-fuzzer/tpm"
)

func Fuzz(data []byte) int {
	// Handle any panics to prevent the fuzzer from stopping.
	defer func() {
		if r := recover(); r != nil {
			// Panics are considered as crashes by the fuzzer.
		}
	}()

	err := tpm.ProcessInput(data)
	if err != nil {
		// Return 0 to indicate the input didn't cause a crash but wasn't interesting.
		return 0
	}

	// Return 1 to indicate the input was processed successfully and is interesting.
	return 1
}
