package somePackage

import (
    "bytes"
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "os"
    "regexp"
    "strconv"
    "strings"
    "text/template"
    "time"

    "github.com/golang/protobuf/jsonpb"
    "github.com/golang/protobuf/proto"
    protoempty "github.com/golang/protobuf/ptypes/empty"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"

    gozip "github.com/straightdave/gozip/lib"
    "github.com/straightdave/lesphina"
    trunks "github.com/straightdave/trunks/lib"
)

func Func1() {

}

func Func() {

}

func Func2(a string) bool {
	return false
}

func (a *Astruct) Func3() int {
	return 0
}

type Int0 interface{
    Name(user *User) (haha map[string]string)
    SayHello(name, lastName string, age int) (t []string, t2 error)
}

type Int1 interface{}

type Int2 interface{
    Name(user *User) (haha map[string]string)
    SayHello(name, lastName string, age int) (t []string, t2 error)
}

type Str1 struct {}
type Str2 struct {
    Name string `json:"some,omitempty"`
}

// alias
type Aha int

var v1 int
var v2 string

