require 'open3'

def main
  puts "enter payload here"
  input = gets.chomp
  stdin, stdout, stderr, wait_thr = Open3.popen3("./test.bat #{input}")
  stdout.each_line { |line| puts line }
  stderr.each_line { |line| puts line }
  exit_status = wait_thr.value
  unless exit_status.success?
    puts "Command execution failed with exit code #{exit_status.exitstatus}"
  end
end

if __FILE__ == $PROGRAM_NAME
  main
end
