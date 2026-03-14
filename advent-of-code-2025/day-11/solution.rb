# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$graph = $file_lines.each_with_object({}) do |line, acc|
  source = line[...3]
  targets = line[4..].split

  acc[source] = targets
end

# PROBLEM 1

timer_start = Time.now

$cache = {}
$cache['out'] = 1

def path_count_from(server_rack)
  $cache[server_rack] ||= $graph[server_rack].map { |x| path_count_from(x) }.sum
end

$paths_count = path_count_from('you')

puts "Problem 1 Result: #{$paths_count} ● #{format('%.1f', Time.now - timer_start)}s" # 688 ● 0.0s

# PROBLEM 2

# timer_start = Time.now

# puts "Problem 2 Result: #{0} ● #{format('%.1f', Time.now - timer_start)}s" # 19210 ● 0.3s
