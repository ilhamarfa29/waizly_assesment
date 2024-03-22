package problem_solving

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PlusMinus() {
	var (
		arraySize          int64
		arrNInput          string
		countPostNumber    int64 = 0
		countNegtNumber    int64 = 0
		countZeroNumber    int64 = 0
		positiveProportion float64
		negativeProportion float64
		zeroProportion     float64
	)

	fmt.Scanln(&arraySize)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrNInput = scanner.Text()
	arrNStr := strings.Split(arrNInput, " ")

	for _, arrNItem := range arrNStr {
		arrItem, _ := strconv.ParseInt(arrNItem, 10, 64)

		if arrItem > 0 {
			countPostNumber++
		} else if arrItem < 0 {
			countNegtNumber++
		} else {
			countZeroNumber++
		}
	}

	positiveProportion = float64(countPostNumber) / float64(arraySize)
	negativeProportion = float64(countNegtNumber) / float64(arraySize)
	zeroProportion = float64(countZeroNumber) / float64(arraySize)

	fmt.Printf("%.6f\n", positiveProportion)
	fmt.Printf("%.6f\n", negativeProportion)
	fmt.Printf("%.6f\n", zeroProportion)
}
