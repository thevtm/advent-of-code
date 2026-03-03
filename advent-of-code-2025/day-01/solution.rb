# INPUTS

$input_path = "./input.txt"
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

# PROBLEM 1

$dial = 50

$rotations = $file_lines.map do |n|
  n[1..].to_i * (n[0] === "L" ? -1 : 1)
end

$dial_positions = $rotations.map do |r|
  $dial = (100 + $dial + r) % 100
  $dial
end

$dial_zero_count = $dial_positions.count(0)

puts "Problem 1 Result: #{$dial_zero_count}"

# PROBLEM 2

# puts "Problem 2 Result:"
