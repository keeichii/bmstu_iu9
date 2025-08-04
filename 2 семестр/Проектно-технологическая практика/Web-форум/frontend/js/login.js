document.getElementById("login-form").addEventListener("submit", async (e) => {
    e.preventDefault();
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
  
    const response = await fetch("http://127.0.0.1:8000/auth/login", {
      method: "POST",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: `username=${encodeURIComponent(username)}&password=${encodeURIComponent(password)}`
    });
  
    if (response.ok) {
      const data = await response.json();
      localStorage.setItem("token", data.access_token);
      alert("Вход выполнен");
      window.location.href = "index.html";
    } else {
      alert("Ошибка входа");
    }
  });
  