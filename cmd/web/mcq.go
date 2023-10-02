package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Suy56/GradeUpNow/internal/models"
	"github.com/gorilla/mux"
)
var currentQuestionIndex=1
var serverVariable=0
func (app *application)mcq_handler(w http.ResponseWriter,r *http.Request){
    vars:=mux.Vars(r)
    subject:=vars["subject"]
    _,err:=app.user.Get(models.G_CurrentUserSession)
    //TODO:
    //scr+=10 for every correct ans
    //update user databse using app.user.Update_score(subject,scr)
    if err!=nil{
        app.serverError(w,err)
        app.errorLog.Print("Error in mcq_hander while request user data",err)
    }
    question,err:=app.user.Get_Mcq(subject)
    if err!=nil{
        app.serverError(w,err)
        app.errorLog.Print("Error in mcq_handler while requesting mcq data",err)
    }
    if r.Method==http.MethodPost{
        userChoice,err:=strconv.Atoi(r.FormValue("choice"))
        if err!=nil{
            app.serverError(w,err)
            app.errorLog.Print("error in mcq_handler",err)
        }
        question[currentQuestionIndex].UserChoice=userChoice
    }
    tmpl,err:=template.ParseFiles("./ui/html/mcq.html")
    if err!=nil{
        app.serverError(w,err)
        app.errorLog.Print("Error in mcq_hanlder while paring html",err)
    }
    err=tmpl.Execute(w,question[currentQuestionIndex])
    if err!=nil{
        app.serverError(w,err)
        app.errorLog.Fatal(err)
    }
}

