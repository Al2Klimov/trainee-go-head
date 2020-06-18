// head | (c) 2020 NETWAYS GmbH | GPLv2+

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var linesParameter = flag.Int("n", 10, "print NUM lines instead of 10.")

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		for filePOS, fileName := range flag.Args() {
			file, fileErr := os.Open(fileName)
			if fileErr != nil {
				fmt.Fprintln(os.Stderr, fileErr)
				os.Exit(1)
			}

			countAndPrint(file)

			if filePOS != len(flag.Args())-1 {
				fmt.Print("\n")
			}

			file.Close()
		}
	} else {
		countAndPrint(os.Stdin)
	}
}

func countAndPrint(input *os.File) {
	if len(flag.Args()) > 1 {
		fmt.Printf("==> %s <==\n", input.Name())
	}

	buf := bufio.NewReader(input)
	var counter int

	for {
		if counter == *linesParameter {
			break
		}

		data, dataErr := buf.ReadBytes('\n')
		if dataErr != nil && dataErr != io.EOF {
			fmt.Fprintln(os.Stderr, dataErr)
			os.Exit(1)
		}

		_, _ = os.Stdout.Write(data)
		counter++

		if dataErr == io.EOF {
			break
		}
	}
}
