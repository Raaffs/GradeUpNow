package main
import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

//models.leader_board return a slice of user struct and we need to loop over that slice to show all the users
//however  the indexing starts from 0, value of 1 needs to be added in each iteration. As there is no direct way
//to add values in html, this function is needed.  
func add(a,b int)int{
	return a+b
}