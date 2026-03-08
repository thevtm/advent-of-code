# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

# PROBLEM 1

timer_start = Time.now

$selected_batteries = $file_lines.map do |batteries|
  batteries_split = batteries.chars

  left_num_str, left_index = batteries_split[...-1]
                             .each_with_index
                             .map { |x, i| [x, i] }
                             .max_by do |x|
    [
      x[0], -x[1]
    ]
  end

  right_num_str = batteries_split[(left_index + 1)..].max

  [left_num_str, right_num_str]
end

$joltages_sum = $selected_batteries.sum { |bs| bs[0].to_i * 10 + bs[1].to_i }

puts "Problem 1 Result: #{$joltages_sum} ● #{format('%.1f', Time.now - timer_start)}s" # 17408 ● 0.0s

# PROBLEM 2

timer_start = Time.now

$selected_batteries = $file_lines.map do |batteries|
  left = 0

  12.downto(1).map do |n|
    right = batteries.length - n
    battery, index = batteries[left..right].chars.each_with_index.max_by { |b, _| b }
    left = left + index + 1
    battery
  end
end

$joltages = $selected_batteries.map do |bs|
  bs.each_with_index.map { |x, i| x[0].to_i * (10**(11 - i)) }.sum
end

$joltages_sum = $joltages.sum

puts "Problem 2 Result: #{$joltages_sum} ● #{format('%.1f', Time.now - timer_start)}s" # 172740584266849 ● 0.0s
