import subprocess


def execute(cmd):
	popen = subprocess.Popen(cmd,stdout=subprocess.PIPE)
	for stdout_line in iter(popen.stdout.readline, ""):
	    yield stdout_line 
	popen.stdout.close()
	return_code = popen.wait()


for path in execute(["../Worker/Worker", "5 8 + *"]):
	f1=open('../Result/testfile.txt', 'a+')
	f1.write(path)
	f1.close()
