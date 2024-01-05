package main

import (
  "encoding/json"
  "fmt"

  "github.com/hamba/avro/v2"
)

type Person struct {
  Name string `json:"name" avro:"name"`
  Age  int    `json:"age" avro:"age"`
}

func main() {

  var john *Person = &Person{
    Name: "john",
    Age:  36,
  }
  fmt.Println(john) // &{john 36}

  //json example
  json_bytes, err := json.Marshal(john)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(json_bytes))// {"name":"john","age":36}
  fmt.Println(json_bytes) // [123 34 110 97 109 101 34 58 34 106 111 104 110 34 44 34 97 103 101 34 58 51 54 125]
  fmt.Println(len(json_bytes)) // 24

  //avro schema
  var schema_string string = `{
    "type":"record",
    "name":"person",
    "fields":[
      {
        "name":"name",
        "type":"string"
       },
      {
        "name":"age",
        "type":"int"
       }
     ]
  }`


  schema, err := avro.Parse(schema_string)
  if err != nil {
    panic(err)
  }
  fmt.Println(schema.Type()) // record

  avro_bytes, err := avro.Marshal(schema, john)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(avro_bytes))// johnH
  fmt.Println(avro_bytes) // [8 106 111 104 110 72]
  fmt.Println(len(avro_bytes)) //6

///Unmarshal
{
  out := Person{}
  err = avro.Unmarshal(schema, avro_bytes, &out)
  if err != nil {
    panic(err)
  }

  fmt.Println(out) // {john 36}
  fmt.Println(out.Name) // john
  fmt.Println(out.Age) // 36
}




  //Schema creation from code
  name_field, err := avro.NewField("name", avro.MustParse(`{"type": "string"}`))
  if err != nil {
      panic(err)
  }
  
  age_field, err := avro.NewField("age", avro.MustParse(`{"type": "int"}`))
  if err != nil {
      panic(err)
  }


  generated_schema, err := avro.NewRecordSchema("person", "", []*avro.Field{name_field, age_field})
  if err != nil {
      panic(err)
  }
  

  fmt.Println(generated_schema.String())

  avro_bytes, err = avro.Marshal(generated_schema, john)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(avro_bytes)) // johnH
  fmt.Println(avro_bytes) // [8 106 111 104 110 72]
  fmt.Println(len(avro_bytes)) // 6



  ///Unmarshal
  {
    out := Person{}
    err = avro.Unmarshal(generated_schema, avro_bytes, &out)
    if err != nil {
      panic(err)
    }

    fmt.Println(out) // {john 36}
    fmt.Println(out.Name) // john
    fmt.Println(out.Age) // 36
  }
}