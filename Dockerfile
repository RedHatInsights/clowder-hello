FROM golang as compiler 
RUN CGO_ENABLED=0 go get -a -ldflags '-s' \ 
github.com/redhatinsights/clowder-hello

FROM scratch 
COPY --from=compiler /go/bin/clowder-hello . 
CMD ["./clowder-hello"]