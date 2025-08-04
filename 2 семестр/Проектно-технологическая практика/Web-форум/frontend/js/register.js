document.getElementById("register-form").addEventListener("submit", async (e) => {
    e.preventDefault();
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
  
    const response = await fetch("http://127.0.0.1:8000/auth/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password })
    });
  
    if (response.ok) {
      alert("Регистрация успешна!");
      window.location.href = "login.html";
    } else {
      const error = await response.text();
      alert("Ошибка: " + error);
    }
  });
  