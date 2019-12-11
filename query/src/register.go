package src

import (
	"log"

	"github.com/graphql-go/graphql"
)

func RegisterTypes() graphql.Schema {
	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			// we define the name and the fields of our
			// object. In this case, we have one solitary
			// field that is of type string
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"tutorials": &graphql.Field{
					// we'll use NewList to deal with an array
					// of int values
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	var tutorialType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"author": &graphql.Field{
					// here, we specify type as authorType
					// which we've already defined.
					// This is how we handle nested objects
					Type: authorType,
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(commentType),
				},
			},
		},
	)
	tutorials := populateTutorial()

	// Schema
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type: tutorialType,
			// it's good form to add a description
			// to each field.
			Description: "Get Tutorial By ID",
			// We can define arguments that allow us to
			// pick specific tutorials. In this case
			// we want to be able to specify the ID of the
			// tutorial we want to retrieve
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// take in the ID argument
				id, ok := p.Args["id"].(int)
				if ok {
					// Parse our tutorial array for the matching id
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							// return our tutorial
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		// this is our `list` endpoint which will return all
		// tutorials available
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}
