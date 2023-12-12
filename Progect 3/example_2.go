package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/linkedin/goavro/v2"
)

func main() {
	avroSchema := `
	{
	  "type": "record",
	  "name": "test_schema",
	  "fields": [
		{
		  "name": "time",
		  "type": "long"
		},
		{
		  "name": "customer",
		  "type": "string"
		}
	  ]
	}`

	// Writing OCF data
	var ocfFileContents bytes.Buffer
	writer, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:      &ocfFileContents,
		Schema: avroSchema,
	})
	if err != nil {
		fmt.Println(err)
	}
	err = writer.Append([]map[string]interface{}{
		{
			"time":     1617104831721,
			"customer": "customer1",
		},
		{
			"time":     1717104831722,
			"customer": "customer2",
		},
	})

	err = writer.Append([]map[string]interface{}{
		{
			"time":     1617104831723,
			"customer": "customer3",
		},
		{
			"time":     1717104831724,
			"customer": "customer4",
		},
		{
			"time":     1717104831725,
			"customer": "customer5",
		},
	})


	fmt.Println("ocfFileContents", ocfFileContents.String())

	// Reading OCF data
	ocfReader, err := goavro.NewOCFReader(strings.NewReader(ocfFileContents.String()))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Records in OCF File");
	for ocfReader.Scan() {
		record, err := ocfReader.Read()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("record", record)
	}
}