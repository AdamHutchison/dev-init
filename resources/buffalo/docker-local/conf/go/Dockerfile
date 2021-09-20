

# Set the Current Working Directory inside the container
WORKDIR /var/www/

USER root

RUN mkdir /.cache && chmod -R 777 /.cache

USER ${UID}:${GID}

# install buffalo cli
RUN wget https://github.com/gobuffalo/cli/releases/download/v0.17.2/buffalo_0.17.2_Linux_x86_64.tar.gz \
    && tar -xvzf buffalo_0.17.2_Linux_x86_64.tar.gz \
    && mv buffalo /usr/local/bin/buffalo

ENV ADDR=0.0.0.0

# Run the executable
CMD [ "buffalo", "dev" ]