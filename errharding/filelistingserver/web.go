package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type appHanlder func(writer http.ResponseWriter,
	request *http.Request) error

type usererror string

func (e usererror) error() string {
	return e.Message()
}

func (e usererror) Message() string {
	return string(e)
}

func errWrapper(
	handler appHanlder) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic :%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s",
				err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				//http.Error(writer,                          //向谁汇报，也就是向哪一个出错的writer返回信息
				//	http.StatusText(http.StatusNotFound),   //返回的内容
				//	http.StatusNotFound)                    //返回的状态码
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

func handleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	arr, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = writer.Write(arr)
	if err != nil {
		return nil
	}
	return nil
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/",
		errWrapper(handleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
