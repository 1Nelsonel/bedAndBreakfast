package main

import (
	"fmt"
	"testing"

	"github.com/1Nelsonel/bedAndBreakfast/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	// logic here
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do mithing; test pass

	default:
		t.Error(fmt.Sprintf("Typr id no *chi.Mux, type is %Y", v))
	}
}