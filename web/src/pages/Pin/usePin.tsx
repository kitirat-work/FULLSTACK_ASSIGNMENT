import * as React from "react";
import { useNavigate } from "react-router-dom";
import { useUserDispatch, useUserState } from "../../store/UserContext";
import { PinState } from "./model";
import AuthService from "../../service/AuthService";

export function usePin() {
  const { user } = useUserState();
	const dispatchUser = useUserDispatch();
  const navigate = useNavigate();
  const [state, setState] = React.useState<PinState>({
    pin: "",
    errorMessage: "",
  });

  const handleKeyPress = (key: string) => {
    const curPin = state.pin + key;
    if (key === "del") {
      setState((prevState) => ({
        ...prevState,
        pin: prevState.pin.slice(0, -1),
      }));
    } else if (curPin.length <= 6) {
      setState((prevState) => ({
        ...prevState,
        pin: prevState.pin + key,
      }));
    }

    if (curPin.length === 6) {
      handleSubmitPin(curPin);
    }
  };

  async function handleSubmitPin(pin: string) {
    try {
      console.log(user.userId, pin);
      
      const res = await AuthService.Login(user.userId, pin);
      if (!res.data) {
        throw new Error("Invalid PIN");
      }
      dispatchUser({ type: 'update', payload: res.data });
      navigate('/bank-main');
    } catch (error) {
      setState((prevState) => ({
        ...prevState,
        pin: "",
        errorMessage: error.message,
      }));
    }
  }

  React.useEffect(() => {
    document.title = "Pin Screen";
    if (!user) {
      navigate('/splash');
    }
    
  }, []);

  return { state, user, handleKeyPress };
};