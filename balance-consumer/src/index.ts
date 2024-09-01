import 'dotenv/config';

import express from 'express';
import { createPool } from 'mysql2/promise';
import { ListAccountBalanceUseCase } from './usecase/ListAccountBalance';
import { AccountRepository } from './repository/AccountRepository';

const database = createPool({
  host: 'balances-mysql',
  user: 'root',
  password: 'root',
  database: 'balances',
});

(async () => {
  await database.query(
    'CREATE TABLE IF NOT EXISTS accounts (id VARCHAR(255), balance FLOAT, updated_at DATETIME)',
  );
  await database.query(
    'INSERT INTO accounts (id, balance, updated_at) VALUES ("1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed", 100.0, "2024-06-12 10:0000")',
  );
})();

const app = express();
const port = process.env.PORT;

app.get('/balances/:accountId', async (req, res) => {
  try {
    const accountRepository = new AccountRepository(database);
    const listAccountBalanceUseCase = new ListAccountBalanceUseCase(
      accountRepository,
    );
    const result = await listAccountBalanceUseCase.execute({
      id: req.params.accountId,
    });
    res.status(200);
    res.send(result);
  } catch (e) {
    res.status(400);
    res.send({ error: (e as Error).message ?? 'Unexpected error' });
  }
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
