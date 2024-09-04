FROM mysql:5.7

COPY init.sh /docker-entrypoint-initdb.d/
COPY init.sql /app/

RUN chmod +x /docker-entrypoint-initdb.d/init.sh
