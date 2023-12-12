## Avro official examples

### Setup:

```
go mod init avro-examples
go get github.com/linkedin/goavro/v2
```

### Start:

```
go run example_1.go
go run example_2.go
go run example_3.go
go run example_4.go -count -schema users.avro
go run example_4.go -count -schema event.avro pets.avro
```