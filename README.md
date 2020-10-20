# Find-and-replace patterns for the Go language. 

Find curated `comby` patterns for replacing certain Go patterns in this repo. These patterns aim to:

- generate virtually zero false positives (i.e., every match should be something you can act on improving)
- match/replace code that can be objectively better (i.e., not stylistically controversial)

If you find that a pattern doesn't do this effectively, please file an issue.

## Current patterns

- `go-staticcheck.toml`: a subset of 22 patterns inspired by simple [staticcheck](https://staticcheck.io/docs/checks) patterns.

## Running

- `comby -config go-staticcheck.toml -f .go` finds matches
- `comby -config go-staticcheck.toml -f .go -i` replaces file contents
  - **You should run `gofmt` after running the above**

Exclude `vendor` and `.` directories with:

- `comby -config go-staticcheck.toml -f .go -i -exclude-dir vendor,.`
