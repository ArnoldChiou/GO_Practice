package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	//create a file
	userFile := "arrnoldc.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()

	//write to the file
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()

	// Read from the file and write to stdout
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
	}
	//fout.Close()
	//fl.Close()

	time.Sleep(1000 * time.Millisecond)
	if err := os.Remove(userFile); err != nil {
		fmt.Println("Remove error:", err)
	}
}
