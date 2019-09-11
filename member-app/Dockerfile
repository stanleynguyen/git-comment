FROM node
WORKDIR /app
COPY . .
RUN npm install
CMD node index.js
ENV PORT 5001
EXPOSE ${PORT}
