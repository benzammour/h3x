// Imports and globals
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func hexdump(s string) ([]byte, error) {
	content, err := ioutil.ReadFile(s)
	if err != nil {
		panic("Cannot open file")
	}

	return content, nil
}

func print_hd(arr []byte) {
	row_width := 16
	curr_row_width := 16
	seen_bytes := 0

	for i := 0; i < len(arr); i += row_width {
		if (len(arr) - seen_bytes) < row_width { // check if our file-size is multiple of 16
			curr_row_width = (len(arr) - seen_bytes)
		}
		row := arr[i:(seen_bytes + curr_row_width)]

		// print hex values
		for i := 0; i < row_width; i++ {
			if i < curr_row_width {
				fmt.Printf("%02x ", row[i])
			} else {
				// if our file is no multiple of 16, pad it to align the "|" symbols
				fmt.Print(strings.Repeat(" ", 3))
			}
		}

		fmt.Print("|")
		fmt.Print(" ")

		// print ascii representation
		for i := 0; i < row_width; i++ {
			if i < curr_row_width {
				if row[i] >= 0x20 && row[i] < 0x7f {
					fmt.Print(string(row[i]))
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(strings.Repeat(" ", 3))
			}
		}
		fmt.Print("|")
		fmt.Print("\n")

		seen_bytes += row_width
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: h3x <FILEPATH>")
		os.Exit(1)
	}

	var content, _ = hexdump(os.Args[1])
	print_hd(content)
}
