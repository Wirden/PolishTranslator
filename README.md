# PolishTranslator

Here is a little project to calculate a serie of expressions written in RPN ( Reverse Polish Notation).

As a reminder this :

5 * ( 5 - 2 )

is equivalent to this in RPN :

5 5 2 - *

## Getting Started

Just clone the repo on your computer. Be sure you can execute Ruby and Python scripts. 
The Go script is compiled and should not require any specific environment. (Unless you mess around the src code and want to build it)

### How to use it

By calling the Ruby script with the right parameters.

Here is the default call when in the main folder

```
ruby ./IO/IO.rb  ../TEST_INPUT/TEST ../Result/testfile.txt

```

1st Argument : Input file
2nd Argument : Log file

### Input File

Here is how the input should look like : 

N - number of input expressions to solve
X1 - RPN expression to evaluate, separated by \n
X2
Xn

Example :

```
2 
3 4 + 
5 1 2 + 4 * + 3 - 
```

### Log File

The log file keeps a track of the Input expressions, the result and the calculated time.

Example (for the previous Example given) :

```
3 4 + is equals to 7, 0.00104999542236
5 1 2 + 4 * + 3 - is equals to 14, 0.00100016593933

It took the program 0.00215888023376 to calculate the entries. 
```

## Built With

* [Golang](https://golang.org/) - For the Worker
* [Python](https://www.python.org/) - For the API
* [Ruby](https://www.ruby-lang.org/) - For the IO

## Authors

* **Antoine Chevrot** - *Initial work* - [Wirden](https://github.com/Wirden)



