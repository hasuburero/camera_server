package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func (self *API) Get(w http.ResponseWriter, r *http.Request) {
	err := self.cam.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println("camera.Get error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fd, err := os.Open(self.cam.Filename)
	if err != nil {
		fmt.Println(err)
		fmt.Println("os.Open error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer fd.Close()

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	fw, err := mw.CreateFormFile("file", self.Filename)
	_, err = io.Copy(fw, fd)
	if err != nil {
		fmt.Println(err)
		fmt.Println("io.Copy error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := mw.FormDataContentType()
	err = mw.Close()

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(body.Bytes())
}
