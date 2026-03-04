# INPUTS

$input_path = "./input.txt"
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$id_ranges = $file_lines[0].split(",").map do |range_str|
  left, right = range_str.split("-")
  [left.to_i, right.to_i]
end

# PROBLEM 1

$invalid_ids_count = 0

$id_ranges.each do |id_range|
  left, right = id_range

  (left..right).each do |n|
    decimals = Math.log10(n).floor + 1
    next if decimals % 2 === 1

    half_decimals = decimals / 2
    is_invalid = (0...half_decimals).zip(half_decimals...decimals)
      .all? {|a,b| (n / 10 ** a) % 10 === (n / 10 ** b) % 10}

    $invalid_ids_count += n if is_invalid
  end
end

puts "Problem 1 Result: #{$invalid_ids_count}" # 16793817782

# PROBLEM 2

# puts "Problem 2 Result: #{$dial_pass_by_zero_count}" # 6623
