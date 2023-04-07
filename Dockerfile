FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o ./ssh
EXPOSE 4202
CMD ["/app/ssh"]
