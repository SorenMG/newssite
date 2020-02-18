package sites

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	site := Site{
		1,
		"bt",
		"http://www.bt.dk/",
		"da_DK",
		SiteScrapeDef{
			"pub",
			"edit",
			"desc",
			"title",
			"tag",
			"section",
			"imgurl",
			Dimension{
				50,
				60,
			},
		},
	}

	if err := json.NewEncoder(w).Encode(site); err != nil {
		panic(err)
	}
}

func Init(r *mux.Router) {
	r.HandleFunc("/", HandleHome)
}
