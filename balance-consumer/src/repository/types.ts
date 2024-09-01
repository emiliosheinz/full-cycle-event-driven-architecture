import { Account } from '../entity/Account';

export interface AccountRepositoryInterface {
  findById: (id: string) => Promise<Account | null>;
}
