import { Account } from '../entity/Account';
import { AccountRepositoryInterface } from '../repository/types';

type UpdateAccountBalanceInputDTO = {
  id: string;
  balance: number;
};

export class UpdateAccountBalanceUseCase {
  constructor(private accountRepository: AccountRepositoryInterface) {}

  public async execute({
    id,
    balance,
  }: UpdateAccountBalanceInputDTO): Promise<void> {
    const account = new Account(id, balance);
    await this.accountRepository.save(account);
  }
}
