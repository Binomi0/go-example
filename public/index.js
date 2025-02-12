document.addEventListener("DOMContentLoaded", () => {
  console.log("script loaded");
  // Agregar listener al formulario para el evento submit
  document.getElementById("myForm").addEventListener("submit", function (e) {
    e.preventDefault(); // Evitar que el formulario se envíe de manera tradicional

    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    // Validar datos
    if (!username || !password) {
      console.error("Username and password are required.");
      return; // Salir si los campos están vacíos
    }

    // Aquí puedes enviar los datos al servidor usando AJAX o formularios de JavaScript
    console.log("Username:", username);
    console.log("Password:", password);

    // Ejemplo de envío con XMLHttpRequest
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "login", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        console.log("Response from server:", xhr.responseText);
      }
    };
    xhr.send(
      "username=" +
        encodeURIComponent(username) +
        "&password=" +
        encodeURIComponent(password)
    );
  });
});
