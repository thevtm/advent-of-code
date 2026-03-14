# INPUTS

$input_path = './input.txt'
# $input_path = "./input-sample.txt"

$file_content = File.read(File.join(__dir__, $input_path))
$file_lines = $file_content.lines(chomp: true)

$inputs = $file_lines.map do |line|
  _, lights, buttons, joltages = line.match(/\[(.+)\] (.+) \{(.+)\}/).to_a

  lights = lights.chars.map { |c| c == '#' }
  buttons = buttons.gsub(/[()]/, '').split.map { |bs| bs.split(',').map(&:to_i) }
  joltages = joltages.split(',').map(&:to_i)

  [lights, buttons, joltages]
end

# PROBLEM 1

timer_start = Time.now

$total_button_presses = 0

$inputs.each_with_index do |(target_lights, buttons), _index|
  initial_lights = Array.new(target_lights.length, false)
  queue = [[initial_lights, 0, 0]]
  min_button_presses = Float::INFINITY

  until queue.empty?
    lights, button_presses, prev_button_index = queue.shift

    if lights == target_lights
      min_button_presses = button_presses
      break
    end

    buttons.each_with_index.drop(prev_button_index).each do |button, button_index|
      new_lights = lights.dup
      button.each { |l| new_lights[l] = !new_lights[l] }
      queue.push([new_lights, button_presses + 1, button_index])
    end
  end

  $total_button_presses += min_button_presses
end

puts ''
puts "Problem 1 Result: #{$total_button_presses} ● #{format('%.1f', Time.now - timer_start)}s" # 484 ● 0.3s

# PROBLEM 2

puts ''
timer_start = Time.now

require 'z3'

$total_button_presses = 0

$inputs.each_with_index do |(_, buttons, target_lights), _index|
  solver = Z3::Optimize.new

  button_ints = []
  buttons_for_lights = Array.new(target_lights.length) { [] }

  buttons.each_with_index do |button, _index|
    button_int = Z3.Int("button(#{button.join(',')})")
    button_ints.push(button_int)

    solver.assert(button_int >= 0)

    button.each { |l| buttons_for_lights[l].push(button_int) }
  end

  buttons_for_lights.zip(target_lights).each do |bfl, tl|
    solver.assert(bfl.reduce { |acc, n| acc + n } == tl)
  end

  total_button_presses_int = Z3.Int('total_button_presses')
  solver.assert(button_ints.reduce(&:+) == total_button_presses_int)
  solver.minimize(total_button_presses_int)

  throw 'Unsatisfiable' unless solver.satisfiable?

  $total_button_presses += solver.model[total_button_presses_int].to_i
end

puts "Problem 2 Result: #{$total_button_presses} ● #{format('%.1f', Time.now - timer_start)}s" # 19210 ● 0.3s
