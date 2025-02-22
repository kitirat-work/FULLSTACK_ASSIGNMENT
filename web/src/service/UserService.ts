import { User } from "../model/user";

const baseURL = process.env.API_URL + '/user';

async function GetUserById(id: string): Promise<{data: User}> {
  const response = await fetch(`${baseURL}/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  const result = await response.json();

  if (!response.ok) {
    throw new Error(result.message);
  }

  return result;
}



export default {
  GetUserById,
};