# Template generated code from trgen 0.16.4
JAR = C:/Users/Kenne/Downloads/antlr4-4.10.1-complete.jar
CLASSPATH = $(JAR)\;.
.SUFFIXES: .g4 .go
ANTLRGRAMMARS ?= $(wildcard *.g4)
GENERATED = parser/simpleboolean_parser.go parser/simpleboolean_lexer.go 
SOURCES = $(GENERATED) Test.go
default: classes
classes: setup $(SOURCES) program
setup:
	-export | grep GO
	-go env
	-go version
	export GO111MODULE=on; go get github.com/antlr/antlr4/runtime/Go/antlr@4.10
program:
	export GO111MODULE=on; go build Test.go
clean:
	rm -f *.tokens *.interp
	rm -f $(GENERATED)
run:
	export GO111MODULE=on; trwdog go run Test.go $(RUNARGS)
test: FORCE
	bash test.sh
parser/simpleboolean_parser.go : parser/SimpleBoolean.g4
	java -jar $(JAR) -encoding utf-8 -Dlanguage=Go -o parser -lib parser -package parser  $<
parser/simpleboolean_lexer.go : parser/SimpleBoolean.g4
	java -jar $(JAR) -encoding utf-8 -Dlanguage=Go -o parser -lib parser -package parser  $<

FORCE: ;
