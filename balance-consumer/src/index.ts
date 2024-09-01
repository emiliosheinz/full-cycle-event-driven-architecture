import 'dotenv/config';

import express from 'express';

const app = express();
const port = process.env.PORT;

app.get('/balances/:accountId', (req, res) => {
  res.send(
    `You requested balances for account with ID: ${req.params.accountId}`,
  );
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
