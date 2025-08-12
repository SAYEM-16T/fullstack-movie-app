FROM node:18-alpine

WORKDIR /app

COPY ./node-backend/package*.json ./

RUN npm install

COPY ./node-backend ./
COPY ./frontend ./frontend

EXPOSE 3100

CMD [ "npm", "start" ]
