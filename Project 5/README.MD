## Avro with github.com/hamba/avro/v2 examples

### Setup:

```
go mod init avro-examples
go get github.com/hamba/avro/v2
```

### Start:

```
go run example_1.go

go run example_2.go

go run example_3.go // Marshal
go run example_4.go // NewEncoder
go run example_5.go // NewEncoderForSchema

go run example_6.go // Unmarshal
go run example_7.go // NewDecoder
go run example_8.go // NewDecoderForSchema

go run example_9.go // Parse
go run example_10.go // MustParse
```
