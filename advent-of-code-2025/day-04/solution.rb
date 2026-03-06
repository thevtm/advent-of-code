# INPUTS

$input_path = "./input.txt"
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$rows = $file_lines.length
$cols = $file_lines[0].length

$grid = $file_lines

# PROBLEM 1

timer_start = Time.now

$accessible_rolls = 0

def has_roll(x, y)
  return false if x < 0 || x >= $rows
  return false if y < 0 || y >= $rows
  $grid[x][y] == "@"
end

def count_adjacent_rolls(row, col)
  adjacent_count = 0

  adjacent_count += 1 if has_roll(row - 1, col - 1)
  adjacent_count += 1 if has_roll(row - 1, col)
  adjacent_count += 1 if has_roll(row - 1, col + 1)

  adjacent_count += 1 if has_roll(row, col - 1)
  adjacent_count += 1 if has_roll(row, col)
  adjacent_count += 1 if has_roll(row, col + 1)

  adjacent_count += 1 if has_roll(row + 1, col - 1)
  adjacent_count += 1 if has_roll(row + 1, col)
  adjacent_count += 1 if has_roll(row + 1, col + 1)

  adjacent_count
end

(0...$rows).each do |row|
  (0...$cols).each do |col|
    next if $grid[row][col] == "."
    $accessible_rolls += 1 if (count_adjacent_rolls(row, col) <= 4)
  end
end

puts "Problem 1 Result: #{$accessible_rolls} ● #{"%.1f" % (Time.now - timer_start)}s" # 1351 ● 0.1s

# PROBLEM 2

timer_start = Time.now

$removed_rolls_count = 0

loop do
  removed_a_roll = false

  (0...$rows).each do |row|
    (0...$cols).each do |col|
      next if $grid[row][col] == "."

      if (count_adjacent_rolls(row, col) <= 4)
        $grid[row][col] = "."
        $removed_rolls_count += 1
        removed_a_roll = true
      end
    end
  end

  break unless removed_a_roll
end

puts "Problem 2 Result: #{$removed_rolls_count} ● #{"%.1f" % (Time.now - timer_start)}s" # 8345 ● 0.7s
