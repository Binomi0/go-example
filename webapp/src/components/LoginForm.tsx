import React, { useRef, FormEvent } from "react";

const LoginForm = () => {
  const usernameRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault(); // Evitar que el formulario se envíe de manera tradicional

    var username = usernameRef.current?.value;
    var password = passwordRef.current?.value;

    // Validar datos
    if (!username || !password) {
      console.error("Username and password are required.");
      return; // Salir si los campos están vacíos
    }

    fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Server response:", data);
        window.location.href = "/";
      })
      .catch((error) => {
        console.error("Error sending request:", error);
      });
  };

  return (
    <div>
      <h1>Welcome to the Login Page v2</h1>
      <form id="myForm" onSubmit={handleSubmit}>
        <label htmlFor="username">Username:</label>
        <input
          autoComplete="webauthn-username"
          ref={usernameRef}
          type="text"
          id="username"
          name="username"
        />
        <br />
        <br />
        <label htmlFor="password">Password:</label>
        <input
          autoComplete="webauthn-password"
          ref={passwordRef}
          type="password"
          id="password"
          name="password"
        />
        <br />
        <br />
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default LoginForm;
