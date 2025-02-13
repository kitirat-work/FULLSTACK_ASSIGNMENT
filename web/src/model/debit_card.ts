export type DebitCard = {
  cardId: string;
  userId: string | null;
  name: string | null;

  // has one design
  design: DebitCardDesign;
  // has one detail
  details: DebitCardDetails;
  // has one status
  status: DebitCardStatus;
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