#!/bin/bash

set -e

docker-entrypoint.sh mysqld &

sleep 10

mysql -u myuser -ppass greetings < /app/init.sql

wait
