# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM ubuntu:latest

# Add Maintainer Info
LABEL maintainer="Alexandr Usanin <hidan9812@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./api .

# Environment Variables
ENV APP_SECRET=mega_secret_key_146
ENV APP_PORT=8080
ENV USE_SYSLOG=0
ENV LOG_DIR=/run/syslog-ng/json-log.fifo

CMD ./api