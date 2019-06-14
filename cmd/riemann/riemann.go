package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func FontBytes() (b []byte) {
	in := res["font"]
	r, err := gzip.NewReader(bytes.NewBuffer(in))
	if err != nil {
		return
	}
	defer r.Close()
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}
