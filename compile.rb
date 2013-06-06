code = []
code << 'package main'
code << 'import "fmt"'
code << 'func main() {'
File.open("main.ket", 'r').each_line do |line|
  next if line[0] == "#"
  next if line[0] == "\n"
  code << 'user := make(map[string]string)' if line == "user\n"
  code << 'user["name"] = "slavik"' if line == "  name:slavik\n"
  code << 'fmt.Println(user["name"])' if line == "= user/name\n"
end
code << '}'

File.open("try.go", 'w') do |file|
  file.write(code.join("\n"))
end

`go build try.go`
