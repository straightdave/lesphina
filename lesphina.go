package lesphina

import (
	"os"
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
