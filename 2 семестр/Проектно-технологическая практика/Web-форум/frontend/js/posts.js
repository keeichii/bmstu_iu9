/* eslint-disable */
// @ts-nocheck

document.addEventListener("DOMContentLoaded", () => {
  const postList = document.getElementById("post-list");
  const API_BASE = "http://localhost:8000";
  const token = localStorage.getItem("token");

  /* ───── Получение всех постов ───── */
  const fetchPosts = () => {
    postList.innerHTML = "<p>Загрузка постов...</p>";

    fetch(`${API_BASE}/posts/`, {
      headers: {
        "Content-Type": "application/json",
        ...(token && { Authorization: `Bearer ${token}` })
      }
    })
      .then(res => {
        if (!res.ok) throw new Error(`Ошибка: ${res.status}`);
        return res.json();
      })
      .then(posts => {
        postList.innerHTML = "";
        if (!posts.length) {
          postList.innerHTML = "<p>Постов пока нет.</p>";
          return;
        }
        posts.sort((a, b) => b.id - a.id);

        posts.forEach(post => {
          const node = document.createElement("div");
          node.classList.add("post");
          node.innerHTML = `
            <div class="post__author">Автор: ${post.author || "неизвестен"}</div>
            <div class="post__title">${post.title}</div>
            <div class="post__content">${post.content}</div>
            <div class="comments" id="comments-${post.id}">Загрузка комментариев...</div>
            ${
              token
                ? `<form class="comment-form" data-post-id="${post.id}" style="margin-top:10px;">
                     <input type="text" placeholder="Оставьте комментарий..." class="comment-input" required />
                     <button type="submit">Отправить</button>
                   </form>`
                : `<p style="color:#888;margin-top:10px;">Авторизуйтесь, чтобы комментировать</p>`
            }`;
          postList.append(node);

          // Комментарии
          fetch(`${API_BASE}/comments/${post.id}`)
            .then(r => r.json())
            .then(comments => {
              const box = document.getElementById(`comments-${post.id}`);
              box.innerHTML = comments.length
                ? comments.map(c => `<div class="comment"><strong>${c.author}</strong>: ${c.content}</div>`).join("")
                : "<p style='color:#aaa;'>Комментариев нет</p>";
            });
        });
      })
      .catch(err => {
        postList.innerHTML = `<p style="color:red">Ошибка: ${err.message}</p>`;
      });
  };

  fetchPosts();

  /* ───── Создание поста ───── */
  const postForm = document.getElementById("create-post-form");
  if (postForm) {
    postForm.addEventListener("submit", e => {
      e.preventDefault();
      if (!token) { alert("Вы не авторизованы"); return; }

      const title = document.getElementById("post-title").value;
      const content = document.getElementById("post-content").value;

      fetch(`${API_BASE}/posts/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`
        },
        body: JSON.stringify({ title, content })
      })
        .then(res => {
          if (!res.ok) throw new Error("Ошибка при создании поста");
        })
        .then(() => { postForm.reset(); fetchPosts(); })
        .catch(err => alert(err.message));
    });
  }

  /* ───── Добавление комментария ───── */
  document.addEventListener("submit", async e => {
    if (e.target.classList.contains("comment-form")) {
      e.preventDefault();
      const form = e.target;
      const postId = form.dataset.postId;
      const input = form.querySelector(".comment-input");
      const text = input.value.trim();
      if (!text) return;

      const res = await fetch(`${API_BASE}/comments/${postId}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`
        },
        body: JSON.stringify({ content: text })
      });

      if (res.ok) {
        const c = await res.json();
        const box = document.getElementById(`comments-${postId}`);
        box.innerHTML += `<div class="comment"><strong>${c.author}</strong>: ${c.content}</div>`;
        input.value = "";
      } else {
        alert("Ошибка при отправке комментария");
      }
    }
  });
});
