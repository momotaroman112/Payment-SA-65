import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import React, { Fragment, useEffect, useState } from "react";

import SignIn from "./components/SignIn";
import Home from "./components/Home";
import Navbar from "./components/Navbar";
import Bills from "./components/Bill";

import BillCreate from "./components/BillCreate";

function App() {
  const [token, setToken] = useState<string | null>();

  useEffect(() => {
    const getToken = localStorage.getItem("id");
    if (getToken) {
      setToken(getToken);
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  return (
    <Router>
      <div>
        {token && (
          <Fragment>
            <Navbar />
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/create" element={<BillCreate />} />
              <Route path="/history" element={<Bills />} />
            </Routes>
          </Fragment>
        )}
      </div>
    </Router>
  );
}

export default App;
