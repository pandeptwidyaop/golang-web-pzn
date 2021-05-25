package webnew

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintf(w, "Hello %s", "Uname")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=pande", nil)

	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameters(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s\n", firstName, lastName)
}

func TestMultipleParameters(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/name?first_name=pande&last_name=ganteng", nil)

	recorder := httptest.NewRecorder()

	MultipleParameters(recorder, request)

	response := recorder.Result()

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultiValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	names := query["name"]

	fmt.Fprint(w, strings.Join(names, ","))
}

func TestMultiValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/name?name=Pande&name=Putu&name=widya", nil)

	recorder := httptest.NewRecorder()

	MultiValues(recorder, request)

	response := recorder.Result()

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}
