package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
)

// -------------------------------------------------------

const (
	port     = ":8080"
	host     = "smtp.gmail.com"
	mailPort = ":587"
	user     = "techhivedevs@gmail.com"
	password = "jigteoorzbvvpktc"
	apiKey   = "mario_secret"
)

// -------------------------------------------------------

// Mail Hanndler Struct
type Mail struct {
	To      string
	Subject string
	Body    string
}

// Error Hanlder Struct
type Error struct {
	Name    string
	Code    string
	Message string
}

// -------------------------------------------------------

// Send Mail
// the function sends a mail by gmail to certain email
// passed as to with subject and body through a struct
// of type Mail{To, Subject, Body}
func sendMail(mail Mail) error {
	if mail.To == "" || mail.Subject == "" || mail.Body == "" {
		return errors.New(" Mail must have a valid  to, subject, body")
	}
	return smtp.SendMail(
		host+mailPort,
		smtp.PlainAuth("", user, password, host),
		user,
		[]string{mail.To},
		[]byte(""+
			"To:"+mail.To+"\r\n"+
			"Subject:"+mail.Subject+"\r\n\r\n"+
			mail.Body+"\r\n"),
	)
}

// -------------------------------------------------------

// Mail Handler
// the handler aims to act as controller for path POST "/"
// with a body bind to type Mail and use the sendMail method
// to send the mail then respons w json response
func mailHandler(w http.ResponseWriter, r *http.Request) {
	var mail Mail
	// Decode Request Body to Struct Mail
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		errorHandler(w, r, Error{Name: "DecodeMailFailure", Code: "400", Message: err.Error()})
		return
	}
	// Sending Mail
	err = sendMail(mail)
	if err != nil {
		errorHandler(w, r, Error{Name: "SendMailFaillure", Code: "500", Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{` +
		`"success":"true",` +
		`"url":"` + r.URL.Path + `",` +
		`"method":"` + r.Method + `",` +
		`"status":"202",` +
		`"message": "success to send mail"` +
		`}`))
}

// -------------------------------------------------------

// Index Handler
// acts as a controller for path GET "/" for the app,
// server a html welcome message to the client
func indexHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(
		[]byte("" +
			"<center>" +
			"<h1>Welcome to Golang Mailing Service using Gmail !</h1>" +
			"</center>"))
}

// -------------------------------------------------------

// Error Handler
// Send A Error in json for to the clinet, errors contains
// success boolean false ,url of request , method of request ,
// error of type Error{Name, Code, Message}
func errorHandler(w http.ResponseWriter, r *http.Request, e Error) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{` +
		`"success": "false",` +
		`"url":"` + r.URL.Path + `",` +
		`"method":"` + r.Method + `",` +
		`"status":"` + e.Code + `",` +
		`"errorName":"` + e.Name + `",` +
		`"message": "` + e.Message + `"` +
		`}`))
}

// -------------------------------------------------------

// Router
// function aims to return a http.HandlerFunc by routing
// between if path is GET "/" server indexHandler
// and if path is POST "/" server mailHandler
// else will serve the errorHandler 404 NotFound
func router() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" && r.Method == "GET" {
			indexHanlder(w, r)
			return
		}
		if r.URL.Path == "/" && r.Method == "POST" {
			mailHandler(w, r)
			return
		}

		errorHandler(w, r,
			Error{Name: "NotFound", Code: "404",
				Message: "url path method is not found"})
	}
}

// -------------------------------------------------------

// Main
// Entry Point for app and the server launcher
func main() {
	fmt.Println("Server is Listintg on port", port)
	http.HandleFunc("/", router())
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
