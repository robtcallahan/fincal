FROM mysql:8.0

COPY fincal_init.sql /tmp
COPY fincal_init2.sql /tmp
COPY mysqld.cnf /etc/mysqld

CMD ["mysqld"]
CMD ["mysql -u root < /tmp/fincal_init2.sql"]
