
FROM golang:1.21.6

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o houseFayzulin ./app.go

CMD ["./houseFayzulin"]