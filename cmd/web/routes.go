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
    
    mux.HandleFunc("/login", app.login) 
    mux.HandleFunc("/home", app.home)
    mux.HandleFunc("/signup", app.sign_up)
    mux.HandleFunc("/home/leaderboard",app.leader_board)
    mux.HandleFunc("/home/profile",app.profile_handler)
	mux.HandleFunc("/home/{username}/profile", app.profile_handler)
    mux.HandleFunc("/home/{subject}/{type}", app.q_type_handler)
    mux.HandleFunc("/",app.root_hander)
    return mux
}