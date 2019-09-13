package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stanleynguyen/git-comment/comment-app/domain"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`DROP TABLE IF EXISTS comments; DROP TABLE IF EXISTS schema_migrations`)
	if err != nil {
		log.Fatal(err)
	}
	testStartSignal := make(chan bool)
	go startInTest(testStartSignal)
	<-testStartSignal
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGetFromOrg(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	resp, err := http.Get("http://localhost:5000/orgs/xendit/comments")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	respMap, err := getCommentsForOrg("xendit")
	if err != nil {
		t.Fatal(err.Error())
	}

	if comments, ok := respMap["comments"]; ok {
		if len(comments) != 2 {
			t.Errorf("len(comments): EXPECTED=%d, GOT=%d", 2, len(comments))
		}
		for i, c := range comments {
			if c.Org != "xendit" {
				t.Errorf("comments[%d].Org: EXPECTED=%s, GOT=%s", i, "xendit", c.Org)
			}
		}
	} else {
		t.Error("Expect response to have comments array")
	}
}


func TestPostComment(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	expectedComment := "one of the Big Four"

	resp, err := http.PostForm("http://localhost:5000/orgs/google/comments", url.Values{"comment": {expectedComment}})
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		t.Fatalf("Request failed with status %d", resp.StatusCode)
	}

	c := domain.Comment{}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal(bodyBytes, &c)
	if err != nil {
		t.Fatal(err.Error())
	}

	if c.Comment != expectedComment {
		t.Errorf("c.Comment: EXPECTED='%s', GOT='%s'", expectedComment, c.Comment)
	}
	if c.Org != "google" {
		t.Errorf("c.Org: EXPECTED='%s', GOT='%s'", "google", c.Org)
	}

	respMap, err := getCommentsForOrg("google")
	if err != nil {
		t.Fatal(err.Error())
	}
	if comments, ok := respMap["comments"]; ok {
		if len(comments) != 2 {
			t.Errorf("len(comments): EXPECTED=%d, GOT=%d", 2, len(comments))
		}
		for i, c := range comments {
			if c.Org != "google" {
				t.Errorf("comments[%d].Org: EXPECTED=%s, GOT=%s", i, "google", c.Org)
			}
		}
	} else {
		t.Error("Expect response to have comments array")
	}
}

func TestDeleteComment(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	cli := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:5000/orgs/xendit/comments", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp, err := cli.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		t.Fatalf("Request failed with status %d", resp.StatusCode)
	}

	respMap, err := getCommentsForOrg("xendit")
	if err != nil {
		t.Fatal(err.Error())
	}
	if comments, ok := respMap["comments"]; ok {
		if len(comments) != 0 {
			t.Errorf("len(comments): EXPECTED=%d, GOT=%d", 2, len(comments))
		}
	} else {
		t.Error("Expect response to have comments array")
	}
}

func setup(t *testing.T) (teardown func()) {
	_, err := db.Exec(
		`INSERT INTO comments(org, comment) VALUES ($1, $2), ($3, $4), ($5, $6)`,
		"xendit", "Looking to hire SE Asia's top dev talent!",
		"xendit", "Accept and Detect Online Payments",
		"google", "Don't be evil.")
	if err != nil {
		t.Fatal(err.Error())
	}

	return func () {
		_, err := db.Exec(`DELETE FROM comments`)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func getCommentsForOrg(org string) (map[string][]domain.Comment, error) {

	resp, err := http.Get(fmt.Sprintf("http://localhost:5000/orgs/%s/comments", org))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respMap := map[string][]domain.Comment{}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &respMap)
	if err != nil {
		return nil, err
	}

	return respMap, nil
}
