
  package main
  func main() {
  user    := map[string]interface{} { "first": first, "someone": someone }
  someone := map[string]interface{} { "hello": hello }
  first   := map[string]interface{} { "second": second }
  second  := map[string]interface{} { "name": "slavik", "age": 21 }
  hello   := map[string]interface{} { "super": "hero"}
  println(user)
}