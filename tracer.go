package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	tracerPort := 5003
	fmt.Printf("PicPay Tracer %v 2.0 \n", tracerPort)

	http.HandleFunc("/", traceHandler)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", tracerPort),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}

func traceHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusAccepted)
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	if err != nil {
		log.Printf("Erro de leitura: %v", err)
		http.Error(writer, "Erro recebendo body", http.StatusBadRequest)
		return
	}
	echoLog(string(body), request.Header)
}

func bodyTransformer(rawBody string, header http.Header) (string, string) {

	ctype := header.Get("content-type")

	var body, name string
	var err error

	switch ctype {

	case "application/x-www-form-urlencoded":
		parsedBody, _ := url.ParseQuery(rawBody)
		body = parsedBody.Get("trace")
		name = parsedBody.Get("name")
	case "application/json":
		var res map[string]interface{}
		err = json.Unmarshal([]byte(rawBody), &res)
		body = res["trace"].(string)
		rawName := res["name"]
		if rawName != nil {
			name = rawName.(string)
		} else {
			name = ""
		}
	default:
		body = rawBody
		name = ""
	}

	if name == "" {
		t := time.Now()
		name = t.Format(time.Stamp)
	}
	if err != nil {
		log.Println("Erro decodificando payload")
	}

	return body, name
}

func echoLog(body string, header http.Header) {
	parsedBody, parsedHead := bodyTransformer(body, header)

	log.Println("---------------------")
	log.Println(parsedHead)
	log.Println("---------------------")
	log.Println("")
	log.Println(string(parsedBody))
	log.Println("")
	log.Println("---------------------")
}