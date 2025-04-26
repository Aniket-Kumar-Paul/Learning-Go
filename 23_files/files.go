package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close() // <-- move defer as early as possible after successful open!

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)	
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File size:", fileInfo.Size()) // In bytes
	fmt.Println("File mode:", fileInfo.Mode()) // File permissions
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is directory:", fileInfo.IsDir())

	// Read
	buffer := make([]byte, 10)
	length, err := f.Read(buffer)
	if err != nil {
		panic(err)		
	}
	fmt.Println("Data", length, buffer)

	// Reading using readfile
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	myFileInfo, err := dir.ReadDir(-1) // n<=0 => return all the files/folders, >0 => return at most n files
	for _, fi := range myFileInfo {
		fmt.Println(fi.Name())
	}

	// create a file
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Hi Ani")
	file.WriteString("Watchu Doin?") // appends to existing text

	bytes := []byte("Hello Go")
	file.Write(bytes)


	// Read & Write from one file to another (Streaming Fashion)
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	destFile, err := os.Create("DestFile.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)
	for {
		b, errR := reader.ReadByte()
		if errR != nil {
			if errR.Error() != "EOF" {
				panic(errR)
			}
			break
		}

		errW := writer.WriteByte(b)
		if errW != nil {
			panic(errW)
		}
	}
	writer.Flush()

	// deleting file
	errors := os.Remove("example.txt")
	if errors != nil {
		panic(errors)
	}
}