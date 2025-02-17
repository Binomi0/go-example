import React from "react";
import { createRoot } from "react-dom/client";
import App from "./App";
import LoginForm from "./components/LoginForm";

const app = document.getElementById("root");
const login = document.getElementById("login");

if (!app && !login) throw new Error("No root element found");
// Render your React component instead
if (app) {
  const root = createRoot(app);
  root.render(<App />);
} else if (login) {
  const root = createRoot(login);
  root.render(<LoginForm />);
}
