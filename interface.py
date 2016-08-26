import json
from subprocess import Popen, PIPE, STDOUT

print 'Launching ruby process...'
slave = Popen(['ruby', 'logger.rb'], stdin=PIPE, stdout=PIPE, stderr=STDOUT)

def read_expressions():
    total = int(raw_input('Enter number of expressions:'))

    expressions = [''] * total

    for i in xrange(total):
        line = raw_input('Enter expression :')
        expressions[i] = line

    return expressions


def evaluate_expressions(expressions):
    results = [0] * len(expressions)

    for idx, expr in enumerate(expressions):
        results[idx] = evaluate_expression(expr)

    return results


def evaluate_expression(expr):
    slave.stdin.write(expr + '\n')
    return slave.stdout.readline().rstrip()


def show_result(results):
    for result in results:
        result_dict = json.loads(result)
        print result_dict['result'], result_dict['duration']


def main():
    expressions = read_expressions()
    results = evaluate_expressions(expressions)
    show_result(results)


if __name__ == '__main__':
    main()
