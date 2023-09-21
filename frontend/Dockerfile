FROM node:alpine

ENV WORKINGDIR=/src/frontend

WORKDIR ${WORKINGDIR}
COPY . ${WORKINGDIR}

RUN npm install

ENTRYPOINT ["npm", "run", "dev"]
