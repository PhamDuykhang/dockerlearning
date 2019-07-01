FROM mongo:3.4

RUN mkdir -p /data/db

VOLUME [ . "/data/db" ]

EXPOSE 27017

CMD [ "mongod" ]