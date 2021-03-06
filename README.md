# lesphina [![Build Status](https://travis-ci.org/straightdave/lesphina.svg?branch=master)](https://travis-ci.org/straightdave/lesphina)

Syntax analyzer of Golang source code.

![lesphina & arion](https://i.pinimg.com/736x/35/e9/42/35e942e53b10d00138db8156ef6b73d1---s.jpg)

## Usage

### Load from source file
```golang
import "lesphina"
var les = lesphina.Read("my_go_code.go")
```
Now most of language entries are in les.Meta structure.

### Query
```golang
q := les.Query()

var theEntryIWant Entry

// ByKind() and ByName() could be chained
// using First() or All() to resolve this query
theEntryIWant = q.ByKind(KindInterface).ByName("someName").First()

theInterfaceIWant, ok := theEntryIWant.(*Interface)

// use it if ok!
```

Once `les := lesphina.Read()` is called, `les.Query()` could be called many times and would return a clean query handle each time.

You can use `~` at the beginning or the end of query words in `ByName()`:
```golang
entries := q.ByName("Prefix~").All()
entries2 := q.ByName("~Suffix").All()
```

