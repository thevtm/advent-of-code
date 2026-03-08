# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$blank_line_index = $file_lines.index { |l| l == '' }

$fresh_ranges = $file_lines[...$blank_line_index].map do |fresh_range_str|
  left, right = fresh_range_str.split('-')
  [left.to_i, right.to_i]
end

$ingredient_ids = $file_lines[($blank_line_index + 1)..].map(&:to_i)

# PROBLEM 1

timer_start = Time.now

$fresh_ingredients = $ingredient_ids.count do |ingredient_id|
  $fresh_ranges.any? { |r| ingredient_id.between?(r[0], r[1]) }
end

puts "Problem 1 Result: #{$fresh_ingredients} ● #{format('%.1f', Time.now - timer_start)}s" # 733 ● 0.0s

# PROBLEM 2

timer_start = Time.now

$fresh_ranges_sorted = $fresh_ranges.sort

$collapsed_ranges = [$fresh_ranges_sorted[0]]

$fresh_ranges_sorted.drop(1).each do |range|
  left, right = range
  _, prev_right = $collapsed_ranges[-1]

  if left > prev_right
    $collapsed_ranges.push(range)
    next
  end

  $collapsed_ranges[-1][1] = [prev_right, right].max
end

$fresh_ids_count = $collapsed_ranges
                   .map { |l, r| r - l + 1 }
                   .sum

puts "Problem 2 Result: #{$fresh_ids_count} ● #{format('%.1f', Time.now - timer_start)}s" # 345821388687084 ● 0.0s
