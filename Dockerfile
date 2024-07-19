# Use golang for builder
FROM golang:1.16-alpine AS build
# Pass in the function to determine path
ARG FUNCTION_NAME=front-end
# Use a standard workdir
WORKDIR /app
# Copy all of the function-specific files to be built
COPY /go/$FUNCTION_NAME .
# pull the packages needed
RUN go mod tidy
# Build a statically linked result
RUN CGO_ENABLED=0 go build -o main .

# Final image based on 2MB distroless static
FROM gcr.io/distroless/static-debian11
# Copy artifact over
COPY --from=build /app/main /app
# Set artifact as entryptoin allowing params in child containers if needed
ENTRYPOINT ["/app"]