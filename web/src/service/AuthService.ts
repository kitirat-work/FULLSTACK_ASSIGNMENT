import { User } from "../model/user";

const baseURL = process.env.API_URL + '/auth';

async function Login(id: string, pin: string): Promise<{data: User}> {
  const response = await fetch(`${baseURL}/login/pin`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ id, pin }),
  });

  const result = await response.json();
  if (!response.ok) {
    throw new Error(result.message);
  }

  return result;
}


export default {
  Login,
};