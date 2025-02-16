export type DebitCard = {
  cardId: string;
  userId: string | null;
  name: string | null;

  // has one design
  debitCardDesign: DebitCardDesign;
  // has one detail
  debitCardDetails: DebitCardDetails;
  // has one status
  debitCardStatus: DebitCardStatus;
};

export type DebitCardDesign = {
  cardId: string;
  userId: string | null;
  color: string | null;
  borderColor: string | null;
};

export type DebitCardDetails = {
  cardId: string;
  userId: string | null;
  issuer: string | null;
  number: string | null;
};

export type DebitCardStatus = {
  cardId: string;
  userId: string | null;
  status: string | null;
};