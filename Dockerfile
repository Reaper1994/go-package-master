# Use the official Golang image with Alpine Linux as the base image for building the application.
# Alpine Linux provides a smaller footprint compared to the standard Golang image.
FROM golang:1.24rc2-alpine

# Set the working directory inside the container to /app.
# All subsequent commands will be run from this directory.
WORKDIR /app

# Initialize a new Go module.
# These files lock the versions of dependencies and are used by go mod tidy to fetch the exact versions.
RUN go mod init github.com/Reaper1994/go-package-master

# Run go mod tidy to download the exact versions of dependencies specified in go.mod and go.sum.
# This step cleans up the go.mod and go.sum files by removing unused modules.
RUN go mod tidy

# Copy the rest of the application code into the container.
# This includes your Go source files, tests, and any other resources your application needs.
COPY . .

# Compile the Go application into a binary named "main".
# The ./main.go argument specifies the entry point of your application.
RUN go build -o main cmd/main.go

# Make the "main" binary executable.
# This step is necessary because the binary is built without executable permissions by default.
RUN chmod +x main

# Inform Docker that the container listens on the specified network ports at runtime.
# Here, port 8080 is exposed, but make sure your application actually listens on this port.
EXPOSE 8080

# OPTIONAL: Set environment variables for Treblle API keys
# These environment variables are used in your Go application to configure Treblle for API ops.
# Set default environment variables (can be overridden by CI/CD secrets or runtime configuration)
ENV TREBLLE_API_KEY=default_key
ENV TREBLLE_PROJECT_ID=default_project_id

# Specify the command to run your application.
# This command is executed when the container starts.
CMD [ "./main" ]