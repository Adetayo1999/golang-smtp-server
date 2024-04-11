package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)


type application struct {
	 smtpAuth *smtp.Auth
	 errorLogger *log.Logger
	 mux *http.ServeMux
}

func main() {

	errorLogger := log.New(os.Stderr, "Error:\t", log.Ldate|log.Ltime|log.Llongfile);

	err := godotenv.Load(".env");

	if err != nil {
		 errorLogger.Fatalln(err);
	}

	port := os.Getenv("PORT");
	email := os.Getenv("EMAIL");
	app_password := os.Getenv("APP_PASSWORD");

	smtpAuth := smtp.PlainAuth("", email, app_password, "smtp.gmail.com");

	mux := http.NewServeMux();

	app := application{
		smtpAuth: &smtpAuth,
		errorLogger: errorLogger,
		mux: mux,
	}

	srv := &http.Server{
	   	Addr: fmt.Sprintf(":%v", port),
		Handler: app.routes(),
		ErrorLog: errorLogger,
	}

	log.Println("Server starting ");

	if err := srv.ListenAndServe(); err != nil {
		 errorLogger.Fatalln(err);
	}

}