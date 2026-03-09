# INPUTS

$input_path = './input.txt'
$connections_count = 1000

# $input_path = "./input-sample.txt"
# $connections_count = 10

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$positions = $file_lines.map { |line| line.split(',').map(&:to_i) }

# PROBLEM 1

timer_start = Time.now

$distances = []

$positions.each_with_index do |pos_a, i|
  $positions.drop(i + 1).each do |pos_b|
    distance = Math.sqrt((pos_a[0] - pos_b[0])**2 + (pos_a[1] - pos_b[1])**2 + (pos_a[2] - pos_b[2])**2)
    $distances.push([pos_a, pos_b, distance])
  end
end

$distances.sort_by! { |(_, _, d)| d }

$parents = {}

def find(pos)
  parent = $parents[pos]

  if parent.nil?
    $parents[pos] = pos
    return pos
  end

  return pos if pos == parent

  $parents[pos] = find(parent)
end

def union(pos_a, pos_b)
  par_a = find(pos_a)
  par_b = find(pos_b)

  $parents[pos_a] = $parents[pos_b] = $parents[par_b] = par_a
end

$distances[...$connections_count].each do |(pos_a, pos_b, _)|
  union(pos_a, pos_b)
end

$circuits_count = $parents.values.map { |p| find(p) }.tally

$answer = $circuits_count.map { |(_, v)| v }.sort.reverse.take(3).reduce(&:*)

puts "Problem 1 Result: #{$answer} ● #{format('%.1f', Time.now - timer_start)}s" # 115885 ● 0.4s

# PROBLEM 2

timer_start = Time.now

$parents = {}
$answer = nil

$distances.each do |(pos_a, pos_b, _)|
  union(pos_a, pos_b)

  next unless $parents.length == $positions.length

  first_parent = find($parents.keys[0])
  all_connected = $parents.values.all? { |p| find(p) == first_parent }

  next unless all_connected

  # puts "pos_a #{pos_a} pos_b #{pos_b}"
  $answer = pos_a[0] * pos_b[0]
  break
end

puts "Problem 2 Result: #{$answer} ● #{format('%.1f', Time.now - timer_start)}s" # 274150525 ● 0.0s
