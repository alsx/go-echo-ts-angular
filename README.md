
Task

Summary
1. Backend application should provide RESTful API to register new user and authenticate users in the service
2. Backend application should provide ability to register new user and authenticate users with Facebook account OR email
3. Backend should to use MySQL to store user authentication data and sessions
4. Backend should to be fully stateless and be able scale horizontal (aka micro-services architecture)
5. API should be versioned

Optional task
1. HTML/JS Frontend with integration with backend and ability to register new user and authenticate users

Deliverables
1. Backend and Frontend (optional) source code and binary files
2. Deployment package

Acceptance Critiries
1. Go code should pass gometalinter default check (https://github.com/alecthomas/gometalinter)
2. All connection variables for MySQL connection should available to change from command line with Go binary start


# Results

### build api server
```
curl https://glide.sh/get | sh
glide install
go build github.com/alsx/enli-task/src/api
```
### create database
```sh
echo "CREATE DATABASE task" | mysql
mysql task < src/api/schema.sql
```
### run api server
```sh
./api -dsn 'user:pwd@tcp(127.0.0.1:3306)/task?parseTime=true'
```


### List of Versions
```sh
curl http://localhost:1323/api
```
```json
{
    "Links": [
        "v1/"
    ]
}
```
### List of Endpoints
```sh
curl http://localhost:1323/api/v1
```
```json
{
    "Links": [
        "signin/",
        "login/",
        "user/",
        "fb-login/",
        "fb-callback/"
    ]
}
```
### Register User
```sh
curl -X POST -d 'name=John Doe' -d 'email=john@doe.com' -d 'password=Pa$$w0rd' http://localhost:1323/api/v1/signin
```
```json
{
    "ID": 12,
    "Name": "John Doe",
    "Email": "john@doe.com"
}
```
### Login
```sh
curl -X POST -d 'email=john@doe.com' -d 'password=Pa$$w0rd' http://localhost:1323/api/v1/login
```
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImpvaG5AZG9lLmNvbSIsImV4cCI6MTUwMTc5OTQ3MH0.nlPd4R3vWb-L5jYCrpvvOr5CZfmDgx5-202Wejx04NU"
}
```

### Call closed url with access token
```sh
curlhttp://127.0.0.1:1323/api/v1/user/ -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImpvaG5AZG9lLmNvbSIsImV4cCI6MTUwMTc5OTYxOH0.R93IlZDw-wChdE5ITTdkkCo24-rNI9Q0NjomFz8S8cY"
```
```json
{
    "ID": 12,
    "Name": "John Doe",
    "Email": "john@doe.com"
}
```
### TBD: fb auth
