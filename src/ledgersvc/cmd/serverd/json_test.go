package main

import (
	"net/http"
	"sort"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func Test_jsonAPIHandler(t *testing.T) {
	// Given:
	r := chi.NewRouter()

	// When:
	jsonAPIHandler(r)

	// Then:
	var routesFound []string
	require.NoError(t, chi.Walk(
		r,
		func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			routesFound = append(routesFound, method+" "+route)
			return nil
		},
	))
	sort.Strings(routesFound)

	routesExp := []string{
		"GET /abc",
	}
	sort.Strings(routesExp)

	require.EqualValues(t, routesExp, routesFound)
}
