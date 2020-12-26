FROM node:14
WORKDIR /app
COPY . /app
RUN npm i
RUN npm run build
EXPOSE 8080
ENTRYPOINT [ "node", "/app/dist/index.js"]