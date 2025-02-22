export type Account = {
  accountId: string;
  userId: string | null;
  type: string | null;
  currency: string | null;
  accountNumber: string | null;
  issuer: string | null;

  // has many flags
  accountFlags: AccountFlag[];
  // has one detail
  accountDetails: AccountDetail;
  // has one balance
  accountBalances: AccountBalance;
};



export type AccountFlag = {
  flagId: number;
  accountId: string;
  userId: string;
  flagType: string;
  flagValue: string;
  createdAt: Date;
  updatedAt: Date;
};

export type AccountDetail = {
  accountId: string;
  userId: string | null;
  color: string | null;
  isMainAccount: boolean | null;
  progress: number | null;
};



export type AccountBalance = {
  accountId: string;
  userId: string | null;
  amount: number | null;
};