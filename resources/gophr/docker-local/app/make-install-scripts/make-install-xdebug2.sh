#!/bin/sh

mkdir -p /usr/xdebug2 && cd /usr/xdebug2 || exit

wget --quiet --trust-server-names http://xdebug.org/files/xdebug-2.2.5.tgz

tar -xzf xdebug-2.2.5.tgz > /dev/null
cd /usr/xdebug2/xdebug-2.2.5 || exit

/usr/php5/bin/phpize
./configure --enable-xdebug --with-php-config=/usr/php5/bin/php-config

make
make install

make test
