// Package ziputil provides convenience functions for accessing zip files
package ziputil // import "timm.io/ziputil"

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// Reads the whole reader into memory and returns an io.ReadCloser for the requested file inside the zip file
func FileFromZipReader(r io.Reader, path string) (io.ReadCloser, error) {
	file, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	fileReader := bytes.NewReader(file)

	zipReader, err := zip.NewReader(fileReader, int64(len(file)))
	if err != nil {
		return nil, err
	}

	for _, f := range zipReader.File {
		if f.Name == path {
			return f.Open()
		}
	}

	return nil, fmt.Errorf("did not find file %s in zip file", path)
}
