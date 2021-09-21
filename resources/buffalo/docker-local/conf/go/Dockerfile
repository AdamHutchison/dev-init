

# Set the Current Working Directory inside the container
WORKDIR /var/www/

RUN mkdir /.cache && chmod -R 777 /.cache

# install buffalo cli
RUN wget https://github.com/gobuffalo/cli/releases/download/v0.17.2/buffalo_0.17.2_Linux_x86_64.tar.gz \
    && tar -xvzf buffalo_0.17.2_Linux_x86_64.tar.gz \
    && mv buffalo /usr/local/bin/buffalo

# Install Soda cli
RUN go get github.com/gobuffalo/pop/... && go install github.com/gobuffalo/pop/soda@latest

ENV ADDR=0.0.0.0

# Run the executable
CMD [ "buffalo", "dev" ]