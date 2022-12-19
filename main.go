package main

// -------------------------------------------------------

import (
	"fmt"
	"net/http"
	"net/smtp"
)

// -------------------------------------------------------

const (
	port     = ":8080"
	user     = "techhivedevs@gmail.com"
	password = "jigteoorzbvvpktc"
	apiKey   = "mario_secret"
)

// -------------------------------------------------------

func sendMail(to string, subject string, body string) error {
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", user, password, "smtp.gmail.com"),
		user,
		[]string{to},
		[]byte("To:"+to+"\r\n"+
			"Subject:"+subject+"\r\n"+
			"\r\n"+
			body+"\r\n"),
	)
	if err != nil {
		return err
	}
	return nil
}

// -------------------------------------------------------

func mailHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendMail("mario3monir@gmail.com", "subject mario", "hello mario")
		w.Write([]byte("Hello World to Golang Mailing Service by Gmail :)"))
	})
}

// -------------------------------------------------------

func main() {
	fmt.Println("Server is Listintg on port", port)
	if err := http.ListenAndServe(port, mailHandler()); err != nil {
		panic(err)
	}
}
