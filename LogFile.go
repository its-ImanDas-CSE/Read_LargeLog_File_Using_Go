package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {

	fmt.Println("hey the output is correct?")
	file, err := os.Open("Day-4_log_file_500mb (1).log")

	if err != nil {
		log.Fatal("there is some error while opening the file", err.Error())

	}
	defer file.Close()
	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal("error in fetching file size ", err.Error())
	}
	fileSize := fileInfo.Size()

	noOfChunks := 500
	chunkSize := fileSize / int64(noOfChunks)
	fmt.Println("here is the size", chunkSize)
	for i := 0; i < 500; i++ {
		start := int64(i) * chunkSize
		end := start + chunkSize

		if i == noOfChunks-1 {
			end = fileSize
		}

		wg.Add(1)
		go readFile(file, start, end)

	}

	wg.Wait()
	fmt.Println("$$$$$$$$$$$$$$   file reading complete $$$$$$$$$$$$$")

}

func readFile(file *os.File, start, end int64) {

	defer wg.Done()

	_, err := file.Seek(start, 0)

	if err != nil {
		log.Fatal("error seeking to start of the chunk", err.Error())
	}

	reader := bufio.NewReader(file)

	// Read the chunk and process it

	bytesToRead := end - start
	buffer := make([]byte, bytesToRead)

	_, err = reader.Read(buffer)
	if err != nil {
		log.Println("Error reading chunk:", err)
		return
	}

	// Process the chunk here (for example, print the content)
	fmt.Printf("the data is %s ", buffer)

	// reader := csv.NewReader(file)
	// for {
	//  data, err1 := reader.Read()

	//  if err1 != nil {
	//      fmt.Println("there is no further data", err1.Error())
	//      break
	//  }
	//  fmt.Println("the data is ", data)

	// }

}
