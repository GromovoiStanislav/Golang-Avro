package main

import (
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

	// Создание данных пользователя
	user := map[string]interface{}{
		"name":   "John Doe",
		"age":    30,
		"emails": []string{"john.doe@example.com", "johndoe@gmail.com"},
	}

	// Кодирование данных в Avro бинарный формат
	avroData, err := codec.BinaryFromNative(nil, user)
	if err != nil {
		fmt.Println("Error encoding Avro data:", err)
		return
	}


	// Подключение к серверу по TCP
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Отправка закодированных данных на сервер
	conn.Write([]byte(avroData))

	fmt.Println("Data sent to server")
}
