import subprocess
import sys

ILLEGAL_ARG_NUM_ERROR = 'Illegal number of arguments. Please enter the file to process as the first argument and the name of the log file as the second argument'
IO_ERROR = 'Please enter a valid file name'
EXPR_NUM_ERROR = 'The number of expressions does not match the number provided'
VALUE_NUM_ERROR = 'The first line of the file must be an integer'
WORKER_PATH = "../Worker/Worker"

def execute(cmd):
	popen = subprocess.Popen(cmd,stdout=subprocess.PIPE)
	for stdout_line in iter(popen.stdout.readline, ""):
	    yield stdout_line 
	popen.stdout.close()
	popen.wait()


def Spliter(expr):
	try:
		with open(expr) as f:
			content = f.readlines()
			content = [x.strip() for x in content]
			return content
	except IOError:
		print IO_ERROR
	finally:
		f.close()


def main(argv):
	if len(argv) != 2:
		print ILLEGAL_ARG_NUM_ERROR
		return

	file_to_process = argv[0]
	log_file = argv[1]

	list_expression = Spliter(file_to_process)

	try:
		num_expr = int(list_expression[0])
	except ValueError:
		print VALUE_NUM_ERROR
		return
	
	if num_expr != (len(list_expression) - 1):
		print EXPR_NUM_ERROR
		return

	for expr in list_expression[1:]:
		for path in execute([WORKER_PATH, expr]):
			f1=open(log_file, 'a+')
			f1.write(path)
			f1.close()



if __name__ == '__main__':
	main(sys.argv[1:])