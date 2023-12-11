package print_in_order

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

var outputFilename = "output.txt"
var expectedOutput = "firstsecondthird"

// deleteFile delete the file.
func deleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	fmt.Printf("File %s deleted successfully.\n", filename)
	return nil
}

// readLinesFromFile reads the content of a file and returns it as a slice of lines.
func readLinesFromFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fileContent := string(content)

	lines := strings.Split(fileContent, "\n")

	return lines, nil
}

// writeTextToFile appends text to a file.
func writeTextToFile(text, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

// permute generates permutations of a given slice.
func permute(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permute(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

// generatePermutations generates all permutations of [1, 2, 3].
func generatePermutations() [][]int {
	var result [][]int
	nums := []int{1, 2, 3}
	permute(nums, 0, &result)
	return result
}

// printFirstFunction represents the action for printing "first".
var printFirstFunction = func() {
	defer wg.Done()
	text := "first"
	err := writeTextToFile(text, outputFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// printSecondFunction represents the action for printing "second".
var printSecondFunction = func() {
	defer wg.Done()
	text := "second"
	err := writeTextToFile(text, outputFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// printThirdFunction represents the action for printing "third".
var printThirdFunction = func() {
	defer wg.Done()
	text := "third"
	err := writeTextToFile(text, outputFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// callFunctions calls the appropriate print functions based on the given permutation.
func callFunctions(f Foo, nums []int) {
	wg.Add(3)
	for _, num := range nums {
		switch num {
		case 1:
			go f.first(printFirstFunction)
		case 2:
			go f.second(printSecondFunction)
		case 3:
			go f.third(printThirdFunction)
		}
	}
	wg.Wait()
	writeTextToFile("\n", outputFilename)
}
