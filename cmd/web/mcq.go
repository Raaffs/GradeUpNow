package main

import (
	"html/template"
    "encoding/json"
	"net/http"
	"strconv"
    "fmt"
    "log"

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
func handd(){
    router:=mux.NewRouter()
router.HandleFunc("/next/{subject}", func(w http.ResponseWriter, r *http.Request) {
    // Move to the next question.
    subject := mux.Vars(r)["subject"]
    currentQuestionIndex = (currentQuestionIndex + 1) 
    redirectURL := fmt.Sprintf("/home/%s/mcq", subject) // Construct the redirect URL
    http.Redirect(w, r, redirectURL, http.StatusSeeOther)
})

router.HandleFunc("/check-answer", func(w http.ResponseWriter, r *http.Request) {
    // Parse the JSON request
    var requestData struct {
        IsCorrect bool `json:"isCorrect"`
    }
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Handle the answer (e.g., log it, update the score, etc.)
    // You can also send a response back to the client if needed.
    // For now, we'll just log it.
    if requestData.IsCorrect {
        serverVariable += 10 // Increment the server-side variable by 10
        fmt.Println("User's answer is correct!", serverVariable)
    } else {
        fmt.Println("User's answer is incorrect.", serverVariable)
    }

    // You can send a response back to the client here if needed.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]bool{"isCorrect": requestData.IsCorrect})
})

router.HandleFunc("/get-variable", func(w http.ResponseWriter, r *http.Request) {
    // Respond with the current value of the server-side variable
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]int{"serverVariable": serverVariable})
})

// Start the HTTP server with Gorilla Mux router.
log.Println("Server listening on :8080...")
err := http.ListenAndServe(":8080", router)
if err != nil {
    log.Fatal(err)
}
}