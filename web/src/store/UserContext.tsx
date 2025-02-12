import * as React from 'react';
import { User } from '../model/user';

/* Reducer */
export type Action = { type: 'update'; payload: User };
export type Dispatch = (action: Action) => void;
export type State = { user: User | null };

const UserStateContext = React.createContext<State | undefined>(undefined);
const UserDispatchContext = React.createContext<Dispatch | undefined>(undefined);

function userReducer(state: State, action: Action): State {
  switch (action.type) {
    case 'update':
      return { ...state, user: action.payload };
    default:
      throw new Error(`Unhandled action type: ${action.type}`);
  }
}

/* Provider */
function UserProvider({ children }: { children: React.ReactNode }) {
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

/* 
  Usage:
  const [userState, userDispatch] = userUserStore();
  userDispatch({ type: 'update', payload: user });

  const { user } = useUserState();

  return (
    <div>
      <h1>{user.name}</h1>
    </div>
  );

  <UserProvider>
    <App />
  </UserProvider>

*/

export { UserProvider, useUserState, useUserDispatch, userUserStore };
