package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Kezzo/fixtheplanet/src/common"
	"golang.org/x/crypto/acme/autocert"
)

// Issue ..
type Issue struct {
	Title    string
	Repo     string
	Number   int
	Language string
}

// IssueResponse ..
type IssueResponse struct {
	Issues []Issue
}

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

	http.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		languages := query["language"]

		for _, lang := range languages {
			if lang == "" {
				http.Error(w, "Url Param 'language' is empty", http.StatusBadRequest)
				return
			}
		}

		issues, err := queryIssues(db, languages)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}

		resp := IssueResponse{
			Issues: issues,
		}

		respBytes, err := json.Marshal(resp)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(respBytes)
	})

	http.HandleFunc("/gitdata", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-token")
		envToken := os.Getenv("GIT-PATH-TOKEN")

		if token != envToken {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		resp, err := common.GetIssuesFromGithub()

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		inserted, err := insertIssuesIntoDB(db, resp)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respString := "Inserted " + strconv.Itoa(inserted) + " issues into database"

		log.Println(respString)
		w.Write([]byte(respString))
	})

	if os.Getenv("ENV") == "LOCAL" {
		log.Fatal(http.ListenAndServe(":80", nil))
	} else {
		listener := autocert.NewListener("cloud.fixthepla.net")
		log.Fatal(http.Serve(listener, nil))
	}
}

func insertIssuesIntoDB(db *common.Database, resp *common.GithubResponse) (int, error) {
	stmt, err := db.SQLDB.Prepare("INSERT INTO issues VALUES( ?, ?, ?, ?, ? )")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	count := 0
	for _, repo := range resp.Data.Search.Edges {
		for _, issue := range repo.Node.Issues.Edges {
			_, err = stmt.Exec(nil, issue.Node.Title, repo.Node.NameWithOwner, issue.Node.Number, repo.Node.PrimaryLanguage.Name)
			if err != nil {
				log.Println("Error inserting issue: " + err.Error())
			} else {
				count++
			}
		}
	}

	return count, nil
}

func queryIssues(db *common.Database, languages []string) ([]Issue, error) {
	queryString := "SELECT * FROM issues"

	for i := 0; i < len(languages); i++ {
		if i == 0 {
			queryString += " WHERE language = ?"
		} else {
			queryString += " OR language = ?"
		}
	}

	queryString += " ORDER BY RAND() LIMIT 20" // TODO: Implement proper random pagination

	stmt, err := db.SQLDB.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var rows *sql.Rows
	if languages != nil && len(languages) > 0 {
		var inputValues []interface{}
		inputValues = make([]interface{}, len(languages))

		for i, language := range languages {
			inputValues[i] = language
		}

		rows, err = stmt.Query(inputValues...)
	} else {
		rows, err = stmt.Query()
	}

	if err != nil {
		return nil, err
	}

	issues := make([]Issue, 0)
	issueID := ""

	defer rows.Close()
	count := 0
	for rows.Next() {
		issue := Issue{}
		err = rows.Scan(&issueID, &issue.Title, &issue.Repo, &issue.Number, &issue.Language)

		if err != nil {
			return nil, err
		}

		issues = append(issues, issue)
		count++
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return issues, nil
}
