package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stanleynguyen/git-comment/comment-app/domain"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
)

// RenderJSON return json object in the http response
func RenderJSON(w http.ResponseWriter, obj interface{}) {
	response, err := json.Marshal(&obj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// RenderErr return error in http response
func RenderErr(w http.ResponseWriter, err error) {
	log.Print(err.Error())

	if httpErr, ok := err.(*domain.HTTPError); ok {
		statusCode := httpErr.StatusCode
		message, _ := json.Marshal(map[string]string{"message": err.Error()})
		http.Error(w, string(message), statusCode)
		return
	}

	http.Error(w, "Something went wrong", 500)
}

// ReadRequestBody read request body values and bind to the interface
func ReadRequestBody(r *http.Request, i interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		if err := r.ParseForm(); err != nil {
			return err
		}
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		fmt.Print(r.PostForm)
		if err := decoder.Decode(i, r.PostForm); err != nil {
			return err
		}
	} else if contentType == "application/json" {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(i); err != nil {
			return err
		}
	} else if strings.Contains(contentType, "multipart/form-data") {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			return err
		}
		decoder := schema.NewDecoder()
		if err := decoder.Decode(i, r.PostForm); err != nil {
			return err
		}
	} else {
		return errors.New("Content-Type Not Accepted")
	}
	return nil
}
