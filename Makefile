benchmark:
	AWS_REGION=us-west-2 go test -bench=.

build:
	go build

run:
	./aws-golang-benchmark

