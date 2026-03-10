# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$positions = $file_lines.map { |line| line.split(',').map(&:to_i) }

# PROBLEM 1

timer_start = Time.now

$largest_area = 0

$positions.each_with_index do |(x_a, y_a), i|
  $positions[(i + 1)..].each do |(x_b, y_b)|
    area = (x_a - x_b + 1).abs * (y_a - y_b + 1).abs
    $largest_area = area if area > $largest_area
  end
end

puts "Problem 1 Result: #{$largest_area} ● #{format('%.1f', Time.now - timer_start)}s" # 4749672288 ● 0.0s

# PROBLEM 2

timer_start = Time.now

$EMPTY = 0
$INSIDE = 1
$OUTSIDE = 2

$compressed_x = $positions.map { |p| p[0] }.sort.uniq.each_with_index.to_h { |x, i| [x, i + 1] }
$compressed_y = $positions.map { |p| p[1] }.sort.uniq.each_with_index.to_h { |y, i| [y, i + 1] }

$max_x = $compressed_x.length
$max_y = $compressed_y.length

$grid = Array.new($max_y + 2) { Array.new($max_x + 2, $EMPTY) }

$positions.each_with_index do |pos_a, i|
  x_a = $compressed_x[pos_a[0]]
  y_a = $compressed_y[pos_a[1]]

  pos_b = $positions[(i + 1) % $positions.length]
  x_b = $compressed_x[pos_b[0]]
  y_b = $compressed_y[pos_b[1]]

  if x_a === x_b
    y_min = [y_a, y_b].min
    y_max = [y_a, y_b].max

    (y_min..y_max).each { |y| $grid[y][x_a] = $INSIDE }
  else
    x_min = [x_a, x_b].min
    x_max = [x_a, x_b].max

    (x_min..x_max).each { |x| $grid[y_a][x] = $INSIDE }
  end
end

# $grid.each {|x| puts x.map {|x| {$EMPTY => ".", $INSIDE => "#"}[x]}.join(" ")}
# puts ""

$stack = [[0, 0]]
$seen = Set.new

until $stack.empty?
  x, y = pos = $stack.pop

  next if x.negative? || x > $max_x + 1
  next if y.negative? || y > $max_y + 1
  next if $grid[y][x] === $INSIDE
  next if $seen.include?(pos)

  $seen.add(pos)
  $grid[y][x] = $OUTSIDE

  $stack.push([x - 1, y - 1])
  $stack.push([x - 1, y])
  $stack.push([x - 1, y + 1])

  $stack.push([x, y - 1])
  $stack.push([x, y])
  $stack.push([x, y + 1])

  $stack.push([x + 1, y - 1])
  $stack.push([x + 1, y])
  $stack.push([x + 1, y + 1])
end

# $grid.each {|x| puts x.map {|x| {$EMPTY => ".", $INSIDE => "#", $OUTSIDE => "X"}[x]}.join(" ")}
# puts ""

$largest_area = 0

$positions.each_with_index do |(x_a, y_a), i|
  $positions[(i + 1)..].each do |(x_b, y_b)|
    area = (x_a - x_b + 1).abs * (y_a - y_b + 1).abs

    next unless area > $largest_area

    x_min = [$compressed_x[x_a], $compressed_x[x_b]].min
    x_max = [$compressed_x[x_a], $compressed_x[x_b]].max

    y_min = [$compressed_y[y_a], $compressed_y[y_b]].min
    y_max = [$compressed_y[y_a], $compressed_y[y_b]].max

    is_inside = (x_min..x_max).all? { |x| (y_min..y_max).all? { |y| $grid[y][x] != $OUTSIDE } }

    $largest_area = area if is_inside
  end
end

puts "Problem 2 Result: #{$largest_area} ● #{format('%.1f', Time.now - timer_start)}s" # 1479665889 ● 3.4s
