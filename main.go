// head | (c) 2020 NETWAYS GmbH | GPLv2+

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var counter int

	for {
		if counter == 10 {
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
