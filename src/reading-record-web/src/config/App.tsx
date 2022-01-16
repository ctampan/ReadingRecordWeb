import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { BrowserRouter, Link } from "react-router-dom";
import AppRoutes from "./AppRoutes";

function App() {
  return (
    <BrowserRouter>
      <header className="navbar flex justify-end p-2">
        <div className="flex">
          <Link className="px-2 text-white font-bold" to="/">Home</Link>
          <Link className="px-2 text-white font-bold" to="/about">About</Link>
          <Link className="px-2 text-white font-bold"to="/dashboard">Dashboard</Link>
        </div>
      </header>
      <hr />
      <section>
        <AppRoutes />
      </section>
    </BrowserRouter>
  );
}

export default App;
