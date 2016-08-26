require 'logger'
require 'quartz'

$logger = Logger.new('logfile.log')
$client = Quartz::Client.new(file_path: 'worker.go')

def execute(expression)
  begin
    startTime = Time.now
    result = calculate(expression)
    endTime = Time.now

    duration = endTime - startTime

    log(expression, result, duration)

    return {"result" => result, "duration" => duration, "err" => nil}
  rescue Exception => err
    log_err(expression, err)
    return {"result" => nil, "duration" => nil, "err" => err.message}
  end
end

def calculate(expression)
  return $client[:resolver].call('Calc', 'Expression'=>expression)['Result']
end

def log(expression, result, duration)
  $logger.info "Expression '#{expression}' resulted '#{result}' in #{'%.04f' % duration}s."
end

def log_err(expression, err)
  $logger.error "An error ocurred evaluating the expression '#{expression}':  " + err.message
end

def start_slave()
  while cmd = STDIN.gets
    cmd.chop!
    if cmd == "exit"
      break
    else
      print execute(cmd),"\n"
      print "[end]\n"
      STDOUT.flush
    end
  end
end

start_slave()
