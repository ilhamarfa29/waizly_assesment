package problem_solving

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func MiniMaxSum() {
	var (
		numbersInput string
		numbers      []int64
		total        int64 = 0
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbersInput = scanner.Text()
	numbersStr := strings.Split(numbersInput, " ")

	for _, numberStr := range numbersStr {
		number, _ := strconv.ParseInt(numberStr, 10, 64)

		numbers = append(numbers, number)
	}

	maxN := slices.Max(numbers)
	minN := slices.Min(numbers)

	for _, number := range numbers {
		total = total + number
	}

	maxSum := total - minN
	minSum := total - maxN

	fmt.Println(minSum, maxSum)
}
