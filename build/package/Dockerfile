# build executable
FROM golang:1.13-alpine AS builder

# Install git because go get needs it for geting dependencies
RUN apk update; \
    apk add --no-cache git;

WORKDIR ~/projects/shameful-api/
COPY . .

# Fetching dependencies using go modules (>= go1.11)
RUN go get -v -d; \
	# Building the api
	CGOENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /go/bin/shameful-api .;

# Build smaller runtime image
FROM scratch

# Copy our built executables
COPY --from=builder /go/bin/shameful-api /go/bin/shameful-api

# Run the API
ENTRYPOINT ["/go/bin/shameful-api"]
