package main

import (
  "fmt"
  "bytes"

  "github.com/hamba/avro/v2"
)


func main() {
  schema := `{
    "type": "record",
    "name": "simple",
    "namespace": "org.hamba.avro",
    "fields" : [
        {"name": "a", "type": "long"},
        {"name": "b", "type": "string"}
    ]
  }`

  type SimpleRecord struct {
    A int64  `avro:"a"`
    B string `avro:"b"`
  }

  r := bytes.NewReader([]byte{0x36, 0x06, 0x66, 0x6F, 0x6F}) // Your reader goes here
  decoder, err := avro.NewDecoder(schema, r)
  if err != nil {
    fmt.Println("error:", err)
  }

  simple := SimpleRecord{}
  if err := decoder.Decode(&simple); err != nil {
    fmt.Println("error:", err)
  }

  fmt.Printf("%+v", simple)
  // Outputs:  {A:27 B:foo}
}