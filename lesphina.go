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

var (
	ErrNoSuchFile      = errors.New("No Such File")
	ErrDirNotSupported = errors.New("Dir Not Supported")
)

type Lesphina struct {
	FileInfo *os.FileInfo `json:"file_info"`
	Meta     *Meta        `json:"meta"`
}

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

func (les *Lesphina) DumpString() string {
	raw := les.Meta.Json()
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
