#

make
export GO111MODULE=on
go run Test.go -file ../examples/t13.txt &
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10

echo When you are finished, make sure to kill the '"go"' process.
