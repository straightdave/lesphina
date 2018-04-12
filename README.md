# lesphina [![Build Status](https://travis-ci.org/straightdave/lesphina.svg?branch=master)](https://travis-ci.org/straightdave/lesphina)

Syntax analyzer of Golang source code.

![lesphina & arion](https://i.pinimg.com/736x/35/e9/42/35e942e53b10d00138db8156ef6b73d1---s.jpg)

## (Ideal) usage

### load
```golang
import "lesphina"
var les = lesphina.Read("my_go_code.go")
```

### meta
```golang
myStruct := les.Meta.Structs[0]
fmt.Println(myStruct.Name)
fmt.Println(myStruct.Fields[0].Name)
fmt.Println(myStruct.Fields[0].Type)
fmt.Println(myStruct.Fields[0].IsPointer)
fmt.Println(myStruct.Fields[0].Tag)

methods := myStruct.Methods()
fmt.Println(methods[0].Name)
fmt.Println(methods[0].Recv.IsPointer)
fmt.Println(methods[0].Args[0])
fmt.Println(methods[0].Args[0].Name)
fmt.Println(methods[0].Args[0].Type)
fmt.Println(methods[0].Args[0].IsPointer)
```


