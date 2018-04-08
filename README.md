# lesphina

syntax analyzer of Golang source code.

## (Ideal) usage

### load
```golang
import "lesphina"
var les = lesphina.Read("my_go_code.go")
```

### query
```golang
allImports       := les.Query().ByType(lesphina.IMPORT).All()
allStructs       := les.Query().ByType(lesphina.STRUCT).All()
firstInterfaces  := les.Query().ByType(lesphina.INTERFACE).First()
someStructs      := les.Query().ByName("Foo???").ByType(lesphina.STRUCT).All()
someNotExported  := les.Query().ByName("bar*").ByType(lesphina.VAR).Exported(false).First()
```

### read meta
```golang
myStruct := les.Query().ByName("MyStruct").ByType(lesphina.STRUCT).First()
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

### advanced query
```golang
myFunc := les.Query().Which(func (q *lesphina.Query) bool {
	if q.Name == "MyFunc*" && q.Type == lesphina.FUNCTION {
		if strings.Contains(q.RawString(), "/some_pattern/") {
			return true
		}
	}
}).First()
```


