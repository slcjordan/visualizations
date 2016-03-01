# visualizations
An example project showing how to use pprof visualizations for go.

Setup

torch
```
$ go get github.com/uber/go-torch
$ git clone git@github.com:brendangregg/FlameGraph.git
$ export PATH=$PATH:/path/to/FlameGraph
$ ./build
$ ./torch
```

gom
```
$ go get go get github.com/rakyll/gom/cmd/gom
$ go get go get github.com/rakyll/gom/http
$ gom
```

debugcharts
```
$ go get -v -u github.com/mkevac/debugcharts
```
then navigate to localhost:6060
