require 'logger'
require 'quartz'

$logger = Logger.new('logfile.log')

def execute(expression)
  startTime = Time.now
  result = calculate(expression)
  endTime = Time.now

  duration = endTime - startTime

  log(expression, result, duration)
end

def calculate(expression)
  client = Quartz::Client.new(file_path: 'worker.go')
  return client[:resolver].call('Calc', 'Expression'=>expression)['Result']
end

def log(expression, result, duration)
  $logger.info "Expression '#{expression}' resulted '#{result}' in #{'%.04f' % duration}s."
end

def log_err(err)
  $logger.error err
end

execute("2 3 *")
