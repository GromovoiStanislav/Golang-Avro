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

  simple := SimpleRecord{A: 27, B: "foo"}
  b, err := avro.Marshal(schema, simple)
  if err != nil {
    fmt.Println("error:", err)
  }

  fmt.Println(b)
  // Outputs: [54 6 102 111 111]
}