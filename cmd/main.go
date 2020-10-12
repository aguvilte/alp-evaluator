package main

import (
	//"github.com/unlar/alp-evaluator/cmd/api"
	//"github.com/lucperkins/party"
	//"io/ioutil"
	//"net/http"
	//"log"
	//"reflect"
	//"io"
	//"bytes"
	//"mime/multipart"
	//"os"
	//"fmt"
	"github.com/unlar/alp-evaluator/cmd/api"
)

func main()  {
	api.StartApp()
	//handler := &party.MultipartRequestHandler{
	//	MaxBytes: 32 << 20,
	//}
	//
	//srv := &http.Server{
	//	Addr: ":8080",
	//	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		res, err := handler.Handle(w, r)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		log.Println("Filename:", res.Header.Filename)
	//
	//		bs, err := ioutil.ReadAll(res.File)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		log.Println("File contents:")
	//		//log.Print(string(bs))
	//		log.Print(reflect.TypeOf(bs))
	//
	//		req, err := http.NewRequest("POST", "http:localhost:8089/execute/scorpion", res.File)
	//
	//	}),
	//}
	//
	//log.Fatal(srv.ListenAndServe())
}