package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type GetCodeRequest struct {
	Repo   string
	Branch string
	Files  []string
}

func main() {

	http.HandleFunc("/", GetHello)

	http.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
		var requestString []byte
		var reqestJson GetCodeRequest

		requestString, err := io.ReadAll(r.Body)
		
		if err != nil {
			log.Fatalf("Something went wrong %v", err)
		}
		
		err = json.Unmarshal(requestString, &reqestJson)
		
		if err != nil {
			log.Fatalf("Something went wrong %v", err)
		}
		if ! strings.Contains(reqestJson.Repo, "github.com/"){
			fmt.Fprintf(w, "%v is not a valid repo link",reqestJson.Repo)
			return
		}
		destination := fmt.Sprintf("https://%v/blob/%v/%v", reqestJson.Repo, reqestJson.Branch, reqestJson.Files[0])
		fmt.Fprint(w, destination)
	})

	http.ListenAndServe(":3000", nil)
}
