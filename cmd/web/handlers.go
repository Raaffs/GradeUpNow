package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"github.com/Suy56/GradeUpNow/internal/models"
	"github.com/gorilla/mux"
    //"encoding/json"
	//"strconv"
)
func(app *application)root_hander(w http.ResponseWriter, r*http.Request){
    http.Redirect(w,r,"/login",http.StatusSeeOther)
}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/login" {
    	app.notFound(w)
		return
	}
    // Check if it's a POST request
    if r.Method == http.MethodPost {
        // Retrieve the username and password from the form
        username := r.FormValue("username")
        password := r.FormValue("password")
        models.G_CurrentUserSession=username
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
        http.Redirect(w, r, "/home", http.StatusSeeOther)
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
    leaderboard,err:=app.user.Leader_board()
    if err!=nil{
        app.serverError(w,err)
        app.errorLog.Print(err)
    }
    top_five:=leaderboard[:5]

    funcMap := template.FuncMap{
        "add": add,
    }

    data:=struct{
        LeaderBoard [] *models.User
    }{
        LeaderBoard: top_five,
    }
	files := []string{
		"./ui/html/home.html",
	}
    //To access add(a,b int) func. Refer to comments in helpers.go
    tmpl, err := template.New("home.html").Funcs(funcMap).ParseFiles(files...)
    if err != nil {
        app.notFound(w)
    }	
	err=tmpl.ExecuteTemplate(w,"home.html",data)
	if err!=nil{
        app.errorLog.Print(err)
		app.serverError(w,err)
	}
	
}
func (app *application) sign_up(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Retrieve form values
        email := r.FormValue("email")
        username := r.FormValue("username")
        pass := r.FormValue("password")

        if email == "" || username == "" || pass == "" {
            errorMessage := "Email, display name, and password cannot be empty."
            http.Redirect(w, r, "/signup?error="+errorMessage, http.StatusSeeOther)
            return
        }
        

        // Check if username already exists
        usernameExists, err := app.user.Check_if_exist(username, "")
        if err != nil {
            app.serverError(w, err)
            return
        }
        if usernameExists {
            errorMessage := "Username already exists."
            http.Redirect(w, r, "/signup?error="+errorMessage, http.StatusSeeOther)
            return
        }

        // Check if email already exists
        emailExists, err := app.user.Check_if_exist("", email)
        if err != nil {
            app.serverError(w, err)
            return
        }
        if emailExists {
            errorMessage := "Email already exists."
            http.Redirect(w, r, "/signup?error="+errorMessage, http.StatusSeeOther)
            return
        }

        // Proceed with user registration
        _, err = app.user.SignUp(username, email, pass)
        if err != nil {
            app.serverError(w, err)
            return
        }

        // Redirect to the login page after successful signup
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Render the signup form for GET requests
    files := []string{
        "./ui/html/signup.html",
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

    err = tmpl.ExecuteTemplate(w, "signup.html", data)
    if err != nil {
        app.serverError(w, err)
        return
    }
}

func (app *application)profile_handler(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	username:=vars["username"]
    data:=&template_data{}
    if username!=""{
        usr,err:=app.user.Get(username)
        if err!=nil{
            fmt.Fprint(w,"Please login")
            http.Redirect(w,r,"/login",http.StatusSeeOther)
        }
        data.Individual_user_data=usr
        tmpl,err:=template.ParseFiles("./ui/html/profile.html")
        if err!=nil{            

            app.serverError(w,err)
            app.errorLog.Fatal(err)
            
        }
        err=tmpl.ExecuteTemplate(w,"profile.html",data)
        if err!=nil{
            app.serverError(w,err)
            app.errorLog.Fatal(err)
        }
    }
    
    
    if r.URL.Path=="/home/profile"{
        fmt.Print(models.G_CurrentUserSession)
        usr,err :=app.user.Get(models.G_CurrentUserSession)
        data.Individual_user_data=usr
        if err!=nil{
        app.notFound(w)
        app.errorLog.Println(err)
       }
       tmpl,err:=template.ParseFiles("./ui/html/profile.html")
       if err!=nil{
           app.serverError(w,err)
           app.errorLog.Print(err)
       }
       err=tmpl.ExecuteTemplate(w,"profile.html",data)
    }

}

func (app *application)leader_board(w http.ResponseWriter, r *http.Request){
    leaderboard,err:=app.user.Leader_board()
    if err!=nil{
        app.serverError(w,err)
        app.errorLog.Print(err)
    }
    funcMap := template.FuncMap{
        "add": add,
    }

    data:=struct{
        LeaderBoard [] *models.User
    }{
        LeaderBoard: leaderboard,
    }
	files := []string{
		"./ui/html/leaderboard.html",
	}
    
    //To access add(a,b int) func. Refer to comments in helpers.go
    tmpl, err := template.New("leaderboard.html").Funcs(funcMap).ParseFiles(files...)
    if err != nil {
        app.notFound(w)
    }	
	err=tmpl.ExecuteTemplate(w,"leaderboard.html",data)
	if err!=nil{
        app.errorLog.Print(err)
		app.serverError(w,err)
	}

}

func (app *application)select_type(w http.ResponseWriter,r *http.Request){
    vars:=mux.Vars(r)
    subject:=vars["subject"]
    tmpl,err:=template.ParseFiles("./ui/html/display.html")
    if err!=nil{
        app.serverError(w,err)
        app.errorLog.Print(err)
        return
    }
    //to redirect to /home/{subject}/{type}
    data := struct {
        Subject string
    }{
        Subject: subject,
    }
    tmpl.ExecuteTemplate(w,"display.html",data)
    
}
