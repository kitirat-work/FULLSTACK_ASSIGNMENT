import { FunctionComponent } from "react";
import { Account } from "../../model/account";
import * as React from "react";

interface AccountMainProps {
  account: Account;
}

const AccountMain: FunctionComponent<AccountMainProps> = (props) => {
  const { account } = props;
  return (
    <div className="main-acc is-small" style={{ backgroundColor: account.accountDetails.color }}>
      <div className="main-acc__top">
        <h2 className="main-acc__name">{account.type}</h2>
        <span className="main-acc__amount">฿{account.accountBalances.amount}</span>
        {account.accountFlags.map((flag) => (
          <span key={flag.flagId} className="main-acc__flag">{flag.flagValue}</span>
        ))}
      </div>
      <div className="main-acc__bottom">
        <span className="main-acc__detail">Smart account {account.accountNumber}</span>
        <span className="main-acc__detail">Powered by {account.issuer}</span>
      </div>
      <button type="button" className="main-acc__more main-acc__more--small">
        <span className="blind">More Action</span>
      </button>
      <div className="tooltip tooltip--sub-acc">
        <button type="button" className="tooltip__btn-more">Copy account number</button>
        <button type="button" className="tooltip__btn-more">Edit Name and Color</button>
      </div>
      <a href="#" className="main-acc__act main-acc__act--money">
        <span className="blind">Add money</span>
      </a>
      {account.accountDetails.progress === null ||
        <div className="main-acc__circle">
          <svg className="graph-bar" width="100%" height="100%" viewBox="0 0 42 42">
            <circle cx="21" cy="21" r="15.91549430918954" fill="transparent" stroke="rgba(0,0,0,0.07)"
              strokeWidth="1.5"></circle>
            <circle className="gauge" cx="21" cy="21" r="15.91549430918954" fill="transparent" stroke="#fff"
              strokeWidth="1.5" strokeLinecap="round" strokeDashoffset="25"
              style={{ strokeDasharray: `${account.accountDetails.progress} ${100 - account.accountDetails.progress}` }}></circle>
          </svg>
          <div className="main-acc__num">
            <span className="percent">{account.accountDetails.progress}</span>
            <span className="unit">%</span>
          </div>
        </div>
      }
    </div>
  );
};

export default AccountMain;