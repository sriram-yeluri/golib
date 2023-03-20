package golib

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

func WriteHttpResponseToFile(resp *http.Response, filename string) {
	defer resp.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file :", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	written, err := io.Copy(writer, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
	log.Printf("Written %d bytes of data to %s", written, filename)
}

func ReadFromFile(filename string) []byte {
	// data , err := os.ReadFile(filename)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	//open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Cannot open the file:", err)
	}
	//close the file when we are done
	defer file.Close()

	//get file statistics to proceed
	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Cannot get the file stats :", err)
	}
	//create byte array to store the file contents
	data := make([]byte, stat.Size())
	//read the file
	bytes, err := file.Read(data)
	if err != nil {
		log.Fatal("Cannot read the file:", err)
	}
	log.Printf("Read %d bytes from file", bytes)

	return data
}
