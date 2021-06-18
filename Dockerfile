FROM golang:1.16-stretch
WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=1

RUN go install github.com/spf13/cobra/cobra@latest

CMD ["tail", "-f", "/dev/null"]