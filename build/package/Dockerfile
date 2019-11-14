# build executable
FROM base-shameful-api-image:latest AS builder

# Building the api
RUN	CGOENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /go/bin/shameful-api ./...;

# Build smaller runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Copy our built executables
COPY --from=builder /go/bin/shameful-api /go/bin/shameful-api

# Run the API
CMD ["/go/bin/shameful-api"]
