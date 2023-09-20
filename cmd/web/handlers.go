package main

import (
	//"fmt"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Suy56/GradeUpNow/internal/models"
	"github.com/gorilla/mux"
	//"strconv"
)
func (app *application)login(w http.ResponseWriter, r *http.Request) {
	files := []string{
	"./ui/html/login.html",
	"./ui/html/index.html",
	}
	tmpl,err:=template.ParseFiles(files...)

	if err!=nil{
		app.serverError(w,err)
		return
	}
	err=tmpl.ExecuteTemplate(w,"login.html",nil)
	if err!=nil{
		app.serverError(w,err)
		return
	}
	if r.Method==http.MethodPost{
		name:=r.FormValue("username")
		pass:=r.FormValue("password")
		fmt.Print(name,pass)
	}
	
	http.Redirect(w,r,"/",http.StatusSeeOther)

}
func (app *application)home(w http.ResponseWriter, r *http.Request){
	files := []string{
		"./ui/html/index.html",
		"./ui/html/login.html",
	}
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	tmpl,err:=template.ParseFiles(files...)
	if err!=nil{
		app.notFound(w)
	}
	
	err=tmpl.ExecuteTemplate(w,"index.html",nil)
	if err!=nil{
		app.serverError(w,err)
	}
	
}
func (app *application)sign_up(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		return
	}
	username:="Suyash"
	email:="suyash353@gmail.com"
	pass:="gtrwm343"
	_,err:=app.user.SignIn(username,email,pass)
	if err!=nil{
		app.serverError(w,err)
		return
	}
}

func (app *application)get_usr_stats(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	username:=vars["username"]
	fmt.Print(username)
	fmt.Println(username)
	user,err:=app.user.Get(username)
	if err!=nil{
		if errors.Is(err, models.ErrNoRecord){
			app.notFound(w)
		}else{
			app.serverError(w,err)
		}
		return
	}
	fmt.Fprintln(w, "username:", user.Username,"Theory:",user.Theory_score,"Mcq:",user.Mcq_score,"Total:",user.Total_score)
}

func (app *application)type_handler(w http.ResponseWriter,r *http.Request){
	//vars:=mux.Vars(r)
	//subject:=vars["subject"]
	//q_type:=vars["type"]
	
}