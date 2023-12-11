package main

import (
	"fmt"
	"github.com/linkedin/goavro/v2"
)

func main() {
	// Определение схемы Avro
	schemaJSON := `{
		"type": "record",
		"name": "User",
		"fields": [
			{"name": "name", "type": "string"},
			{"name": "age", "type": "int"},
			{"name": "emails", "type": {"type": "array", "items": "string"}}
		]
	}`

	// Создание Avro схемы
	codec, err := goavro.NewCodec(schemaJSON)
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

	fmt.Println("Encoded:", avroData)

	// Декодирование бинарных данных обратно в объект
	native, _, err := codec.NativeFromBinary(avroData)
	if err != nil {
		fmt.Println("Error decoding Avro data:", err)
		return
	}

	decodedUser := native.(map[string]interface{})
	fmt.Println("Decoded:", decodedUser)

	// Получение значений полей "name" и "age"
	name := decodedUser["name"].(string)
	age := int(decodedUser["age"].(int32))

	fmt.Println("name:", name)
	fmt.Println("age:", age)

	// Получение значения поля "emails"
	emails := decodedUser["emails"].([]interface{})

	fmt.Println("emails:", emails)
	fmt.Println("email:", emails[0])
	
}