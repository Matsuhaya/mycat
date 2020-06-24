package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n       = flag.Bool("n", false, "bool flag")
	row int = 1
)

func readFile(fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if *n {
			fmt.Printf("%v: ", row)
			row++
		}
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
