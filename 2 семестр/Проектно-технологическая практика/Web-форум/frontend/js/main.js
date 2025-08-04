document.addEventListener("DOMContentLoaded", () => {
  const logoutBtn        = document.getElementById("logout-btn");
  const loginLink        = document.getElementById("login-link");
  const registerLink     = document.getElementById("register-link");
  const createPostSection= document.getElementById("create-post-section");
  const userInfo         = document.getElementById("user-info");
  const usernameDisplay  = document.getElementById("username-display");
  const userPanel        = document.getElementById("user-panel");

  const postForm   = document.getElementById("create-post-form");
  const toggleBtn  = document.getElementById("toggle-post-form");
  const toggleTitle= document.querySelector(".toggle-title");

  const token = localStorage.getItem("token");

  /* ───── визуальное состояние навбара ───── */
  if (token) {
    logoutBtn.style.display    = "inline-block";
    loginLink.style.display    = "none";
    registerLink.style.display = "none";
    if (createPostSection) createPostSection.style.display = "block";
    if (userPanel)         userPanel.style.display = "block";

    try {
      const payload = JSON.parse(atob(token.split(".")[1]));
      if (usernameDisplay) {
        usernameDisplay.textContent = "Пользователь: " + payload.sub;
        userInfo.style.display = "block";
      }
    } catch (e) {
      console.error("Ошибка при чтении токена:", e);
    }
  } else {
    logoutBtn.style.display    = "none";
    loginLink.style.display    = "inline-block";
    registerLink.style.display = "inline-block";
    if (createPostSection) createPostSection.style.display = "none";
    if (userPanel)         userPanel.style.display = "none";
  }

  logoutBtn.addEventListener("click", () => {
    localStorage.removeItem("token");
    alert("Вы вышли из аккаунта");
    window.location.reload();
  });

  /* ───── сворачивание / разворачивание формы поста ───── */
  if (toggleBtn && postForm && toggleTitle) {
    let isCollapsed = false;

    const togglePostForm = () => {
      isCollapsed = !isCollapsed;
      postForm.classList.toggle("hidden", isCollapsed);
      toggleBtn.textContent = isCollapsed ? "+" : "–";
    };

    toggleBtn.addEventListener("click", togglePostForm);
    toggleTitle.addEventListener("click", togglePostForm);
  }
});
