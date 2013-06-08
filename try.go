package main
import "fmt"
func main() {
user := make(map[string]string)
user["name"] = "slavik"
fmt.Println(user["name"])
}