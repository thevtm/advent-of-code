# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample-part-1.txt"

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

def path_count(server_rack)
  $cache[server_rack] ||= $graph[server_rack].map { |x| path_count(x) }.sum
end

$paths_count = path_count('you')

puts "Problem 1 Result: #{$paths_count} ● #{format('%.1f', Time.now - timer_start)}s" # 688 ● 0.0s

# PROBLEM 2

# $input_path = "./input-sample-part-2.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$graph = $file_lines.each_with_object({}) do |line, acc|
  source = line[...3]
  targets = line[4..].split

  acc[source] = targets
end

timer_start = Time.now

$cache = {}

def path_count(from, to, seen = Set.new)
  return 1 if from == to
  return 0 if from == 'out'

  $cache["#{from}-#{to}"] ||= $graph[from].reject { |x| seen.include?(x) }.map { |x| path_count(x, to, seen) }.sum
end

$svr_to_fft = path_count('svr', 'fft')
$fft_to_dac = path_count('fft', 'dac')
$dac_to_out = path_count('dac', 'out')

$paths_count = $svr_to_fft * $fft_to_dac * $dac_to_out

puts "Problem 2 Result: #{$paths_count} ● #{format('%.1f', Time.now - timer_start)}s" # 293263494406608 ● 0.0s
