package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
)

type SendEmailRequestBody struct {
	Sender string `json:"sender"`;
	Receivers []string `json:"receivers"`;
	Msg string `json:"msg"`;
}

func (app *application) handleSendMail(w http.ResponseWriter, r *http.Request) {



	body, err := io.ReadAll(r.Body);

	var validationErrors = make(map[string]string);


	if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError);
		  return;
	}


	var requestBody =  &SendEmailRequestBody{};
	
    err = json.Unmarshal(body, requestBody)

	if err != nil {
		  http.Error(w, err.Error(), http.StatusBadRequest);
		  return;
	}

	if requestBody.Sender == ""{
		validationErrors["sender"] = "sender required"
	}
	
	if len(requestBody.Receivers) == 0 {
		validationErrors["receivers"] = "receivers empty"
	}

	if requestBody.Msg == ""{
		validationErrors["msg"] = "message required"
	}



	if  len(validationErrors) != 0 {
		
		 data, err := json.Marshal(validationErrors)
		
		 if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError);
		    return;
		 }

		  w.Header().Set("Content-Type", "application/json")		
		_, err = w.Write(data);


		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError);
		    return;
		 }
		 return;
	}


	msg := []byte(requestBody.Msg);

	err = smtp.SendMail("smtp.gmail.com:587", *app.smtpAuth, requestBody.Sender, requestBody.Receivers, msg);


	if err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError);
		  return;
	}

	fmt.Fprintln(w, "Successful");

}