package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: regiontool <filename>")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inRegion := false
	//regionName := ""

	for scanner.Scan() {
		line := scanner.Text()
		if inRegion {
			if line == "// #endregion" {
				inRegion = false
			} else {
				fmt.Println(line)
			}
		} else {
			if line[:10] == "// #region" {
				inRegion = true
				//regionName = line[11:]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
