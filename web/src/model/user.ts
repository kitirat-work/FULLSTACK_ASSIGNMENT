import { Account } from "./account";
import { DebitCard } from "./debit_card";
import { Transaction } from "./transaction";

export type User = {
  userId: string;
  name: string;

  //has many banner
  banners: Banner[];
  //has one user greeting
  userGreetings?: UserGreeting;
  //has many accounts
  accounts: Account[];
  //has many debit cards
  debitCards: DebitCard[];
  //has many transactions
  transactions: Transaction[];
}

export type Banner = {
  bannerId: string;
  userId: string;
  title: string;
  description: string;
  image: string;
}

export type UserGreeting = {
  userId: string;
  greeting: string;
}