// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"html/template"
)

var clientID string
var clientSecret string
var callbackURL string

func main() {
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/notify", notifyHandler)
	http.HandleFunc("/auth", authHandler)
	clientID = os.Getenv("ClientID")
	clientSecret = os.Getenv("ClientSecret")
	callbackURL = os.Getenv("CallbackURL")
	port := os.Getenv("PORT")
	fmt.Printf("ENV port:%s, cid:%s csecret:%s\n", port, clientID, clientSecret)
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Populates request.Form
	code := r.Form.Get("code")
	state := r.Form.Get("state")
	fmt.Printf("Get code=%s, state=%s \n", code, state)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {

}
func authHandler(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").ParseFiles("auth.tmpl")
	check(err)
	noItems := struct {
		ClientID    string
		CallbackURL string
	}{
		ClientID:    clientID,
		CallbackURL: callbackURL,
	}

	err = t.Execute(w, noItems)
	check(err)
}
