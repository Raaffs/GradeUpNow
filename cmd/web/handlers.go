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
func (app *application) login(w http.ResponseWriter, r *http.Request) {
    // Check if it's a POST request
    if r.Method == http.MethodPost {
        // Retrieve the username and password from the form
        username := r.FormValue("username")
        password := r.FormValue("password")

        if username == "" || password == "" {
            // Invalid input, set error message
            errorMessage := "Username and password cannot be empty."
            http.Redirect(w, r, "/login?error="+errorMessage, http.StatusSeeOther)
            return
        }

        // Check if the username and password exist in the database
        user, err := app.user.Get(username)
        if err != nil {
            if errors.Is(err, models.ErrNoRecord) {
                errorMessage := "User not found."
                http.Redirect(w, r, "/login?error="+errorMessage, http.StatusSeeOther)
                return
            } else {
                app.serverError(w, err)
                return
            }
        }

        // Check if the entered password matches the user's password
        if user.Password != password {
            errorMessage := "Invalid password."
            http.Redirect(w, r, "/login?error="+errorMessage, http.StatusSeeOther)
            return
        }

        // Authentication successful, redirect to home page
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // Render the login form
    files := []string{
        "./ui/html/login.html",
        "./ui/html/index.html",
    }
    tmpl, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    // Get the error message from the query parameter (if any)
    errorMessage := r.URL.Query().Get("error")
    data := struct {
        ErrorMessage string
    }{
        ErrorMessage: errorMessage,
    }

    err = tmpl.ExecuteTemplate(w, "login.html", data)
    if err != nil {
        app.serverError(w, err)
        return
    }
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
	_,err:=app.user.SignUp(username,email,pass)
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

func (app *application)leader_board(w http.ResponseWriter, r *http.Request){
	leader_board,err:=app.user.Leader_board()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _,usr:=range leader_board{
		fmt.Fprintf(w,"%v\n",usr)
	}
}

/*func (app *application)q_type_handler(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
//	subject:=vars["subject"]
	q_type:=vars["type"]
	if q_type=="mcq"{
		mcq_list,err:=app.user.Get_Mcq(q_type)
		if err!=nil
	
	}
	
}*/