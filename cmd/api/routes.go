package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

    app.mux.HandleFunc("/", app.handleSendMail);
	
	return app.mux;
}