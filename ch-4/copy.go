package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run copy.go file1 file2")
		os.Exit(1)
	}
	copy(os.Args[1], os.Args[2])
}

func copy(src, dst string) {
	sf, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer sf.Close()

	of, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	buf := make([]byte, 1)
	for {
		_, readErr := sf.Read(buf)
		if readErr != nil && readErr != io.EOF {
			log.Fatal(err)
		}

		if _, err := of.Write(buf); err != nil {
			log.Fatal(err)
		}
		if readErr == io.EOF {
			break
		}
	}
}
