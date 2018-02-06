package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HttpTestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "POST" {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}
		if r.Body != nil {
			jsonResponse := map[string]string{}
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &jsonResponse)
			if jsonResponse["email"] == "testUser" && jsonResponse["password"] == "testPasswd" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("pass"))
			} else {
				fmt.Printf("the email is %v\n ", jsonResponse["email"])
				fmt.Printf("the password is %v\n ", jsonResponse["password"])
			}
		}
	}))
	defer ts.Close()

	data := map[string]string{}

	// i := Aiservice{
	// 	Username: "testUser",
	// 	Password: "testPasswd",
	// }
	data["email"] = "Username"
	data["password"] = "Password"

	payloadBytes, _ := json.Marshal(data)
	body := bytes.NewReader(payloadBytes)

	res, err := http.Post(ts.URL, "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	respbody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "pass", string(respbody))

}

func main() {
	var t *testing.T
	HttpTestServer(t)
}
