#### Go Testing Commands

1. Run all test
```
go test
```
2. Test Specific package
```
go test {pkg1} {pkg2}...
```
3. Run tests in current and decendent directory
```
go test ./...
```
4. Generate verbose output
```
go test -v
```
5. Run only tests matching
```
go test -run {regexp}
```
6. Generate Test Coverage
```
go test -cover
```
7. Generatge test cover profile to find exact functions missing test coverage.
```
go test -coverprofile cover.out
go tool cover -func cover.out
go tool cover -html cover.out
go test -coverprofile count.out -covermode count
go tool cover -html cover.out
```

8. Benchmark testing including other test
```
go test -bench .
go test -bench . -benchtime 10s

```
9. Profile testing - block, cover, cpu, mem, mutex
```
go test -bench . -benchmem
go test -bench SHA1 -benchmem
go test -trace {trace.out}
go test -{type}profile {file}
go test -bench Alloc -memprofile profile.out
## install choco install graphviz for graph
go tool pprof profile.out 
type help to more info. like run svg
```