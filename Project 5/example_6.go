package main

import (
  "fmt"

  "github.com/hamba/avro/v2"
)


func main() {
  schema := avro.MustParse(`{
    "type": "record",
    "name": "simple",
    "namespace": "org.hamba.avro",
    "fields" : [
        {"name": "a", "type": "long"},
        {"name": "b", "type": "string"}
    ]
  }`)

  type SimpleRecord struct {
    A int64  `avro:"a"`
    B string `avro:"b"`
  }

  data := []byte{0x36, 0x06, 0x66, 0x6F, 0x6F} // Your Avro data here
  simple := SimpleRecord{}
  if err := avro.Unmarshal(schema, data, &simple); err != nil {
    fmt.Println("error:", err)
  }

  fmt.Printf("%+v", simple)
  // Outputs: {A:27 B:foo}
}