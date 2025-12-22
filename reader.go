package main

import (
	"fmt"
	"io"
)

func readStream(r io.Reader) {
	buffer := make([]byte, 8)

	for {
		n, err := r.Read(buffer)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, buffer)
		fmt.Printf("b[:n] = %q\n", buffer[:n])

		if err == io.EOF {
			break
		}
	}
}
