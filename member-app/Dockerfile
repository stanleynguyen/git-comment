FROM node
WORKDIR /app
COPY . .
RUN npm install --production
CMD npm start
ENV PORT 5001
EXPOSE ${PORT}
