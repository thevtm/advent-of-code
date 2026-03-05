# INPUTS

$input_path = "./input.txt"
$input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

# PROBLEM 1

timer_start = Time.now

$selected_batteries = $file_lines.map do |batteries|
  batteries_split = batteries.split("")

  left_num_str, left_index = batteries_split[...-1]
    .each_with_index
    .map {|x,i| [x, i]}
    .max_by {|x| [x[0], -x[1]]}

  right_num_str = batteries_split[(left_index + 1)..].max

  [left_num_str, right_num_str]
end

$joltages_sum = $selected_batteries.sum {|bs| bs[0].to_i * 10 + bs[1].to_i }

puts "Problem 1 Result: #{$joltages_sum} ● #{"%.1f" % (Time.now - timer_start)}s" # 17408 ● 0.0s

# PROBLEM 2

# timer_start = Time.now

# $selected_batteries = $file_lines.map do |batteries|
#   batteries
#     .split("")
#     .each_with_index
#     .sort_by {|x, i| [-x.to_i, i]}
#     .first(12)
#     .sort_by {|x, i| i}
#     .map {|x| x}
# end

# # $joltages_sum = $selected_batteries.sum do |bs|
# #   bs.each_with_index.map {|x, i| x.to_i * (10 ** (11 - i))}.sum
# # end

# puts "Problem 2 Result: #{$joltages_sum} ● #{"%.1f" % (Time.now - timer_start)}s" # ??
