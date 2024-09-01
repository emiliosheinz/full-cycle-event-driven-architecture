import { AccountRepositoryInterface } from '../repository/types';

type ListAccountBalanceInputDTO = {
  id: string;
};

type ListAccountBalanceOutputDTO = {
  balance: number;
};

export class ListAccountBalanceUseCase {
  constructor(private accountRepository: AccountRepositoryInterface) {}

  public async execute({
    id,
  }: ListAccountBalanceInputDTO): Promise<ListAccountBalanceOutputDTO> {
    const account = await this.accountRepository.findById(id);
    if (!account) {
      throw new Error(`Account with id ${id} not found.`);
    }
    return { balance: account.balance };
  }
}
