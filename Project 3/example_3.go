package main

import (
	"os"
	"fmt"

	"github.com/linkedin/goavro/v2"
)

const loginEventAvroSchema = `{"type": "record", "name": "LoginEvent", "fields": [{"name": "Username", "type": "string"}]}`

func main() {
	codec, err := goavro.NewCodec(loginEventAvroSchema)
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"Username": "superman",
	}


	binary, err := codec.BinaryFromNative(nil, m)
		if err != nil {
			panic(err)
		}
		_ = binary

	{
		native, _, err := codec.NativeFromBinary(binary)
		if err != nil {
			fmt.Println(err)
		}

		decodedData := native.(map[string]interface{})
		fmt.Println(decodedData)
	}
	

	
	single, err := codec.SingleFromNative(nil, m)
	if err != nil {
		panic(err)
	}
	_ = single
	
	
	{
		native, _, err := codec.NativeFromSingle(single)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(native)
	}

	
	var values []map[string]interface{}
	values = append(values, m)
	values = append(values, map[string]interface{}{"Username": "batman"})
	values = append(values, map[string]interface{}{"Username": "wonder woman"})

	f, err := os.Create("event.avro")
	if err != nil {
		panic(err)
	}
	ocfw, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:     f,
		Codec: codec,
	})
	if err != nil {
		panic(err)
	}
	if err = ocfw.Append(values); err != nil {
		panic(err)
	}


	// Open the Avro file for reading
	f, err = os.Open("event.avro")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	// Create an OCF Reader
	ocfr, err := goavro.NewOCFReader(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Records in OCF File");
	for ocfr.Scan() {
		record, err := ocfr.Read()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("record", record)
	}
}