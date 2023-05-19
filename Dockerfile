FROM golang:1.20.4

# Set working directory according to linux 
WORKDIR /usr/src/app

# Hot Reload 
RUN go install github.com/cosmtrek/air@latest 

# Copy application files
COPY . .

# Ensure required packages are installed
RUN go mod tidy

