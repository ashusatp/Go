package main

import (
	"bufio"
	"os"
)

func main() {

	// -> Read file using Buffer

	/*
		// open file Open()
		f, err := os.Open("example.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// get FileInfo using Stat()
		fileInfo, err := f.Stat()
		if err != nil {
			panic(err)
		}

		// some methods on fileInfo
		fmt.Println("File Name: ", fileInfo.Name())
		fmt.Println("Is Directory or file: ", fileInfo.IsDir())
		fmt.Println("Size of file: ", fileInfo.Size()) // in bytes
		fmt.Println("Last Modify: ", fileInfo.ModTime())

		// create buffer to store bytes to read
		buf := make([]byte, fileInfo.Size())
		data, err := f.Read(buf) //data is in hexa-decimal format
		if err != nil {
			panic(err)
		}
		println(data) //give length of buf in bytes
		// printing data
		for i := 0; i < len(buf); i++ {
			print(string(buf[i]))
		}

	*/

	// -> os.ReadFile() -- (not preferred)
	/*
		f, err := os.ReadFile("example.txt") //load entire file into the memory
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fmt.Println(string(f))
	*/

	//-> Read Folder
	/*
		source := "../" // for current put "."
		dir, err := os.Open(source)
		if err != nil {
			panic(err)
		}
		defer dir.Close()

		n := 1 // if n<=0 retuns all files/folders present in the folder
		files, err := dir.ReadDir(n)
		if err != nil {
			panic(err)
		}

		for _, fi := range files {
			fmt.Println("File :", fi.Name())
		}
	*/

	// -> Create file
	/*
		file, err := os.Create("filename.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		file.WriteString("hi go")
		file.WriteString("nice Language") // append

		bytes := []byte("Hello Golang")
		file.Write(bytes)
	*/

	// -> delete file
	/*
		err := os.Remove("filename.txt")
		if err != nil {
			panic(err)
		}
	*/

	// -> Read from one file write into another file (streaming)

	// open source file
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	// create destination file
	destFile, err := os.Create("dest.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// create Reader and writer

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {

		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}

	}

	// Flusing writer bytes
	writer.Flush()
}
