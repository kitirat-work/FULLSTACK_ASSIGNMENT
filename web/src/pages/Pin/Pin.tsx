import * as React from "react";
import { FunctionComponent } from "react";
import { usePin } from "./usePin";

interface PinProps {

}

const Pin: FunctionComponent<PinProps> = () => {
  const { state, user, handleKeyPress } = usePin();
  return (
    <main className="container container--pin-type">
      <div className="pin">
        <div className="pin__top">
          <span className="pin__photo"><img src={user?.banner.image} alt="profile picture" /></span>
          <h1 className="pin__name">{user?.name}</h1>
          <p className="pin__dsc" style={{ display: state.errorMessage ? 'block' : 'none' }}>
            {state.errorMessage}<br />
            You have {state.attemptsLeft} attempts left.
          </p>
          <div className="pin__dots">
            {Array.from({ length: 6 }).map((_, index) => (
              <span key={index} className={`pin__dot ${index < state.pin.length ? 'is-filled' : ''}`}></span>
            ))}
          </div>
        </div>
        <div className="pin__btm">
          <a href="#" className="pin__login">Login with ID / Password </a>
          <span className="pin__kb">Powered by TestLab</span>
          <div className="pin__keys">
            {['1', '2', '3', '4', '5', '6', '7', '8', '9'].map((key) => (
              <button
                key={key}
                type="button"
                className="pin__key"
                onClick={() => handleKeyPress(key)}
              >
                {key}
              </button>
            ))}
            <span className="pin__key pin__key--space"></span>
            <button
              key="0"
              type="button"
              className="pin__key"
              onClick={() => handleKeyPress('0')}
            >
              0
            </button>
            <button
              type="button"
              className="pin__key pin__key--del"
              onClick={() => handleKeyPress('del')}
            >
              <span className="blind">Delete</span>
            </button>
          </div>
        </div>
      </div>
    </main>
  );
};


export default Pin;