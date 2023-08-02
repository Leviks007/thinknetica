package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

func OpenWeb(fnmap map[string]func(http.ResponseWriter, *http.Request)) {
	mux := mux.NewRouter()
	for key, fn := range fnmap {
		mux.HandleFunc(key, fn).Methods(http.MethodGet)
	}
	http.ListenAndServe(":8080", mux)
}
