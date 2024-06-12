# syntax=docker/dockerfile:1

FROM golang:1.22.4

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Install make
RUN apt update && apt install -y make

# Build the app
RUN make build

# Expose the port
EXPOSE 8080

# When running the container, run "make run"
CMD ["make", "run"]

