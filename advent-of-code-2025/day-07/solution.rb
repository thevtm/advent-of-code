# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$rows = $file_lines.length
$cols = $file_lines[0].length

# PROBLEM 1

timer_start = Time.now

$grid = Marshal.load(Marshal.dump($file_lines))
$beam_split_count = 0

(1...$rows).each do |row_index|
  (0...$cols).each do |col_index|
    cell = $grid[row_index][col_index]
    cell_above = $grid[row_index - 1][col_index]

    next if cell_above != 'S' && cell_above != '|'
    next if cell == '|'

    if cell == '^'
      $beam_split_count += 1

      $grid[row_index][col_index - 1] = '|' if col_index.positive? && $grid[row_index][col_index - 1] == '.'

      $grid[row_index][col_index + 1] = '|' if col_index < $cols - 1 && $grid[row_index][col_index + 1] == '.'
    else
      $grid[row_index][col_index] = '|'
    end
  end
end

puts "Problem 1 Result: #{$beam_split_count} ● #{format('%.1f', Time.now - timer_start)}s" # 1675 ● 0.0s

# PROBLEM 2

timer_start = Time.now

$map = $file_lines
$S_col = $file_lines[0].index("S")

$queue = [[1, $S_col, 1]]

$timelines_count = 0

def upsert_beam(row, col, count)
  queue_element_index = $queue.index {|(r, c)| row == r && col == c}

  if queue_element_index.nil?
    $queue.push([row, col, count])
  else
    $queue[queue_element_index][2] += count
  end
end

until $queue.empty?
  row, col, count = $queue.shift

  next if col < 0 || col >= $cols

  if row >= $rows
    $timelines_count += count
    next
  end

  cell = $map[row][col]

  if cell == "."
    upsert_beam(row + 1, col, count)
  elsif cell == "^"
    upsert_beam(row + 1, col - 1, count)
    upsert_beam(row + 1, col + 1, count)
  end
end

puts "Problem 2 Result: #{$timelines_count} ● #{format('%.1f', Time.now - timer_start)}s" # 187987920774390 ● 0.0s
