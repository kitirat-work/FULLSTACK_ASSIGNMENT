import * as React from 'react';
import { User } from '../model/user';
import UserService from "../service/UserService";

/* Reducer */
export type Action = { type: 'update'; payload: User; };
export type Dispatch = (action: Action) => void;
export type State = { user: User | null; };

const UserStateContext = React.createContext<State | undefined>(undefined);
const UserDispatchContext = React.createContext<Dispatch | undefined>(undefined);

function userReducer(state: State, action: Action): State {
  console.log('userReducer', state, action);

  switch (action.type) {
    case 'update':
      console.log({ ...state, user: action.payload });

      return { ...state, user: action.payload };
    default:
      throw new Error(`Unhandled action type: ${action.type}`);
  }
}

/* Provider */
function UserProvider({ children }: { children: React.ReactNode; }) {
  const [state, dispatch] = React.useReducer(userReducer, { user: null });
  return (
    <UserStateContext.Provider value={state}>
      <UserDispatchContext.Provider value={dispatch}>
        {children}
      </UserDispatchContext.Provider>
    </UserStateContext.Provider>
  );
}

/* Hooks */
// const USER_ID = "000018b0e1a211ef95a30242ac180002";
// async function initUser() {
//   try {
//     const res: User = await UserService.GetUserById(USER_ID);
//     dispatchUser({ type: 'update', payload: res });

//   } catch (error) {
//     console.error(error);
//   }
// }

function useUserState() {
  const context = React.useContext(UserStateContext);
  if (context === undefined) {
    throw new Error('useUserState must be used within a UserProvider');
  }

  return context;
}

function useUserDispatch() {
  const context = React.useContext(UserDispatchContext);
  if (context === undefined) {
    throw new Error('useUserDispatch must be used within a UserProvider');
  }
  return context;
}

function userUserStore() {
  return [useUserState(), useUserDispatch()];
}

export { UserProvider, useUserState, useUserDispatch, userUserStore };
