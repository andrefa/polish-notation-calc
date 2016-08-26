require 'quartz'

client = Quartz::Client.new(file_path: 'worker.go')

puts client[:resolver].call('Calc', 'Expression'=>'1 2 +')['Result']
