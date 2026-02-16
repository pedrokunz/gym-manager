package graphql

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/pedrokunz/gym-manager/backend/internal/db"
)

var classType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Class",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int},
		"name":     &graphql.Field{Type: graphql.String},
		"trainer":  &graphql.Field{Type: graphql.String},
		"schedule": &graphql.Field{Type: graphql.String},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"classes": &graphql.Field{
			Type: graphql.NewList(classType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rows, err := db.DB.Query("SELECT id, name, trainer, schedule FROM classes")
				if err != nil {
					return nil, err
				}
				defer rows.Close()

				var results []map[string]interface{}
				for rows.Next() {
					var id int
					var name, trainer, schedule string
					rows.Scan(&id, &name, &trainer, &schedule)
					results = append(results, map[string]interface{}{
						"id":       id,
						"name":     name,
						"trainer":  trainer,
						"schedule": schedule,
					})
				}
				return results, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

func Handler(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query string `json:"query"`
	}
	json.NewDecoder(r.Body).Decode(&params)

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: params.Query,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
