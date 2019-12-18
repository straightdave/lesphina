package lesphina

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// Errors ...
var (
	ErrNoSuchFile      = errors.New("No Such File")
	ErrDirNotSupported = errors.New("Dir Is Not Supported")
)

// Lesphina ...
type Lesphina struct {
	FileInfo *os.FileInfo `json:"fileInfo"`
	Meta     *Meta        `json:"meta"`
}

// Read reads source meta from a file and returns a lesphina instance.
// This is probably the most common used entrance of this package.
func Read(source string) (*Lesphina, error) {
	var fi os.FileInfo
	var err error
	if fi, err = os.Stat(source); err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoSuchFile
		}
		return nil, err
	}

	if fi.IsDir() {
		return nil, ErrDirNotSupported
	}

	var meta *Meta
	if meta, err = parseSource(source); err != nil {
		return nil, err
	}

	return &Lesphina{
		FileInfo: &fi,
		Meta:     meta,
	}, nil
}

// DumpString dumps source meta into a compressed base64 string.
func (les *Lesphina) DumpString() string {
	raw := les.Meta.JSON()
	var buff bytes.Buffer
	gz := gzip.NewWriter(&buff)
	if _, err := gz.Write([]byte(raw)); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buff.Bytes())
}

// Restore restores lesphina instance from a compressed base64 string.
func Restore(raw string) *Lesphina {
	raw = strings.TrimSpace(raw)
	zippedStr, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		panic(err)
	}

	rdata := bytes.NewReader(zippedStr)
	upzippedStr, err := gzip.NewReader(rdata)
	if err != nil {
		panic(err)
	}

	jsonStr, err := ioutil.ReadAll(upzippedStr)
	if err != nil {
		panic(err)
	}

	var meta Meta
	if err := json.Unmarshal(jsonStr, &meta); err != nil {
		panic(err)
	}

	return &Lesphina{Meta: &meta}
}
