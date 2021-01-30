FROM php:8.0-fpm

RUN apt-get update && apt-get install -y \
        zsh \
        vim \
        libfreetype6-dev \
        libjpeg62-turbo-dev \
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
        iproute2 \
        npm

RUN pecl install xdebug \
    && docker-php-ext-enable xdebug \
    && docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-install -j$(nproc) intl \
    && docker-php-ext-install -j$(nproc) zip \
    && docker-php-ext-install -j$(nproc) pdo_mysql \
    && docker-php-ext-install -j$(nproc) exif \
    && docker-php-ext-install -j$(nproc) pcntl \
    && docker-php-ext-install -j$(nproc) bcmath \
    && docker-php-ext-install -j$(nproc) soap

RUN mv /usr/local/etc/php/php.ini-development /usr/local/etc/php/php.ini

RUN pear config-set php_ini /usr/local/etc/php/php.ini

# # PHP 7.4 and below
# RUN echo "xdebug.idekey=PHPSTORM" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
#     && echo "xdebug.mode=debug" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
#     && echo "xdebug.remote_autostart=off" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
#     && echo "xdebug.discover_client_host=1" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini

# PHP 8.0 and above
RUN echo "xdebug.mode=debug" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo "xdebug.start_with_request=yes" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini

WORKDIR /var/www

# Update npm
RUN npm i npm@latest -g

# Install composer
COPY install_composer.sh /root/
COPY composer /usr/local/bin/

RUN chmod 755 /root/install_composer.sh && mkdir /var/www/storage && mkdir /var/www/storage/logs && /root/install_composer.sh

RUN chmod 755 /usr/local/bin/composer && chmod 755 /usr/local/etc/php/composer.phar

RUN mkdir /home/root && chmod 777 /home/root

ENV HOME /home/root
