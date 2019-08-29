# Spell Checker 

한글을 입력하면 원본 한글과 수정된 한글을 비교해서 출력해주는 프로그램입니다.

## Install dependencies

```go
$ go get github.com/sergi/go-diff/...
```

## Test

```go
$ go test ./...
```

## Run

```go
$ go run cmd/spellchecker/main.go "외않되"

# 외않되왜 안돼
```

## Sources

* https://github.com/sergi/go-diff
