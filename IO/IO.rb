entryFile = ARGV[0]
outputFile = ARGV[1]

 PYNAME = '../API/API.PY'
# ARGS = '../TEST_INPUT/TEST ../Result/tstfile.txt'
python_output =  `python #{PYNAME} #{entryFile} #{outputFile}`

puts python_output
