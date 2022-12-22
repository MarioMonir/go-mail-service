# go-mail-service

#### Project aims to have a simple native golang service ( api server )

---

### 1. Technologies :

- just golang ( GO )

### 2. Installation :

- no need for any installation if you already have golang

### 3. Run :

Api server will run on port 8080

    go run main.go

### 4. API Documentation

just 2 edpoints on  the same url "/" : http://localhost:8080/

---

- Index Home page : [ GET ]

  - serves a welcome html message :

    ```
    <center>
    <h1>Welcome to Golang Mailing Service using Gmail !</h1>
    </center>
    ```

---

- Send mail : [ POST ]
  - req body:
    ```
    {
        "to":"<email@example.com>",
        "subject":"email_Subject",
        "body":"email_body"
    }
    ```
  - serve a json response body
    ```
    {
        "success": "true",
        "url": "/",
        "method": "POST",
        "status": "202",
        "message": "success to send mail"
    }
    ```

---

###### @author MarioMonir
