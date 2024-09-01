import { RowDataPacket } from 'mysql2';
import { Account } from '../entity/Account';
import { Database } from '../types/Database';
import { AccountRepositoryInterface } from './types';

export class AccountRepository implements AccountRepositoryInterface {
  constructor(private db: Database) {}

  public async findById(id: string): Promise<Account | null> {
    const [[result]] = await this.db.query<[RowDataPacket]>(
      'SELECT id, balance, updated_at FROM accounts WHERE id = ?',
      [id],
    );
    if (!result) return null;
    return new Account(result.id, result.balance, result.date);
  }
}
