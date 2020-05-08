FROM node:10.16-alpine

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json /usr/src/app/

RUN npm config set unsafe-perm true

# Since we use babel so npm prod and dev production
RUN npm i npm@latest -g && \
    npm install && \
    npm audit fix

COPY . /usr/src/app/

CMD [ "npm", "start" ]

