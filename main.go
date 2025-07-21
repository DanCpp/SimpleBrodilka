package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type _map struct {
	field [][]byte
	rows  int
	cols  int
}

func ReadMap(path string) (Map _map) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var reader bufio.Reader = *bufio.NewReader(file)

	var index_line = 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		Map.field = append(Map.field, []byte(line))
		index_line += 1
		Map.cols = max(Map.cols, len(line))
	}
	Map.rows = index_line
	return Map
}

func main() {
	var Map _map = ReadMap("map.txt")

	for _, line := range Map.field {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Print("\n")
	}
}
