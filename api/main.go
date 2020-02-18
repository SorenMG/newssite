package main

import (
	"fmt"
	"log"
	"strings"

	"net/http"

	"newssite/api/middlewares/logger"
	sitesRoute "newssite/api/routes/sites"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func initRouter() {
	r := mux.NewRouter()

	// Handle routes
	r.HandleFunc("/", HomeHandler)
	sitesRoute.Init(r.PathPrefix("/sites").Subrouter())

	// Use middleware
	r.Use(logger.Logger)

	walkRoutes(r)

	// Serve server
	log.Fatal(http.ListenAndServe(":8000", r))
}

func walkRoutes(r *mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	mongo := New()

}
