# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$id_ranges = $file_lines[0].split(',').map do |range_str|
  left, right = range_str.split('-')
  [left.to_i, right.to_i]
end

# PROBLEM 1

timer_start = Time.now

$invalid_ids_count = 0

$id_ranges.each do |id_range|
  left, right = id_range

  (left..right).each do |n|
    decimals = Math.log10(n).floor + 1
    next if decimals % 2 === 1

    half_decimals = decimals / 2
    is_invalid = (0...half_decimals).zip(half_decimals...decimals)
                                    .all? { |a, b| (n / 10**a) % 10 === (n / 10**b) % 10 }

    $invalid_ids_count += n if is_invalid
  end
end

puts "Problem 1 Result: #{$invalid_ids_count} ● #{format('%.1f', Time.now - timer_start)}s" # 16793817782 ● 8.6s

# PROBLEM 2

timer_start = Time.now

$invalid_ids_count = 0

$id_ranges.each do |id_range|
  left, right = id_range

  (left..right).each do |n|
    n_str = n.to_s
    decimals = n_str.length
    half_decimals = decimals / 2

    (1..half_decimals).each do |repeat_len|
      next if decimals % repeat_len != 0

      repetitions = decimals / repeat_len

      is_invalid = (1...repetitions).all? do |repetition_index|
        offset = repetition_index * repeat_len
        (0...repeat_len).all? { |i| n_str[i] == n_str[i + offset] }
      end

      next unless is_invalid

      # puts "+ #{n}" if is_invalid
      $invalid_ids_count += n
      break
    end
  end
end

puts "Problem 1 Result: #{$invalid_ids_count} ● #{format('%.1f', Time.now - timer_start)}s" # 27469417404 ● 7.8s
