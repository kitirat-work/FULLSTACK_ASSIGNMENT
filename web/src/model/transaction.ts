export type Transaction = {
  transaction_id: string;
  user_id: string | null;
  name: string | null;
  image: string | null;
  isBank: boolean | null;
};