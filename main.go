package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const URl = "0.0.0.0:8008/predict"

func think(w http.ResponseWriter, r *http.Request) {

	hint, ok := r.URL.Query()["hint"]

	if !ok || len(hint[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	image, ok := r.URL.Query()["image"]

	if !ok || len(hint[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	var body []byte
	if hint[0] == "object" {
		body = []byte(`{
"service": "ilsvrc_googlenet",
"parameters": {
"output": {
"best": 3
},
"mllib": {
"gpu": true
}
},
"data": [
"` + image[0] + `"
]
}`)
	}

	fmt.Println(string(body))
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, URl, bytes.NewBuffer(body))
	log.Println(err)

	response := response{}
	resp, err := client.Do(req)
	log.Println(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(data, &response)
		w.Header().Add("Content-type", "application/json")

		w.Write(data)
		w.WriteHeader(200)
		fmt.Println(image[0], hint[0], "predicted")
	}
}

type response struct {
}

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	http.HandleFunc("/think", think)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
