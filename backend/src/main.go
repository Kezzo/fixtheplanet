package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Kezzo/fixtheplanet/src/common"
	"golang.org/x/crypto/acme/autocert"
)

// Issue ..
type Issue struct {
	Title    string
	Repo     string
	Number   int
	Language string
	Labels   []IssueLabel
}

// IssueLabel ..
type IssueLabel struct {
	Name  string
	Color string
}

// IssueResponse ..
type IssueResponse struct {
	Issues     []Issue
	PagingSeed int
	NextOffset int
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
		log.Println("Received /issues")

		unixTime := time.Now().Unix()

		if db.LastActiveTableCheckTime+60 <= unixTime {
			err := db.CheckActiveTable()

			if err != nil {
				log.Println("ERROR: " + err.Error())
			}
		}

		query := r.URL.Query()
		languages := query["language"]

		for _, lang := range languages {
			if lang == "" {
				http.Error(w, "Url Param 'language' is empty", http.StatusBadRequest)
				return
			}
		}

		pagingSeeds := query["seed"]
		pagingSeed := 0

		if len(pagingSeeds) > 0 && pagingSeeds[0] != "" {
			pagingSeed, err = strconv.Atoi(pagingSeeds[0])

			if err != nil {
				http.Error(w, "Url Param 'seed' is invalid", http.StatusBadRequest)
				return
			}
		}

		pagingOffsets := query["offset"]
		pagingOffSet := 0

		if len(pagingOffsets) > 0 && pagingOffsets[0] != "" {
			pagingOffSet, err = strconv.Atoi(pagingOffsets[0])

			if err != nil {
				http.Error(w, "Url Param 'offset' is invalid", http.StatusBadRequest)
				return
			}
		}

		issues, nextPagingSeed, nextPagingOffset, err := queryIssues(db, languages, pagingSeed, pagingOffSet)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}

		resp := IssueResponse{
			Issues:     issues,
			PagingSeed: nextPagingSeed,
			NextOffset: nextPagingOffset,
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
		log.Println("Received /gitdata")

		token := r.Header.Get("x-token")
		envToken := os.Getenv("GIT-PATH-TOKEN")

		if token != envToken {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		nextActiveTable := "issues" + strconv.FormatInt(time.Now().Unix(), 10)

		err := db.CreateNextIssuesTable(&nextActiveTable)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		lastCursor := ""
		lastPageFound := false
		insertedTotal := 0

		for !lastPageFound {
			resp, err := common.GetIssuesFromGithub(lastCursor)

			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if len(resp.Data.Search.Edges) == 0 {
				lastPageFound = true
			} else {
				lastRepo := resp.Data.Search.Edges[len(resp.Data.Search.Edges)-1]
				lastCursor = lastRepo.Cursor

				inserted, err := insertIssuesIntoDB(&nextActiveTable, db, resp)

				if err != nil {
					log.Println(err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				insertedTotal += inserted
			}
		}

		respString := "Inserted " + strconv.Itoa(insertedTotal) + " issues into database"

		err = db.AddActiveTableEntry(&nextActiveTable)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = db.DeleteOldIssuesTables()

		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(respString)
		w.Write([]byte(respString))
	})

	if os.Getenv("ENV") == "LOCAL" {
		log.Fatal(http.ListenAndServe(":80", nil))
	} else {
		log.Println("Starting https listener")

		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("cloud.fixthepla.net"),
			Cache:      autocert.DirCache("/tmp/certs"),
		}
		server := &http.Server{
			Addr: ":https",
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}

		go func() {
			h := certManager.HTTPHandler(nil)
			log.Fatal(http.ListenAndServe(":http", h))
		}()

		log.Fatal(server.ListenAndServeTLS("", ""))
	}
}

func insertIssuesIntoDB(table *string, db *common.Database, resp *common.GithubResponse) (int, error) {
	stmt, err := db.SQLDB.Prepare("INSERT INTO " + *table + " VALUES( ?, ?, ?, ?, ?, ? )")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	count := 0
	for _, repo := range resp.Data.Search.Edges {
		for _, issue := range repo.Node.Issues.Edges {

			labels := make([]IssueLabel, len(issue.Node.Labels.Nodes))
			for i, issueLabel := range issue.Node.Labels.Nodes {
				labels[i] = IssueLabel{
					Name:  issueLabel.Name,
					Color: issueLabel.Color,
				}
			}

			labelBytes, err := json.Marshal(labels)

			if err != nil {
				log.Println("Error marshalling issue labels: " + err.Error())
				continue
			}

			_, err = stmt.Exec(nil, issue.Node.Title, repo.Node.NameWithOwner, issue.Node.Number, repo.Node.PrimaryLanguage.Name, labelBytes)
			if err != nil {
				log.Println("Error inserting issue: " + err.Error())
			} else {
				count++
			}
		}
	}

	return count, nil
}

func queryIssues(db *common.Database, languages []string, pagingSeed int, pagingOffset int) (nextIssues []Issue, nextPagingSeed int, nextPagingOffset int, err error) {
	var queryStringBuilder strings.Builder

	queryStringBuilder.WriteString("SELECT * FROM ")
	queryStringBuilder.WriteString(db.ActiveTable)

	for i := 0; i < len(languages); i++ {
		if i == 0 {
			queryStringBuilder.WriteString(" WHERE language = ?")
		} else {
			queryStringBuilder.WriteString(" OR language = ?")
		}
	}

	pageSize := 20

	if pagingSeed == 0 {
		pagingSeed = rand.Int()
	}

	queryStringBuilder.WriteString(" ORDER BY RAND(")
	queryStringBuilder.WriteString(strconv.Itoa((pagingSeed)))
	queryStringBuilder.WriteString(") LIMIT ")
	queryStringBuilder.WriteString(strconv.Itoa((pageSize)))

	queryStringBuilder.WriteString(" OFFSET ")
	queryStringBuilder.WriteString(strconv.Itoa((pagingOffset)))

	stmt, err := db.SQLDB.Prepare(queryStringBuilder.String())

	if err != nil {
		return nil, 0, 0, err
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
		return nil, 0, 0, err
	}

	defer rows.Close()

	issues := make([]Issue, 0)
	issueID := ""

	count := 0
	for rows.Next() {
		issue := Issue{}
		var labelBytes []byte

		err = rows.Scan(&issueID, &issue.Title, &issue.Repo, &issue.Number, &issue.Language, &labelBytes)

		if err != nil {
			return nil, 0, 0, err
		}

		err = json.Unmarshal(labelBytes, &issue.Labels)

		if err != nil {
			return nil, 0, 0, err
		}

		issues = append(issues, issue)
		count++
	}

	err = rows.Err()

	if err != nil {
		return nil, 0, 0, err
	}

	pagingOffset += pageSize

	return issues, pagingSeed, pagingOffset, nil
}
