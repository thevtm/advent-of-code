# INPUTS

$input_path = "./input.txt"
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$number_lines = $file_lines[...-1].map {|line| line.split(" ").map {|x| x.to_i}}
$operators = $file_lines[-1].split(" ")

# PROBLEM 1

timer_start = Time.now

$operations_map = {"+" => :+, "*" => :*}

$answers = (0...$operators.length).map do |i|
  operator_fn = $operations_map[$operators[i]]
  $number_lines.map {|l| l[i]}.reduce(operator_fn)
end

$grand_total = $answers.sum

puts "Problem 1 Result: #{$grand_total} ● #{"%.1f" % (Time.now - timer_start)}s" # 6757749566978 ● 0.0s

# PROBLEM 2

timer_start = Time.now

$longest_number_line_length = $file_lines[...-1].map {|l| l.length}.max

$char_parse_map = {nil => nil, " " => nil, "0" => 0, "1" => 1, "2" => 2,
  "3" => 3, "4" => 4, "5" => 5, "6" => 6, "7" => 7, "8" => 8, "9" => 9}

$numbers = [[]]

(0...$longest_number_line_length).each do |i|
  digits = $file_lines[...-1].map {|l| $char_parse_map[l[i]]}

  if (digits.all? {|x| x.nil?})
    $numbers.push([])
    next
  end

  num = digits.reject {|x| x.nil?}.reduce(0) {|acc, x| acc * 10 + x}
  $numbers[-1].push(num)
end

$answers = $numbers.zip($operators).map do |ns, operator|
  operator_fn = $operations_map[operator]
  ns.reduce(operator_fn)
end

$grand_total = $answers.sum

puts "Problem 2 Result: #{$grand_total} ● #{"%.1f" % (Time.now - timer_start)}s" # 10603075273949 ● 0.0s
