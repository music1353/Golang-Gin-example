package main

import (
    db "my-gin/database"
    "my-gin/router"
)

func main() {
    defer db.SqlDB.Close()

    router := router.InitRouter()
    router.Run()
}