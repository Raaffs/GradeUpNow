package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Suy56/GradeUpNow/internal/models"
	"github.com/gorilla/mux"
)

var score int
var total_score=0
func (app *application) mcq_handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subject := vars["subject"]
	_, err := app.user.Get(models.G_CurrentUserSession)
	if err != nil {
		app.serverError(w, err)
		app.errorLog.Print("Error in mcq_hander while request user data", err)
	}

	question, err := app.user.Get_Mcq(subject)
	if err != nil {
		app.serverError(w, err)
		app.errorLog.Print("Error in mcq_handler while requesting mcq data", err)
	}
	currentQuestionIndex, err := strconv.Atoi(r.URL.Query().Get("q"))
	if err != nil {
		currentQuestionIndex = 0
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			app.serverError(w, err)
			return
		}
		userChoice, err := strconv.Atoi(r.FormValue("choice"))
		if err != nil {
			app.serverError(w, err)
			app.errorLog.Print("error in mcq_handler", err)
		}
		question[currentQuestionIndex].UserChoice = userChoice
	}
	hasNextQuestion := currentQuestionIndex+1 < len(question)
	if currentQuestionIndex >= 0 && currentQuestionIndex < len(question) {
		tmpl, err := template.ParseFiles("./ui/html/mcq.html")
		if err != nil {
			app.serverError(w, err)
			app.errorLog.Print("Error in mcq_hanlder while paring html", err)
		}
		fmt.Println(score)
		
		err = tmpl.Execute(w, struct {
			MQ_num            int
			MQ_question       string
			MQ_ans            int
			MQ_type           string
			Options           []string
			NextQuestionIndex int // Include the next question index
			UserChoice        int
			HasNextQuestion bool
		}{
			MQ_num:            question[currentQuestionIndex].MQ_num,
			MQ_question:       question[currentQuestionIndex].MQ_question,
			MQ_ans:            question[currentQuestionIndex].MQ_ans,
			MQ_type:           question[currentQuestionIndex].MQ_type,
			Options:           question[currentQuestionIndex].Options,
			NextQuestionIndex: currentQuestionIndex + 1, // Pass the next question index
			UserChoice:        question[currentQuestionIndex].UserChoice,
			HasNextQuestion: hasNextQuestion,
		})

		if err != nil {
			app.serverError(w, err)
			app.errorLog.Fatal(err)
		}
        total_score+=score
	} else {
		// Display a message when all questions have been answered
		fmt.Fprint(w, "All questions have been answered.")
	}
	err=app.user.Update_score(subject,score)
	if err!=nil{
		app.serverError(w,err)
		app.errorLog.Println("in mcq handler",err)
	}

    fmt.Println("total score",total_score)
      fmt.Println("score",score)

}

func (app *application) Update_mcq_score(w http.ResponseWriter, r *http.Request) {
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
		score = 10 // Increment the server-side variable by 10
		fmt.Println("User's answer is correct!", score)
	} else {
		score=0
		fmt.Println("User's answer is incorrect.", score)
	}

	// You can send a response back to the client here if needed.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"isCorrect": requestData.IsCorrect})

	// Update the user's score based on the response if needed.
	// Example: app.user.Update_score(subject, score)
}
