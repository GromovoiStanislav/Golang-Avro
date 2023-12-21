package main

import (
	"fmt"
	"github.com/linkedin/goavro/v2"
)

func main() {
	// Определение схемы Avro
	schemaJSON := `{
		"type": "record",
		"name": "Visit",
		"fields": [
			{"name": "city", "type": "string"},
			{"name": "zipCodes", "type": {"type": "array", "items": "string"}},
			{"name": "visits", "type": "int"}
		]
	}`

	// Создание Avro схемы
	codec, err := goavro.NewCodec(schemaJSON)
	if err != nil {
		fmt.Println("Error creating Avro codec:", err)
		return
	}

	// Пример данных для кодирования
	data1 := map[string]interface{}{"city": "Seattle", "zipCodes": []string{"98101"}, "visits": 3}
	data2 := map[string]interface{}{"city": "NYC", "zipCodes": []string{}, "visits": 0}
	data3 := map[string]interface{}{"city": "Cambridge", "zipCodes": []string{"02138", "02139"}, "visits": 2}

	// Кодирование данных в Avro бинарный формат
	buf1, err := codec.BinaryFromNative(nil, data1)
	if err != nil {
		fmt.Println("Error encoding Avro data:", err)
		return
	}

	buf2, err := codec.BinaryFromNative(nil, data2)
	if err != nil {
		fmt.Println("Error encoding Avro data:", err)
		return
	}

	buf3, err := codec.BinaryFromNative(nil, data3)
	if err != nil {
		fmt.Println("Error encoding Avro data:", err)
		return
	}


	// Пример массива закодированных данных
	bufs := [][]byte{buf1, buf2, buf3}

	// Декодирование каждого буфера
	for _, buf := range bufs {
		native, _, err := codec.NativeFromBinary(buf)
		if err != nil {
			fmt.Println("Error decoding Avro data:", err)
			return
		}

		// native теперь содержит декодированные данные
		fmt.Println("Decoded Data:", native)
	}
}
