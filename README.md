## Introduction 

This repository contains a small web application, which fulfill the requirements for [array backengineering task](https://gitlab.com/array.com/tests-backend/-/blob/master/test-1.md). 

It provides three API endpoints:
  * to register user with email, password
  * to login user using registered email / password 
  * to logout user session 

The project is tested with `golang==1.6.x`, `mysql==8.0.x` and `redis==5.0.x`. 
## Setup 

The easiest way to run this application is with docker-compose. 

```sh
$ docker-compose build
$ docker-compose up -d
```

However, you can easily set `DATABASE_URL` and `REDIS_ADDR` environment variables to point to relevant values and start the app without using docker as well.

```sh
$ export DATABASE_URL=test:test@(localhost:3306)/test
$ export REDIS_ADDR=localhost:6379
$ export PORT=":8080"
$ go run app.go 
```

### Testing 

To run the package tests, after setting `DATABASE_URL` and `REDIS_ADDR` environment variables, use `go test`.

```sh
$ go test ./dbm
$ go test ./api
```

You can also use `httpie` tool to quickly test the API endpoints from command line. 

- To register a user account

```sh
$ trhura @ eightcat > http post http://localhost:8080/users/register name="a" email="a@gmail.com" password="a"

HTTP/1.1 200 OK
Content-Length: 54
Content-Type: application/json
Date: Tue, 23 Nov 2021 14:48:38 GMT
Set-Cookie: ssn=36952c0ce79499e4; Path=/; Max-Age=86400; HttpOnly

{
    "message": "user created successfully",
    "success": true
}
```

- To login with registed user

```sh
$ http --session=./cookie.json -a "a@gmail.com:a" post http://localhost:8080/users/login

HTTP/1.1 200 OK
Content-Length: 50
Content-Type: application/json
Date: Tue, 23 Nov 2021 14:49:50 GMT
Set-Cookie: ssn=3787f5146f6cc4df; Path=/; Max-Age=86400; HttpOnly

{
    "message": "user login successful",
    "success": true
}
```

- To logout an existing session

```sh
$ http --session=./cookie.json post http://localhost:8080/users/logout

HTTP/1.1 200 OK
Content-Length: 51
Content-Type: application/json
Date: Tue, 23 Nov 2021 14:50:55 GMT
Set-Cookie: ssn=; Path=/; Expires=Tue, 23 Nov 2021 14:50:55 GMT; Max-Age=0; HttpOnly

{
    "message": "user logout successful",
    "success": true
} 
```