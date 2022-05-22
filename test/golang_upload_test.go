package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func uploadFile(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, handler, err := request.FormFile("myFile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MINE Headers: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		panic(err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(writer, "Successfully Uploaded FIle\n")
	fmt.Println(string(fileBytes[:]))
}

func Test(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		uploadFile(writer, request)
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
