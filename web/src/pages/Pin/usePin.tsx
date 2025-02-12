import * as React from "react";
import { useNavigate } from "react-router-dom";
import UserService from "../../service/UserService";
import { useUserState } from "../../store/UserContext";
import { PinState } from "./model";

export function usePin() {
  const { user } = useUserState();
  const navigate = useNavigate();
  const [state, setState] = React.useState<PinState>({
    pin: "",
    attemptsLeft: 3,
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
      await UserService.Login(pin);
      navigate('/bank-main');
    } catch (error) {
      setState((prevState) => ({
        ...prevState,
        pin: "",
        attemptsLeft: prevState.attemptsLeft - 1,
        errorMessage: "Invalid PIN",
      }));

      if (state.attemptsLeft === 1) {
        navigate('/splash');
      }
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