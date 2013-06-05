File.open("main.ket", 'r').each_line do |line|
  next if line[0] == "#"
  next if line[0] == "\n"
  puts line.inspect
end
