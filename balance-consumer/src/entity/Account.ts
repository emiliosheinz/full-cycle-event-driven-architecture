export class Account {
  constructor(
    public id: string,
    public balance: number,
    public updatedAt = new Date(),
  ) {}
}
