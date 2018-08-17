FROM golang:1.7.3 as golang-build
WORKDIR /go/src/test-genie/
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v gopkg.in/yaml.v2 
COPY . .
WORKDIR src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o test-genie .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /build/
COPY --from=golang-build /go/src/test-genie/ .
CMD ["./bin/test-genie"]
EXPOSE 8000