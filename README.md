


Overview
This project implements a User Authentication and Authorization API using Golang, Echo, and MySQL. It supports:

    1. User Registration (Sign Up)
    2. User Authentication (Sign In) with JWT Tokens
    3. Authorization Middleware to protect API routes
    4. Token Revocation (Logout)
    5. Token Refreshing

Installation & Setup
1. Install Dependencies
Ensure you have Go, MySQL, and Postman/cURL installed.

Install Go: Download here
Install MySQL: MySQL Download

2. Clone the Repository
    https://github.com/AbhinashKumarSingh/authentication_in_go.git

3. install dependency 
   go mod tidy

4. Set Up MySQL Database
        app->docs->schema run mysql queries
        InitDB -> DB_USER=root
                  DB_PASSWORD=yourpassword
                  DB_NAME=users_db
5. 



//signup
curl --location 'http://localhost:8000/user-service/user/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"ahhh@gmail.com",
    "password":"1234"
}'

//signin
curl --location 'http://localhost:8000/user-service/user/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"ahhh@gmail.com",
    "password":"1234"
}'


//refresh-tokn
curl --location 'http://localhost:8000/user-service/user/refresh-token' \
--header 'Content-Type: application/json' \
--data '{
    "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA5MDYwNjR9.uWUwMP-ULSQtUVyXTabajmvSbSQpO8FLzfkZDeSwxZ4"
}'

//revole token
curl --location 'http://localhost:8000/user-service/user/revoke-token' \
--header 'Content-Type: application/json' \
--data '{
    "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA5MDYwNjR9.uWUwMP-ULSQtUVyXTabajmvSbSQpO8FLzfkZDeSwxZ4"
}'