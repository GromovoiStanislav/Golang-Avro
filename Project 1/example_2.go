package main

import (
	"fmt"
	"github.com/linkedin/goavro/v2"
	"io/ioutil"
)

func encodeAndWriteToFile(codec *goavro.Codec, data map[string]interface{}, filename string) error {
	// Кодирование данных в Avro бинарный формат
	avroData, err := codec.BinaryFromNative(nil, data)
	if err != nil {
		return fmt.Errorf("Error encoding Avro data: %v", err)
	}

	// Запись бинарных данных в файл
	err = ioutil.WriteFile(filename, avroData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	fmt.Println("Data written to", filename)
	return nil
}

func readAndDecodeFromFile(codec *goavro.Codec, filename string) (map[string]interface{}, error) {
	// Считывание бинарных данных из файла
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading from file: %v", err)
	}

	// Декодирование бинарных данных обратно в объект
	native, _, err := codec.NativeFromBinary(fileData)
	if err != nil {
		return nil, fmt.Errorf("Error decoding Avro data: %v", err)
	}


	// Convert native Go form to textual Avro data
    // textual, err := codec.TextualFromNative(nil, native)
    // if err != nil {
    //     fmt.Println(err)
    // }
	//fmt.Println(string(textual))


	decodedData := native.(map[string]interface{})
	return decodedData, nil
}

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

	// Кодирование данных и запись в файл
	err = encodeAndWriteToFile(codec, user, "user.avro")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Считывание данных из файла и декодирование
	decodedUser, err := readAndDecodeFromFile(codec, "user.avro")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Decoded:", decodedUser)
}
