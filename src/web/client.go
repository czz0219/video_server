package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"log"
	"net/http"
)

//未经初始化的变量
var httpClient *http.Client

func init() {
	httpClient = &http.Client{} //初始化
	log.Println("main package client init")
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) error{
	var resp *http.Response
	var err error
	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf(err.Error())
			return err
		}
		normalResponse(w, resp)
	case http.MethodPost: //只有POST才有 REQUEST_BODY
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf(err.Error())
			return err
		}
		normalResponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", b.Url, nil)
		req.Header = r.Header
		log.Println(" MethodDelete will delete :"+b.Url)
		resp, err = httpClient.Do(req)
		log.Println(" after MethodDelete")
		if err != nil {
			log.Printf(err.Error())
			return err
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
	}
	return nil
}
func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}
	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}
