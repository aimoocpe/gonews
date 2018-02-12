package app

import (
	"net/http"
)

//Mount handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // list all news
	// /news/:path
	// /admin/login
	// /admin/posts
	// /admin/edit
	// /admin/create
}
