import React from "react";
import { createRoot } from "react-dom/client";
import App from "./App";
var app = document.getElementById("root");
if (!app)
    throw new Error("No root element found");
// Render your React component instead
var root = createRoot(app);
root.render(React.createElement(App, null));
//# sourceMappingURL=index.js.map