FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/configuration ./configuration
COPY --from=builder /app/certificate ./certificate
COPY --from=builder /app/domain ./domain
COPY --from=builder /app/infrastructure ./infrastructure
COPY --from=builder /app/interactor ./interactor
COPY --from=builder /app/repository ./repository
COPY --from=builder /app/usecase ./usecase
COPY --from=builder /app/main .

CMD ["./main"]