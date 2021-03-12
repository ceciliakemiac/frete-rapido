FROM golang:1.14.1-alpine as builder

RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
WORKDIR /GreetingAPI
COPY . /greeting
WORKDIR /greeting
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o greeting

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main . 

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /greeting .
CMD ["./greeting"]

COPY --from=builder /dist/ /

ENTRYPOINT ["/main"] 
