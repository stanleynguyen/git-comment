const os = require('os');
const express = require('express');

const app = express();

app.get('/orgs/xendit/members', (req, res) => {
  res.end(`Hello from ${os.hostname()}`);
});

app.listen(process.env.PORT, () => {
  console.log(`UP AND RUNNING ON ${process.env.PORT}`);
});
