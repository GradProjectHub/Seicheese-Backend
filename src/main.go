package main

import (
    "fmt"
    "seicheese/database"
)

func main() {
    _, err := database.InitDB()
    if err != nil {
        fmt.Printf("データベースに接続できませんでした: %v\n", err)
        return
    }
    
    fmt.Println("データベースへの接続が成功しました")
}