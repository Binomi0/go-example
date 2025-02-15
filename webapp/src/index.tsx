import React from "react";
import { createRoot } from "react-dom/client";
import App from "./App";

const app = document.getElementById("root");
if (!app) throw new Error("No root element found");
// Render your React component instead
const root = createRoot(app);
root.render(<App />);
