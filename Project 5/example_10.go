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
  
  fmt.Println(schema.Type())
  // Outputs: record
}