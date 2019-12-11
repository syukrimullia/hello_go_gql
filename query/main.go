package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"

	"github.com/syukrimullia/hello_go_gql/query/src"
)

func main() {
	schema := src.RegisterTypes()

	// Query
	query := `
        {
			tutorial (id: 4) {
				id
				title
				author {
					name
				}
			}
        }
    `
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
