import 'dotenv/config';

import { createPool } from 'mysql2/promise';
import express from 'express';

import { ListAccountBalanceUseCase } from './usecase/ListAccountBalance';
import { AccountRepository } from './repository/AccountRepository';
import { KafkaConsumer } from './libs/kafka/Consumer';
import { Topic } from './libs/kafka/Topic';
import { Message } from './libs/kafka/Message';
import { BalanceUpdatedPayload } from './libs/kafka/MessagePayload';
import { UpdateAccountBalanceUseCase } from './usecase/UpdateAccountBalance';

/** Database setup */
const database = createPool({
  host: 'balances-mysql',
  user: 'root',
  password: 'root',
  database: 'balances',
});
(async () => {
  await database.query(
    'CREATE TABLE IF NOT EXISTS accounts (id VARCHAR(255), balance FLOAT, updated_at DATETIME, PRIMARY KEY (id))',
  );
})();

/** Kafka consumer */
const consumer = new KafkaConsumer();
consumer.consume(Topic.Balances, ({ name, payload }) => {
  if (name === Message.BalanceUpdated) {
    const {
      account_id_from,
      account_id_to,
      ballance_account_from,
      ballance_account_to,
    } = payload as BalanceUpdatedPayload;
    const accountRepository = new AccountRepository(database);
    const updateAccountBalanceUseCase = new UpdateAccountBalanceUseCase(
      accountRepository,
    );
    updateAccountBalanceUseCase.execute({
      id: account_id_from,
      balance: ballance_account_from,
    });
    updateAccountBalanceUseCase.execute({
      id: account_id_to,
      balance: ballance_account_to,
    });
  }
});

/** Express setup */
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
