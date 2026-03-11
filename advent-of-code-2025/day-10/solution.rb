# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$inputs = $file_lines.map do |line|
  _, lights, buttons, _ = line.match(/\[(.+)\] (.+) \{(.+)\}/).to_a

  lights = lights.chars.map {|c| c == "#"}
  buttons = buttons.gsub(/[\(\)]/, "").split(" ").map {|bs| bs.split(",").map {|b| b.to_i}}

  [lights, buttons]
end

# PROBLEM 1

timer_start = Time.now

$total_button_presses = 0

$inputs.each_with_index do |(target_lights, buttons), index|
  initial_lights = Array.new(target_lights.length, false)
  queue = [[initial_lights, 0]]
  min = Float::INFINITY

  until queue.empty?
    lights, button_presses = queue.shift

    # The first one we found should be the best one since we doing BFS
    next if button_presses >= min

    if (lights == target_lights)
      min = [min, button_presses].min
      break
    end

    buttons.each do |button|
      new_lights = lights.dup
      button.each {|l| new_lights[l] = !new_lights[l]}
      queue.push([new_lights, button_presses + 1])
    end
  end

  puts "index #{index} min #{min} target_lights #{target_lights} buttons #{buttons}"
  $total_button_presses += min
end

puts ""
puts "Problem 1 Result: #{$total_button_presses} ● #{format('%.1f', Time.now - timer_start)}s" # 484 ● 14.3s

# PROBLEM 2

# timer_start = Time.now


# puts "Problem 2 Result: #{0} ● #{format('%.1f', Time.now - timer_start)}s" # ?? ● 0.0s
