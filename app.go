package main

import (
	"github.com/graphql-go/graphql"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func main(){
	http.HandleFunc("/graphql",middleware(CreateSchema()))
	
	http.ListenAndServe(":8080",nil)
}

func middleware(schema graphql.Schema)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// extract raw payload from client
		data,err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		
		
		result := graphql.Do(graphql.Params{
			Schema:schema,
			RequestString:string(data),
		})
		
		w.Header().Set("Content-Type","application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}
