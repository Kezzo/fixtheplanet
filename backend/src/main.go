package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Kezzo/fixtheplanet/src/common"
)

func main() {
	log.Println("Starting server...")

	db, err := common.InitDatabase()

	if err != nil {
		log.Panicln(err.Error())
	}

	//db.SQLDB.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's alive!"))
	})

	http.HandleFunc("/gitdata", func(w http.ResponseWriter, r *http.Request) {
		resp, err := common.GetIssuesFromGithub()

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
		}

		insert, err := db.SQLDB.Prepare("INSERT INTO issues VALUES( ?, ?, ?, ?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()

		count := 0
		for _, repo := range resp.Data.Search.Edges {
			for _, issue := range repo.Node.Issues.Edges {
				_, err = insert.Exec(nil, issue.Node.Title, repo.Node.NameWithOwner, repo.Node.PrimaryLanguage.Name, "github")
				if err != nil {
					log.Println("Error inserting issue: " + err.Error())
				} else {
					count++
				}
			}
		}

		respString := "Inserted " + strconv.Itoa(count) + " issues into database"

		log.Println(respString)
		w.Write([]byte(respString))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
