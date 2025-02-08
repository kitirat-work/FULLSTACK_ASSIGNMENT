import * as React from "react";
import { FunctionComponent } from "react";

interface PinProps {

}

const Pin: FunctionComponent<PinProps> = () => {
  React.useEffect(() => {
    document.title = "Pin Screen";
  }, []);

  return (
    <div className="wrap">
      {/* <!-- container --> */}
      <main className="container container--pin-type">
        {/* <!-- Enter PIN --> */}
        <div className="pin">
          <div className="pin__top">
            <span className="pin__photo"><img src="https://dummyimage.com/200x200/999/fff" alt="" /></span>
            <h1 className="pin__name">Interview User</h1>
            <p className="pin__dsc" style={{ display: 'none' }}>Invalid PIN Code.<br />You have 3 attempts left.</p>
            <div className="pin__dots">
              <span className="pin__dot is-filled"></span>
              <span className="pin__dot is-filled"></span>
              <span className="pin__dot"></span>
              <span className="pin__dot"></span>
              <span className="pin__dot"></span>
              <span className="pin__dot"></span>
            </div>
          </div>
          <div className="pin__btm">
            <a href="#" className="pin__login">Login with ID / Password </a>
            <span className="pin__kb">Powered by TestLab</span>
            <div className="pin__keys">
              <button type="button" className="pin__key">1</button>
              <button type="button" className="pin__key">2</button>
              <button type="button" className="pin__key">3</button>
              <button type="button" className="pin__key">4</button>
              <button type="button" className="pin__key">5</button>
              <button type="button" className="pin__key">6</button>
              <button type="button" className="pin__key">7</button>
              <button type="button" className="pin__key">8</button>
              <button type="button" className="pin__key">9</button>
              <span className="pin__key pin__key--space"></span>
              <button type="button" className="pin__key">0</button>
              <button type="button" className="pin__key pin__key--del"><span className="blind">Delete</span></button>
            </div>
          </div>
        </div>
        {/* <!-- //Enter PIN --> */}
      </main>
      {/* <!-- //container --> */}
    </div>
  );
};

export default Pin;