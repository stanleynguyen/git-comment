const app = require('./src/app');

if (process.env.NODE_ENV === 'production') {
  app.inProdMode().start();
} else {
  app.inDevMode().start();
}
