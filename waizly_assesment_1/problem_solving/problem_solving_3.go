package problem_solving

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func TimeConversion() {
	var (
		timeInput string
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	timeInput = scanner.Text()

	format12hour := "03:04:05PM"
	format24hour := "15:04:05"
	t, err := time.Parse(format12hour, timeInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t.Format(format24hour))

}
