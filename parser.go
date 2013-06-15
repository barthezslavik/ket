package main
import (
  "os"
  "io/ioutil"
  "strings"
  "fmt"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func get_max(values []int) int {
  max := 0
  for _,v := range values {
    if v > max {
      max = v
    }
  }
  return max
}

func recurse_from_deep(lines []string, max int) {
  for _,line := range lines {
      symbols := []byte(line)
      deep := 0
      for _,symbol := range symbols {
        if(symbol == 32) {
          deep++
        }
      }
      println(line)
  }
}

func parse_object(a []string) []string {
  contents,_ := ioutil.ReadFile(os.Args[1]+".ket")
  lines := strings.Split(string(contents), "\n")
  deep_array := make([]int, 0)
  for _,line := range lines {
    if(strings.Contains(line, " ")) {
      symbols := []byte(line)
      deep := 0
      for _,symbol := range symbols {
        if(symbol == 32) {
          deep++
        }
      }
      //println(line)
      deep_array = append(deep_array, deep)
    }
  }

  max := get_max(deep_array)
  recurse_from_deep(lines, max)

  fmt.Println(max)

  a = append(a, `
  user    := map[string]interface{} { "first": first, "someone": someone }`)
  a = append(a, `
  someone := map[string]interface{} { "hello": hello }`)
  a = append(a, `
  first   := map[string]interface{} { "second": second }`)
  a = append(a, `
  second  := map[string]interface{} { "name": "slavik", "age": 21 }`)
  a = append(a, `
  hello   := map[string]interface{} { "super": "hero"}`)
  a = append(a, `
  println(user)`)
  return a
}

func main() {
  a := make([]string, 0)
  a = append(a, `
  package main`)
  a = append(a, `
  func main() {`)
  a = parse_object(a)
  a = append(a, `
}`)

//for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
//  a[i], a[j] = a[j], a[i]
//}

f, err := os.Create(os.Args[1]+".go")
check(err)
for i := 0; i < len(a); i++ {
  f.WriteString(a[i])
}

//j, _ := json.Marshal(user)
//fmt.Println(a)

}
