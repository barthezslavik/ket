package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, _ := sql.Open("mysql", "root:@/go_learn")
  defer db.Close()

  var name string
  db.QueryRow("SELECT name FROM users WHERE id=?", 1).Scan(&name)
  fmt.Printf(name)
}
