package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Image(writer http.ResponseWriter, request *http.Request) (string, string) {
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		panic(err)
	}

	file, handler, err := request.FormFile("myFile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(writer, "Successfully Uploaded FIle\n")

	x := strconv.Itoa(int(handler.Size))
	y := string(fileBytes)
	return x, y
}
