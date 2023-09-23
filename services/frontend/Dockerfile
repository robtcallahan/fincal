FROM node:alpine as base

FROM base as dev

ENV WORKINGDIR=/app/frontend

WORKDIR ${WORKINGDIR}
COPY . ${WORKINGDIR}

RUN npm install

ENTRYPOINT ["npm", "run", "dev"]
