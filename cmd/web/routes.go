package main

import (
	"net/http"

	"github.com/gorilla/mux"
)
// This method returns a servemux containing our application routes.
func (app *application) routes() *mux.Router {
    mux := mux.NewRouter()
    //access files in /static/ to render css, javascript and images. 
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))
    
    mux.HandleFunc("/",app.root_hander)
    mux.HandleFunc("/login", app.login)
    mux.HandleFunc("/signup", app.sign_up)
    mux.HandleFunc("/home", app.home)
    mux.HandleFunc("/home/leaderboard",app.leader_board)
    mux.HandleFunc("/home/profile",app.profile_handler)
	mux.HandleFunc("/home/{username}/profile", app.profile_handler)
    mux.HandleFunc("/home/{subject}", app.select_type)
    mux.HandleFunc("/home/{subject}/mcq",app.mcq_handler)
    mux.HandleFunc("/home/{subject}/theory",app.theory_handler)
    mux.HandleFunc("/check-answer",app.Update_mcq_score)
    
    return mux
}