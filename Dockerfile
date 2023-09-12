FROM golang as compiler
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY clowder-hello.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /clowder-hello

FROM scratch 
COPY --from=compiler /clowder-hello . 
CMD ["./clowder-hello"]