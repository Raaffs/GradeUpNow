package main

import (
	"net/http"

	"github.com/gorilla/mux"
)
// This method returns a servemux containing our application routes.
func (app *application) routes() *mux.Router {
    mux := mux.NewRouter()
    //tbh I don't know what this exactly does but without this
    //css doesn't work
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/login", app.login) 
    mux.HandleFunc("/signup", app.sign_up)
	mux.HandleFunc("/home/{username}/stats", app.get_usr_stats)

    mux.HandleFunc("/home/{subject}/{type}", app.q_type_handler)

    return mux
}