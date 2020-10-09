# Start with the golang v1.13 image
FROM golang:1.13

# Set a working directory
WORKDIR /go/src/app

# Copy the entire project
COPY . .
# Install all go_modules
RUN go get -d -v ./...
RUN go install -v ./...

# Set environment variables
ENV HTTP_PORT=8080
ENV MONGO_SERVICE_HOST=
ENV MONGO_SERVICE_PORT=27017

# Expose the port
EXPOSE "8080:8080"

# Specify the command to run

CMD ["bin"]
