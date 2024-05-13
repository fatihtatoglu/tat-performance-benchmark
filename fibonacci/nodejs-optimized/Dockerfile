FROM node:alpine
WORKDIR /app
COPY . .
CMD [ "node", "--optimize-for-size", "--stack-size=102400", "main.js" ]