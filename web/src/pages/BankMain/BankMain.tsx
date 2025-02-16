
import * as React from "react";
import { FunctionComponent } from "react";
import { useUserState } from "../../store/UserContext";
import { useNavigate } from "react-router-dom";
import { Transaction } from '../../model/transaction';
import AccountMain from "./AccountMain";

interface BankMainProps {

}

const BankMain: FunctionComponent<BankMainProps> = () => {
  const { user } = useUserState();
  const navigate = useNavigate();

  React.useEffect(() => {
    document.title = "Bank Main Screen";
    if (!user) {
      navigate('/splash');
    }
    console.log("User", user);

  }, []);

  return (
    <>

      <header className="header ">
        <a href="#" className="header__lft header__menu"><span className="blind">Menu</span></a>
        <button type="button" className="header__rgt header__cxl"><span className="blind">Cancel</span></button>
      </header>

      <main className="container container--main">
        <div className="content_wrap">
          <div className="main-top">
            <h1 className="main-top__tit main-loading main-loading--order1">Have a nice day Clare</h1>
          </div>

          <div className="main-acc main-acc--large main-loading main-loading--order3">
            <div className="main-acc__top">
              <h2 className="main-acc__name">{user?.accounts[0].type}</h2>
              <span className="main-acc__amount">{`฿${user?.accounts[0].accountBalances?.amount}`}</span>
              <span className="main-acc__detail main-acc__detail--num">Smart account {user?.accounts[0].accountNumber}</span>
              <span className="main-acc__detail">Powered by {user?.accounts[0].issuer}</span>
            </div>

            <div className="main-acc__bottom">
              <div className="main-acc__ani_box">
                <div className="main-acc__ani__item">
                  <span className="main-acc__ani_img1"></span>
                  <span className="main-acc__ani_img2"></span>
                </div>
                <div className="main-acc__ani__item2">
                  <span className="main-acc__ani_img3"></span>
                </div>
              </div>
              <div className="main-acc__link__box">
                <div className="main-acc__link__item">
                  <a href="#" className="main-acc__link main-acc__link--withdrawal">Withdrawal</a>
                  <a href="#" className="main-acc__link main-acc__link--qr">QR scan</a>
                  <a href="#" className="main-acc__link main-acc__link--addmoney">Add money</a>
                </div>
              </div>
            </div>
            <button type="button" className="main-acc__more">
              <span className="blind">More Action</span>
            </button>
            <div className="tooltip " style={{ display: 'none' }}>

              <button type="button" className="tooltip__btn-more">Set main account</button>

              <button type="button" className="tooltip__btn-more">Copy account number</button>

              <button type="button" className="tooltip__btn-more">Edit Name and Color</button>
            </div>
            <div className="tooltip tooltip--bubble tooltip--right-under" style={{ display: 'none' }}>
              <span className="tooltip__txt">Change your main account for <br />Using transfer, Wallet more
                easliy</span>
            </div>
          </div>
          {/* List of Transactions */}
          <div className="rctly__wrap main-loading main-loading--order5">
            <ul className="rctly__lst">
              {
                user?.transactions.map((transaction: Transaction) => (
                  <li key={transaction.transactionId} className="rctly__item">
                    <a href="#" className="rctly__link">
                      <span className="rctly__thumb"><img src={transaction.image || 'https://dummyimage.com/54x54/999/fff'}
                        alt="" /></span>
                      <span className="rctly__name">{transaction.name}</span>
                    </a>
                  </li>
                ))
              }
            </ul>
          </div>
          {/* Banner */}
          {
            user?.banners.map(banner => (
              <a className="main-make main-loading main-loading--order6" style={{ display: 'none' }}>
                <strong className="main-make__tit">{banner.title}</strong>
                <p className="main-make__dsc">{banner.description}</p>
              </a>
            ))
          }
          {/* Debit */}
          <div className="debit-swipe__wrap main-loading main-loading--order6">
            <div className="debit-swipe__inner">
              <div className="debit-swipe__lst" style={{ width: '1595px' }}>
                {
                  user?.debitCards.map(debit => (
                    <a href="#" className="debit-swipe__item"
                  style={{  backgroundColor: debit.debitCardDesign.color, borderColor: debit.debitCardDesign.borderColor } as React.CSSProperties}>
                      <strong className="debit-swipe__name">{debit.name}</strong>
                      {debit.debitCardStatus.status == "Active"
                        ? <span className="debit-swipe__etc debit-swipe__etc--active">
                          <span className="debit-swipe__etc__num">{debit.debitCardDetails.number}</span>
                        </span>
                        : <span className="debit-swipe__etc">{debit.debitCardStatus.status}</span>
                      }
                      <span className="debit-swipe__issue">Issued by {debit.debitCardDetails.issuer}</span>
                    </a>
                  ))
                }
                <a href="#" className="debit-swipe__item debit-swipe__item--all">
                  See all
                </a>
              </div>
            </div>
          </div>

          {/* accounts */}
          {
            user?.accounts.map(account => (
              <AccountMain account={account} />
            ))
          }
          
          {/* Banner */}
          {
            user?.banners.map(banner => (
              <a href="#" className="main-prod">
                <span className="main-prod__cms-ico"><img src={banner.image} alt="" /></span>
                <strong className="main-prod__tit">{banner.title}</strong>
                <p className="main-prod__dsc">{banner.description}</p>
              </a>
            ))
          }

          <div className="main-tb">
            <a href="#" className="link-to">Total Balance</a>
          </div>
        </div>
      </main>
    </>
  );
};

export default BankMain;