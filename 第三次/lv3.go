package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	_ = os.Remove("plan.txt")
	WriteData()
	ReadData()

}

func WriteData() {
	file, err := os.Create("plan.txt")
	if err != nil {
		log.Fatalf("create file error : %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte("I’m not afraid of difficulties and insist on learning programming"))
	if err != nil {
		log.Fatalf("write data error : %v", err)
	}
}

func ReadData() {
	buffer := make([]byte, 2048)
	file, err := os.OpenFile("plan.txt", os.O_RDWR, 777)
	if err != nil {
		log.Fatalf("open file error : %v", err)
	}
	defer file.Close()

	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatalf("read data error : %v", err)
	}

	fmt.Printf("read %d byte : %s", n, string(buffer))
}
