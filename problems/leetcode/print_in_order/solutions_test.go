package print_in_order

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func Test_AllSolutions(t *testing.T) {
	permutations := generatePermutations()
	for _, nums := range permutations {
		f := []Foo{NewFooMutex(), NewFooCond(), NewFooChannel()}
		for i := 0; i < len(f); i++ {
			callFunctions(f[i], nums)
		}
	}
	lines, err := readLinesFromFile(outputFilename)
	assert.Equal(t, err, nil)
	for i := 0; i < len(lines)-1; i++ {
		assert.Equal(t, lines[i], expectedOutput)
	}
	deleteFile(outputFilename)
}
