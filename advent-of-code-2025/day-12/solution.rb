# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$shapes = (0..5).map do |i|
  line_start = i * 5 + 1
  line_end = line_start + 3

  $file_lines[line_start...line_end]
end

$regions = $file_lines[30..].map do |line|
  _, width_str, height_str, shapes_freq_str = line.match(/(\d+)x(\d+): (.+)/).to_a

  width = width_str.to_i
  height = height_str.to_i
  shapes_freq = shapes_freq_str.split(" ").map(&:to_i)

  [width, height, shapes_freq]
end

# PROBLEM 1

timer_start = Time.now

$shapes_filled_space = $shapes.map {|s| s.join.count("#")}

$valid_regions = $regions.count do |(width, height, shapes_freq)|
  area = width * height
  area_required = shapes_freq.each_with_index.sum {|f, i| f * $shapes_filled_space[i]}

  puts "width #{width} height #{height} area #{area} area_required #{area_required} shapes_freq #{shapes_freq}"

  area_required <= area
end

puts "Problem 1 Result: #{$valid_regions} ● #{format('%.1f', Time.now - timer_start)}s" # 528 ● 0.0s

# PROBLEM 2

# timer_start = Time.now

# It doesn't exist!!

# puts "Problem 2 Result: #{0} ● #{format('%.1f', Time.now - timer_start)}s" # ?? ● 0.0s
