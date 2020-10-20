#!/bin/bash

comby -config ../go-staticcheck.toml -f test.go -stdout > test.tmp.go

diff test.tmp.go test.expect.go
ret=$?
if [ $ret -ne 0 ]; then
    echo "diff output is wrong"
    echo "update with: comby -config ../go-staticcheck.toml -f test.go -stdout > test.expect.go"
    rm test.tmp.go
    exit 1
fi

gofmt test.tmp.go > /dev/null

ret=$?
if [ $ret -ne 0 ]; then
    echo "gofmt should succeed"
    rm test.tmp.go
    exit 1
fi

rm test.tmp.go
