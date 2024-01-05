package main

import (
  "fmt"
  "log"

  "github.com/hamba/avro/v2"
)


func main() {
  schema, err := avro.Parse(`{
    "type": "record",
    "name": "simple",
    "namespace": "org.hamba.avro",
    "fields" : [
        {"name": "a", "type": "long"},
        {"name": "b", "type": "string"}
    ]
  }`)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(schema.Type())
  // Outputs: record
}