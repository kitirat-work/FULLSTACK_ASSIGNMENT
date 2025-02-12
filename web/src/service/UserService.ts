import { Banner } from "../model/banner";
import { User } from "../model/user";

const baseURL = process.env.API_URL + '/user';

async function GetUserById(id: string): Promise<User> {
  
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      const user: User = {
        user_id: "user_id",
        name: "name",
        banner: {
          banner_id: "banner_id",
          user_id: "user_id",
          title: "title",
          description: "description",
          image: "image",
        } as Banner,
      };
      resolve(user);
    }, 1000);
  });
  // const response = await fetch(`${baseURL}/${id}`, {
  //   method: 'GET',
  //   headers: {
  //     'Content-Type': 'application/json',
  //   },
  // });

  // if (!response.ok) {
  //   throw new Error('Failed to get user');
  // }

  // return response.json();
}

async function Login(pin: string): Promise<void> {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (pin !== '111111') {
        reject(new Error('Invalid PIN'));
      }
      resolve();
    }, 1000);
  });
  // const response = await fetch(`${baseURL}/login`, {
  //   method: 'POST',
  //   headers: {
  //     'Content-Type': 'application/json',
  //   },
  //   body: JSON.stringify({ pin }),
  // });

  // if (!response.ok) {
  //   throw new Error('Failed to login');
  // }

  // return response.json();
}


export default {
 GetUserById,
  Login,
};