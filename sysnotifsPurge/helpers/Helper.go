package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//SetResponse function with json output
func SetResponse(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if i != nil {
		if err := json.NewEncoder(w).Encode(i); err != nil {
			panic(err)
		}
	}
}

//SetTextResponse text output with header
func SetTextResponse(w http.ResponseWriter, statusCode int, resp string) {

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(statusCode)

	w.Write([]byte(resp))
}

//AppendTextResponse add text to response
func AppendTextResponse(w http.ResponseWriter, resp string) {

	w.Write([]byte(resp))
}

//GetID function
func GetID(r *http.Request, idName string) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[idName])

	if err != nil {
		panic(err)
	}

	return id
}

//GetString function
func GetString(r *http.Request, stringName string) string {
	vars := mux.Vars(r)
	stringValue := vars[stringName]

	return stringValue
}

//GetIsActive function
func GetIsActive(r *http.Request) bool {
	vars := mux.Vars(r)
	isactive, err := strconv.ParseBool(vars["isactive"])

	if err != nil {
		panic(err)
	}

	return isactive
}

//GetBody function
func GetBody(r *http.Response, i interface{}) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	json.Unmarshal(body, &i)

}
