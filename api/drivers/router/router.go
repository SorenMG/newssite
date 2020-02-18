package router

import (
	"fmt"
	"log"
	"net/http"
	"newssite/api/middlewares/logger"
	sitesRoute "newssite/api/routes/sites"
	"strings"

	"github.com/gorilla/mux"
)

type Router struct {
	Inst *mux.Router
	Port string
}

func New() (r Router) {
	r = Router{
		mux.NewRouter(),
		"8000",
	}

	setupRoutes(r.Inst)
	setupMiddleWare(r.Inst)

	log.Println("Starting server at port " + r.Port + " 🙌")

	// List routes
	walkRoutes(r.Inst)

	// Run server
	log.Fatal(http.ListenAndServe(":"+r.Port, r.Inst))

	return
}

func setupRoutes(r *mux.Router) {
	// Handle routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	sitesRoute.Init(r.PathPrefix("/sites").Subrouter())
}

func setupMiddleWare(r *mux.Router) {
	// Use middleware
	r.Use(logger.Logger)
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
