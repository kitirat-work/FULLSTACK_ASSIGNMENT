import { Account } from "./account";
import { DebitCard } from "./debit_card";

export type User = {
  userId: string;
  name: string;

  //has many banner
  banner: Banner[];
  //has one user greeting
  userGreeting?: UserGreeting;
  //has many accounts
  accounts: Account[];
  //has many debit cards
  debitCards: DebitCard[];
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