FROM ubuntu:trusty

COPY docker-local/app/make-install-scripts /make-install-scripts
RUN chmod +x /make-install-scripts/*

# Install OS dependencies
RUN apt-get update && apt-get install -y \
    python \
    mc \
    curl \
    make \
    autoconf \
    gcc \
    wget \
    libmysqlclient-dev \
    libxml2-dev \
    libcurl4-openssl-dev \
    libpng-dev \
    libjpeg-dev \
    libmcrypt-dev \
    php5-mcrypt \
    libpcre3-dev \
    imagemagick \
    php5-imagick \
    php-pear \
    php5-dev \
    libmagickwand-dev \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

# Install PHP 5.5.18 from source
RUN mkdir -p /usr/php5/source-code
RUN wget --quiet --trust-server-names http://uk3.php.net/get/php-5.5.18.tar.gz/from/this/mirror -P /usr/php5/source-code

RUN tar -xvf /usr/php5/source-code/php-5.5.18.tar.gz -C /usr/php5/source-code

COPY docker-local/app/php-docker.ini /usr/php5/php.ini

RUN ./usr/php5/source-code/php-5.5.18/configure --prefix=/usr/php5 \
    --enable-bcmath \
    --enable-calendar \
    --enable-ftp \
    --enable-mbstring \
    --enable-opcache \
    --enable-soap \
    --enable-sockets \
    --enable-zip \
    --disable-fileinfo \
    --disable-flatfile \
    --disable-posix \
    --with-curl \
    --with-gd \
    --with-imagick \
    --with-openssl \
    --with-mcrypt \
    --without-sqlite3 \
    --without-pdo-sqlite \
    --without-imap \
    --without-kerberos \
    --without-imap-ssl\
    --without-interbase \
    --without-ldap \
    --without-mssql \
    --without-oci8 \
    --without-pgsql \
    --without-pear \
    --disable-phar \
    --without-snmp \
    --enable-mysqlnd \
    --with-pdo-mysql=mysqlnd \
    --with-mysqli=mysqlnd \
    --with-mysql=mysqlnd \
    --with-config-file-scan-dir=/usr/php5

RUN make
RUN make install

# Install X-Debug from source
RUN /make-install-scripts/make-install-xdebug2.sh

# Download Google App Engine (acts as local server)
RUN wget -P /gophr --quiet --trust-server-names https://storage.googleapis.com/appengine-sdks/featured/google_appengine_1.9.40.zip
RUN cd /gophr && unzip -q google_appengine_1.9.40.zip

WORKDIR /gophr/backend