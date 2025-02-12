import { Banner } from "./banner";

export type User = {
  user_id: string;
  name: string;
  //has one banner
  banner: Banner;
}