package main

import (
    "fmt"
    "net/http"
    "html/template"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/Suy56/GradeUpNow/internal/models"
)

func (app *application) theory_handler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    subject := vars["subject"]
    _, err := app.user.Get(models.G_CurrentUserSession)
    if err != nil {
        app.notFound(w)
        app.errorLog.Print(err)
        return
    }

    questions, err := app.user.Get_Theory(subject)
    if err != nil {
        app.serverError(w, err)
        return
    }

    // Get the current question index from the query parameter, or initialize it to 0
    currentQuestionIndex, err := strconv.Atoi(r.URL.Query().Get("q"))
    if err != nil {
        currentQuestionIndex = 0
    }

    // Check if the request is a POST request (user submitted an answer)
    if r.Method == http.MethodPost {
        // Parse the form data to get the user's answer
        err := r.ParseForm()
        if err != nil {
            app.serverError(w, err)
            return
        }

        // Get the user's answer from the form data
        userAnswer := r.Form.Get("answer")
        frmt_ans:=Format_ans(userAnswer)
        fmt.Println(userAnswer)

        fmt.Print(frmt_ans)
        // Process the user's answer as needed
        // For example, you can compare it to the correct answer and update the score.
    }

    // Check if the current question index is within bounds
    if currentQuestionIndex >= 0 && currentQuestionIndex < len(questions) {
        // Display the current question
        tmpl, err := template.ParseFiles("./ui/html/theory.html")
        if err != nil {
            app.serverError(w, err)
            return
        }

        // Check if there's a next question
        hasNextQuestion := currentQuestionIndex+1 < len(questions)

        // Pass the current question data and whether there's a next question to the template
        err = tmpl.Execute(w, struct {
            TQ_num           int
            TQ_question      string
            TQ_type          string
            HasNextQuestion  bool
            NextQuestionIndex int
        }{
            TQ_num:           questions[currentQuestionIndex].TQ_num,
            TQ_question:      questions[currentQuestionIndex].TQ_question,
            TQ_type:          questions[currentQuestionIndex].TQ_type,
            HasNextQuestion:  hasNextQuestion,
            NextQuestionIndex: currentQuestionIndex + 1,
        })
        if err != nil {
            app.serverError(w, err)
            return
        }
    } else {
        // Display a message when all questions have been answered
        fmt.Fprint(w, "All questions have been answered.")
    }
}
