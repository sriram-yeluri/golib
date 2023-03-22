package golib

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// WriteHttpResponseToFile function saves http response body to a given file name.
func WriteHttpResponseToFile(resp *http.Response, filename string) error {
	defer resp.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	written, err := io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}
	writer.Flush()
	log.Printf("Written %d bytes of data to %s", written, filename)
	return nil
}

// ReadFromFile function reads data from file and returns []byte stream.
// This function reads file line by line.
func ReadFromFile(filename string) ([]byte, error) {
	// data , err := os.ReadFile(filename)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	//open the file
	file, err := os.Open(filename)
	if err != nil {
		Error.Printf("Cannot open the file : %s", filename)
		return nil, err
	}
	//close the file when we are done
	defer file.Close()

	//get file statistics to proceed
	stat, err := file.Stat()
	if err != nil {
		Error.Println("Error getting the file stats :", err)
		return nil, err
	}
	//create byte array to store the file contents
	data := make([]byte, stat.Size())
	//read the file
	bytes, err := file.Read(data)
	if err != nil {
		Error.Println("Cannot read the file:", err)
		return nil, err
	}
	log.Printf("Read %d bytes from file", bytes)

	return data, nil
}

// WriteStructToJsonFile method writes strucure data to a file in json format.
func WriteStructToJsonFile(data []byte, filename string) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
