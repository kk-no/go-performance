package := ""

.PHONY: bench
bench:
	go test -bench . ./... -benchmem

.PHONY: bencho
bencho:
	cd ${package} && go test -bench . ./... -benchmem

.PHONY: bencht
bencht:
	cd ${package} && go test -bench . ./... -benchmem -trace bench.trace

.PHONY: benchp
benchp:
	cd ${package} && go test -bench . ./... -benchmem -cpuprofile bench.prof

.PHONY: trace
trace:
	go tool trace --http localhost:6060 ${package}/bench.trace

.PHONY: pprof
pprof:
	go tool pprof -http :6070 ${package}/bench.prof