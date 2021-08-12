

RUN apt-get update && apt-get install -y \
    zsh \
    vim \
    cron \
    libfreetype6-dev \
    libjpeg62-turbo-dev \
    libmagickwand-dev \
    libmcrypt-dev \
    libpng-dev \
    libicu-dev \
    libpq-dev \
    libxpm-dev \
    libvpx-dev \
    libzip-dev \
    libwebp-dev \
    libxml2-dev \
    gnupg \
    gnupg1 \
    gnupg2 \
    wget \
    unzip \
    git \
    iputils-ping \
    iproute2

# Build imagik from source until pecl release is created (date today 13-April-2021)
RUN git clone https://github.com/Imagick/imagick \
    && cd imagick \
    && phpize && ./configure \
    && make \
    && make install

RUN docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-install -j$(nproc) intl \
    && docker-php-ext-install -j$(nproc) zip \
    && docker-php-ext-install -j$(nproc) pdo_mysql \
    && docker-php-ext-install -j$(nproc) exif \
    && docker-php-ext-install -j$(nproc) pcntl \
    && docker-php-ext-install -j$(nproc) bcmath \
    && docker-php-ext-install -j$(nproc) soap

COPY resources/php.ini /usr/local/etc/php/php.ini
COPY resources/docker-cron-entrypoint /usr/local/bin/docker-cron-entrypoint
COPY resources/crontab /crontab

RUN chmod 777 /usr/local/bin/docker-cron-entrypoint

RUN pear config-set php_ini /usr/local/etc/php/php.ini

WORKDIR /var/www

# Install composer
COPY resources/install_composer.sh /root/
COPY resources/composer /usr/local/bin/

RUN chmod 755 /root/install_composer.sh && mkdir /var/www/storage && mkdir /var/www/storage/logs && /root/install_composer.sh

RUN chmod 755 /usr/local/bin/composer && chmod 755 /usr/local/etc/php/composer.phar

RUN mkdir /home/root && chmod 777 /home/root

ENV HOME /home/root

# Setup cron pointing to artisan:schedule
# We'll only register the cronfile if artisan is present (to allow for auto-registering cron with laravel)
ONBUILD RUN if [ -f "/var/www/artisan" ]; then /usr/bin/crontab /crontab; fi
