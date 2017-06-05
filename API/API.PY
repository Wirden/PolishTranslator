import subprocess
import sys
import time

ILLEGAL_ARG_NUM_ERROR = 'Illegal number of arguments. Please enter the file to process as the first argument and the name of the log file as the second argument'
IO_ERROR = 'Impossible to open the file '
EXPR_NUM_ERROR = 'The number of expressions does not match the number provided'
VALUE_NUM_ERROR = 'The first line of the file must be an integer'
FINAL_TEXT = " It took the program {0} seconds to calculate the entries. "
WORKER_PATH = "../Worker/Worker"

start_time = time.time()

def execute(cmd):
	popen = subprocess.Popen(cmd,stdout=subprocess.PIPE)
	for stdout_line in iter(popen.stdout.readline, ""):
	    yield stdout_line 
	popen.stdout.close()
	popen.wait()


def Spliter(expr):

	try:
		f = open(expr)
	except IOError:
		print IO_ERROR, expr
		return (False,"")

	content = f.readlines()
	content = [x.strip() for x in content]
	f.close
	return (True,content)

def main(argv):
	if len(argv) != 2:
		print ILLEGAL_ARG_NUM_ERROR
		return

	file_to_process = argv[0]
	log_file = argv[1]

	spliter_bool,list_expression = Spliter(file_to_process)

	if not spliter_bool:
		return

	try:
		num_expr = int(list_expression[0])
	except ValueError:
		print VALUE_NUM_ERROR
		return
	
	if num_expr != (len(list_expression) - 1):
		print EXPR_NUM_ERROR
		return

	try:
		f1=open(log_file, 'a+')
	except IOError:
		print IO_ERROR, log_file
		return


	for expr in list_expression[1:]:
		loop_time = time.time()
		for path in execute([WORKER_PATH, expr]):
			log_text = path + ", " + str(time.time() - loop_time)
			f1.write(log_text + '\n')

	final_text = FINAL_TEXT.format(time.time() - start_time)
	f1.write('\n')
	f1.write(final_text + '\n')
	f1.close()



if __name__ == '__main__':
	main(sys.argv[1:])