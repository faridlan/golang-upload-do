package main

import (
	"net/http"

	"github.com/faridlan/golang-upload-do/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		controller.UploadImage(writer, request)
	})
	mux.HandleFunc("/delete", func(writer http.ResponseWriter, request *http.Request) {
		controller.DeleteImage(writer, request)
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
