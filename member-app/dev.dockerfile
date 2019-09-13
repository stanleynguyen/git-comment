FROM node
WORKDIR /app
RUN npm install
CMD npm run dev
ENV PORT 5001
EXPOSE ${PORT}
