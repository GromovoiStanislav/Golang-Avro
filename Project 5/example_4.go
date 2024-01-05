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

  w := &bytes.Buffer{}
  encoder, err := avro.NewEncoder(schema, w)
  if err != nil {
    fmt.Println("error:", err)
  }

  simple := SimpleRecord{A: 27, B: "foo"}
  if err := encoder.Encode(simple); err != nil {
    fmt.Println("error:", err)
  }

  fmt.Println(w.Bytes()) 
  // Outputs:  [54 6 102 111 111]
}