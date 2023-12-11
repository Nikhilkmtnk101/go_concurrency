package print_in_order

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func Test_MutexSolution(t *testing.T) {
	permutations := generatePermutations()
	for _, nums := range permutations {
		f := NewFooMutex()
		callFunctions(f, nums)
	}
	lines, err := readLinesFromFile(outputFilename)
	assert.Equal(t, err, nil)
	for i := 0; i < len(lines)-1; i++ {
		assert.Equal(t, lines[i], "firstsecondthird")
	}
	deleteFile(outputFilename)
}
