# syntax=docker/dockerfile:1

FROM golang:1.22.4

LABEL name="url-shortener-image"

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Install make
RUN apt update && apt install -y make

# Build the app
RUN make build

# When running the container, run "make run"
CMD pwd && echo "---" && ls -a && echo "---" && make run

