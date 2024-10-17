package tpm

import "testing"

func FuzzProcessInput(f *testing.F) {
	// Seed corpus: add initial inputs if needed
	f.Add([]byte("initial input"))
	f.Fuzz(func(t *testing.T, data []byte) {
		err := ProcessInput(data)
		if err != nil {
			// Optionally log the error
			t.Logf("Error: %v", err)
		}
	})
}
