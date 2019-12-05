# Golang-Gin-example

This is a backend Restful API example based on Golang Gin web framework.



## Environment

Backend：Golang 1.13.4

Database：PostgreSQL 12



## Directory Structure

~~~
my-gin
├── database
│   └── postgresql.go
├── models
│   └── person.go
├── router
│   ├── router.go
│		└── apis
│   		└── v1
│						└── person.go
├── main.go
├── go.mod
└── go.sum
~~~



## Get Started

1. #### Go modules：a tool for managing Go modules

   ~~~go
   // open Go modules
   go env -w GO111MODULE=on
   
   // init Go modules, it will generate go.mod file
   go mod init [MODULE_PATH]
   ~~~

   ##### basic instructions of the Go & Go Modules

   * `go get`：pull new dependencies
   * `go get -u`：update existing dependencies
   * `go mod tidy`：organize existing dependencies

2. #### Initial Database and Create Table

   ~~~sql
   CREATE TABLE person (
     id serial PRIMARY KEY NOT NULL,
     first_name varchar DEFAULT '',
     last_name varchar DEFAULT ''
   );
   ~~~

3. #### Replace your module path in go.mod

   replace it to your absolute path

   ~~~go
   require (
       ...
   )
   
   replace (
       my-gin/database => ~/my-gin/database
       my-gin/models => ~/my-gin/models
       my-gin/router => ~/my-gin/router
   )
   ~~~



## License

2019, Ching-Hsuan Su

MIT