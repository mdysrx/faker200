package cmd

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err == nil {
		return
	}

	exit(err.Error())
}

func exit(msg string) {
	fmt.Println("error:", msg)
	os.Exit(1)
}

// func strToInt(s string) (int, error) {
// 	i, err := strconv.Atoi(s)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	return i, nil
// }
