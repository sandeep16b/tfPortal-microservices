console.log("✅ post home.js loaded");

(async function () {
  const token = localStorage.getItem("token");
  console.log("📦 Token: ", token);

  if (!token) {
    alert("⛔ No token. Please login.");
    window.location.href = "/post/index.html";
    return;
  }

  let payload;
  try {
    payload = JSON.parse(atob(token.split('.')[1]));
  } catch (err) {
    console.error("❌ Failed to decode token:", err);
    localStorage.removeItem("token");
    window.location.href = "/post/index.html";
    return;
  }

  const now = Math.floor(Date.now() / 1000);
  console.log("⏱️ Now:", now, "| Exp:", payload.exp);

  if (!payload.exp || now >= payload.exp) {
    alert("⚠️ Session expired. Please login again.");
    localStorage.removeItem("token");
    window.location.href = "/post/index.html";
    return;
  }

  // ✅ Main logic goes here — optionally uncomment if server endpoint is ready
  try {
    const res = await fetch("http://localhost:8080/posts", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  
    if (!res.ok) {
      console.warn("❌ Invalid token on server:", res.status);
      localStorage.removeItem("token");
      alert("Token expired or unauthorized.");
      window.location.href = "/post/index.html";
      return;
    }

    const posts = await res.json();
    const table = document.getElementById("postTable");
    if (!table) return;

    table.innerHTML = "<tr><th>ID</th><th>UserID</th><th>Title</th><th>Body</th></tr>";
    posts.forEach(post => {
      table.innerHTML += `<tr><td>${post.id}</td><td>${post.userId}</td><td>${post.title}</td><td>${post.body}</td></tr>`;
    });

  } catch (err) {
    console.error("🚨 Error loading posts:", err);
    alert("Unexpected error occurred.");
    window.location.href = "/post/index.html";
  }

})();

// 🔄 Create post remains outside the IIFE
async function createPost() {
  const userId = parseInt(document.getElementById("postUserId").value);
  const title = document.getElementById("postTitle").value;
  const body = document.getElementById("postBody").value;

  const token = localStorage.getItem("token");
  if (!token) {
    alert("❌ Session expired. Please log in again.");
    window.location.href = "/post/index.html";
    return;
  }

  try {
    const res = await fetch("http://localhost:8080/posts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ userId, title, body }),
    });

    if (res.ok) {
      alert("✅ Post created successfully!");
      location.reload(); // 🔄 Reload page to refresh posts if needed
    } else {
      const error = await res.text();
      alert("❌ Failed to create post:\n" + error);
    }
  } catch (err) {
    console.error("🚨 Error creating post:", err);
    alert("Unexpected error occurred.");
  }
}
