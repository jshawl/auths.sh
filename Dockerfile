FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o ./ssh
COPY www ./www
EXPOSE 4202
EXPOSE 8080
ENV HOST_PRIVATE_KEY ""
CMD ["/app/ssh"]
