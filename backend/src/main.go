package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's alive!"))
	})

	http.HandleFunc("/gitdata", func(w http.ResponseWriter, r *http.Request) {

		fileData, err := ioutil.ReadFile("../queries/get-issues.gql")

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
		}

		mapData := map[string]string{
			"query":     string(fileData),
			"variables": "",
		}
		requestBody, err := json.Marshal(mapData)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
		}

		req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(requestBody))
		req.Header.Add("Authorization", "bearer "+os.Getenv("GITHUB-TOKEN"))
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
		}

		w.Write(bodyBytes)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
