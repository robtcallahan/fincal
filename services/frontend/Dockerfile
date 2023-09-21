FROM node:alpine

ENV WORKINGDIR=/app/frontend

WORKDIR ${WORKINGDIR}
COPY . ${WORKINGDIR}

RUN npm install

ENTRYPOINT ["npm", "run", "dev"]
