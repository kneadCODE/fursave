package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kneadCODE/fursave/src/ledgersvc/internal/controller/jsonapi"
	ledgermgmtrepo "github.com/kneadCODE/fursave/src/ledgersvc/internal/repository/ledgermgmt"
	ledgermgmtuc "github.com/kneadCODE/fursave/src/ledgersvc/internal/usecase/ledgermgmt"
)

func jsonAPIHandler(rtr chi.Router) {
	rtr.Get("/abc", func(w http.ResponseWriter, r *http.Request) {
		log.Println("abc called")
	})

	ctrl := jsonapi.NewController(
		ledgermgmtuc.New(
			ledgermgmtrepo.New(),
		),
	)

	_ = ctrl
}
