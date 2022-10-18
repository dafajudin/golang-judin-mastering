package main

import (
	"gql-hero/database"
	"gql-hero/graph"
	"gql-hero/graph/generated"
	"net/http"
	"testing"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/steinfletcher/apitest"
)

func graphQLHandler() http.Handler {
	database.Connect()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	return mux
}

func TestGetHeroes_Success(t *testing.T) {
	var query string = `{
		heroes {
			id
			nama_hero
			role
			emblem
			battle_spell
		}
	}`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateHero_Success(t *testing.T) {
	var query string = `mutation {
		createHero(input: {
			nama_hero: "new hero",
			role: "mage",
			emblem: "assasin",
			battle_spell: "inspire"
		}) {
			nama_hero
		}
	}`

	var result string = `{
		"data": {
			"createHero": {
				"nama_hero": "new hero"
			}
		}
	}`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusOK).
		Body(result).
		End()
}

func TestCreateHero_Failed(t *testing.T) {
	var query string = `mutation {
		createHero(input: {
			nama_hero: "new hero",
		}) {
			id
			nama_hero
		}
	}`

	apitest.New().
		Handler(graphQLHandler()).
		Post("/query").
		GraphQLQuery(query).
		Expect(t).
		Status(http.StatusUnprocessableEntity).
		End()
}
