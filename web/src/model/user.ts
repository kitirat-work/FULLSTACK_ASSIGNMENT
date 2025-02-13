import { Account } from "./account";
import { DebitCard } from "./debit_card";

export type User = {
  user_id: string;
  name: string;
  
  //has many banner
  banner: Banner[];
  //has one user greeting
  userGreeting: UserGreeting;
  //has many accounts
  accounts: Account[];
  //has many debit cards
  debitCards: DebitCard[];
}

export type Banner = {
  banner_id: string;
  user_id: string;
  title: string;
  description: string;
  image: string;
}

export type UserGreeting = {
  user_id: string;
  greeting: string;
}