package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GithubResponse ..
type GithubResponse struct {
	Data struct {
		Search struct {
			Edges []GithubRepository `json:"edges"`
		} `json:"search"`
		RateLimit struct {
			Cost      int `json:"cost"`
			Limit     int `json:"limit"`
			Remaining int `json:"remaining"`
		} `json:"rateLimit"`
	} `json:"data"`
}

// GithubRepository ..
type GithubRepository struct {
	Node struct {
		Issues struct {
			Edges      []GithubIssue `json:"edges"`
			TotalCount int           `json:"totalCount"`
		}
		NameWithOwner   string `json:"nameWithOwner"`
		PrimaryLanguage struct {
			Name string `json:"name"`
		} `json:"primaryLanguage"`
	} `json:"node"`
	Cursor string `json:"cursor"`
}

// GithubIssue ..
type GithubIssue struct {
	Node struct {
		Title  string `json:"title"`
		Number int    `json:"number"`
		Labels struct {
			Nodes []GithubIssueLabel `json:"nodes"`
		} `json:"labels"`
	} `json:"node"`
}

// GithubIssueLabel ..
type GithubIssueLabel struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// CursorVar ..
type CursorVar struct {
	LastCursor string `json:"lastCursor"`
}

// GetIssuesFromGithub ..
func GetIssuesFromGithub(lastCursor string) (*GithubResponse, error) {

	var mapData map[string]string

	if lastCursor == "" {
		fileData, err := ioutil.ReadFile("../queries/get-issues.gql")

		if err != nil {
			return nil, err
		}

		mapData = map[string]string{
			"query":     string(fileData),
			"variables": "",
		}
	} else {
		fileData, err := ioutil.ReadFile("../queries/get-issues-paginated.gql")

		if err != nil {
			return nil, err
		}

		queryVar := CursorVar{
			LastCursor: lastCursor,
		}

		queryVarJSON, err := json.Marshal(queryVar)

		if err != nil {
			return nil, err
		}

		mapData = map[string]string{
			"query":     string(fileData),
			"variables": string(queryVarJSON),
		}
	}

	requestBody, err := json.Marshal(mapData)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	req.Header.Add("Authorization", "bearer "+os.Getenv("GITHUB-TOKEN"))
	client := &http.Client{}

	log.Println("Getting data from github")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Println(string(bodyBytes[:]))
		return nil, errors.New(string(bodyBytes[:]))
	}

	//log.Println(string(bodyBytes[:len(bodyBytes)]))

	githubResp := &GithubResponse{}
	err = json.Unmarshal(bodyBytes, githubResp)

	if err != nil {
		return nil, err
	}

	log.Println("Received data from github. Used ratelimit: ", githubResp.Data.RateLimit.Cost, " remaining: ", githubResp.Data.RateLimit.Remaining)

	return githubResp, nil
}
