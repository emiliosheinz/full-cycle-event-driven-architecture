import { Account } from './Account';

describe('Account Entity', () => {
  test('should create a new Account', () => {
    const account = new Account('1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed', 100);
    expect(account.id).toEqual('1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed');
    expect(account.balance).toEqual(100);
    expect(account.updatedAt).toEqual(expect.any(Number));
  });

  test('should create a new account with provided updated at value', () => {
    const updatedAt = Date.now();
    const account = new Account(
      '1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed',
      100,
      updatedAt,
    );
    expect(account.updatedAt).toEqual(updatedAt);
  });
});
