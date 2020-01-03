package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:local-password@tcp(127.0.0.1:3306)/fixtheplanet")

	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Connection to database verified!")

	_, err = db.Exec("USE fixtheplanet")
	if err != nil {
		log.Panicln(err.Error())
	}

	_, err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

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
