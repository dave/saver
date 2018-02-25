package saver

import (
	"io"

	"bytes"

	"github.com/flimzy/jsblob"
	"github.com/gopherjs/gopherjs/js"
)

func Save(filename, mimeType string, contents []byte) {
	js.Global.Call(
		"saveAs",
		jsblob.New([]interface{}{contents}, jsblob.Options{Type: mimeType}),
		filename,
	)
}

func SaveReader(filename, mimeType string, reader io.Reader) error {
	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, reader); err != nil {
		return err
	}
	js.Global.Call(
		"saveAs",
		jsblob.New([]interface{}{buf.Bytes()}, jsblob.Options{Type: mimeType}),
		filename,
	)
	return nil
}
