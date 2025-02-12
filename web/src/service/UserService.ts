import { User } from "../model/user";

const baseURL = process.env.API_URL + '/user';

async function getUserById(id: string): Promise<User> {
  const response = await fetch(`${baseURL}/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Failed to get user');
  }

  return response.json();
}


export default {
  getUserById,
};