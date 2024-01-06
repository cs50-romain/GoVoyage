# Golang image
FROM golang:1.21

# Create app directory
WORKDIR /app

# Copy everything from root dir to /app
COPY . .

# Install Go dependency
RUN go mod download

# Build app with optional configs
RUN go build -o /govoyage

# Tell Docker which netwrok to listen on
EXPOSE 8080

# Run executable command
CMD [ "/govoyage" ]
