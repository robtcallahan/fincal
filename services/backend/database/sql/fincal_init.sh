#!/bin/sh

source ../.env.sh

mysql << EOF
ALTER USER 'root'@'localhost'
  IDENTIFIED WITH caching_sha2_password BY "${MYSQL_ROOT_PASSWORD}";

CREATE USER "${MYSQL_USER}"@'localhost'
  IDENTIFIED WITH caching_sha2_password BY "${MYSQL_PASSWORD}";
EOF

mysql --user=${MYSQL_USER} --password=${MYSQL_PASSWORD} < /data/fincal_init.sql