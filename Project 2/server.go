package main

import (
	"bufio"
	"fmt"
	"net"
	"io/ioutil"

	"github.com/linkedin/goavro/v2"
)

func main() {
	schema, err := ioutil.ReadFile("user_schema.avsc")
	if err != nil {
		panic(err)
	}

	// Создание Avro схемы
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		fmt.Println("Error creating Avro codec:", err)
		return
	}

	// Создание сервера TCP
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 3000")

	for {
		// Принятие входящего подключения
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Client connected!")

		// Чтение данных от клиента с использованием bufio.Scanner
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receivedData := scanner.Text()

			// Декодирование данных и вывод на сервере
			native, _, err := codec.NativeFromBinary([]byte(receivedData))
			if err != nil {
				fmt.Println("Error decoding Avro data:", err)
				conn.Close()
				break
			}

			decodedUser := native.(map[string]interface{})
			fmt.Println("Received Data:", decodedUser)
		}

		conn.Close()
		fmt.Println("Client disconnected!")
	}
}
