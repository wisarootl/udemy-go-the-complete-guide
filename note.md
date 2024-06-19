- build app and run
```shell
go mod init example.com/first-app
go build
./first-app.exe
```

- run without build
```shell
go run .
```

- install dependencies
```shell
go get
go get github.com/Pallinder/go-randomdata
```


- pointer
```go
var agePointer *int

agePointer = &age // get pointer

fmt.Println("Address of age variable:", agePointer)
fmt.Println("Age:", *agePointer) // dereference the pointer
```
