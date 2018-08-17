FROM golang:1.7.3 as golang-build
WORKDIR /go/src/test-genie/
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v gopkg.in/yaml.v2 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./bin/test-genie ./src/*.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/test-genie/
COPY --from=golang-build /go/src/test-genie/ .
ENTRYPOINT ["/bin/sh", "entrypoint.sh"]
EXPOSE 9000