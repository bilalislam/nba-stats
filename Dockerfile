FROM golang:latest as builder

RUN mkdir -p /app
WORKDIR /app

# Force the go compiler to use modules
ENV GO111MODULE on

# <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

RUN chmod +x /app
# Compile application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o messagebridge-api

#Image Diff
#(Not Scratch) 1.23GB
#(Scratch    ) 34.3MB
# <- Second step to build minimal image
FROM scratch
WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app .
# Execite application when container is started
EXPOSE 80/tcp
CMD ["./nba-stats"]
EXPOSE 8080