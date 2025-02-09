import * as React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import BankMain from "./pages/BankMain/BankMain";
import Pin from "./pages/Pin/Pin";
import Splash from "./pages/Splash/Splash";
import "./css/th-bank.css";
const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Splash />} />
        <Route path="/splash" element={<Splash />} />
        <Route path="/bank-main" element={<BankMain />} />
        <Route path="/pin" element={<Pin />} />
      </Routes>
    </Router>
  );
};

export default App;