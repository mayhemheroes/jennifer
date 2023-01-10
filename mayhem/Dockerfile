FROM golang:1.18 as builder

COPY . /jennifer
WORKDIR /jennifer/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzJen.a && \
    clang -fsanitize=fuzzer fuzzJen.a -o fuzzJen.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /jennifer/mayhem/fuzzJen.libfuzzer /