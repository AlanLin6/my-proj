# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /workspace

# Copy the local package files to the container's workspace
COPY . .

# Copy the pre-built binary into the container
COPY bazel-bin/calorie_calculator_/calorie_calculator .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./calorie_calculator"]
