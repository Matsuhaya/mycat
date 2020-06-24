package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func readFile(fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	fileNames := flag.Args()

	for _, fn := range fileNames {
		err := readFile(fn)
		if err != nil { /* エラー処理 */
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
}
