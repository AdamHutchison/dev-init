

# Set the Current Working Directory inside the container
WORKDIR /var/www/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

USER root

# install buffalo cli
RUN wget https://github.com/gobuffalo/cli/releases/download/v0.17.2/buffalo_0.17.2_Linux_x86_64.tar.gz \
    && tar -xvzf buffalo_0.17.2_Linux_x86_64.tar.gz \
    && mv buffalo /usr/local/bin/buffalo

USER www-data

# Run the executable
CMD ["buffalo-dev"]