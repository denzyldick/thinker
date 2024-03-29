package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const URl = "http://localhost:8080/predict"

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

	if hint[0] == "faces" {
		body = []byte(`{
  "service": "faces",
  "parameters": {
    "input": {},
    "output": {
      "confidence_threshold": 0.4,
      "bbox": true
    },
    "mllib": {
      "gpu": true
    }
  },
  "data": [
"` + image[0] + `"
  ]
}
`)
	}
	if hint[0] == "emotions" {
		body = []byte(`{
  "service": "faces_emo",
  "parameters": {
    "input": {},
    "output": {
      "confidence_threshold": 0.4,
      "bbox": true
    },
    "mllib": {
      "gpu": true
    }
  },
  "data": [
    "` + image[0] + `"
  ]
}

`)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, URl, bytes.NewBuffer(body))
	log.Println(err)

	response := response{}
	resp, err := client.Do(req)
	log.Println(err)

	if err == nil {
		fmt.Println("status code", resp.StatusCode)
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			data, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(image[0], hint[0], "predicted")
			json.Unmarshal(data, &response)
			w.Header().Add("Content-type", "application/json")
			w.WriteHeader(200)
			w.Write(data)
		}
	}
}

type response struct {
}

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	http.HandleFunc("/think", think)
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}
