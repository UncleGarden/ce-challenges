File.open(ARGV[0]).each_line do |line|
  s = line.gsub!( /[^\d]/, "").split("").map {|s| s.to_i }
  puts s.inject {|sum,x| sum + x ** s.length } == line.to_i ? "True" : "False"
end
