console.log("✅ user home.js loaded");

(async function () {
  const token = localStorage.getItem("token");
  console.log("📦 Token: ", token);

  if (!token) {
    alert("⛔ No token. Please login.");
    window.location.href = "/user/index.html";
    return;
  }

  const payload = JSON.parse(atob(token.split('.')[1])); 

  if (isTokenExpired(token)) {
    alert("⚠️ Session expired. Please login again.");
    localStorage.removeItem("token");
    window.location.href = "/user/index.html"; // or /post/index.html
    return;
  }
  const now = Math.floor(Date.now() / 1000);
  console.log("⏱️ Now:", now, "| Exp:", payload.exp);

  if (!payload.exp || now >= payload.exp) {
    alert("⚠️ Session expired. Please login again.");
    localStorage.removeItem("token");
    window.location.href = "/user/index.html";
    return;
  }

  try {
    const res = await fetch("http://localhost:8080/users", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      console.warn("❌ Invalid token on server:", res.status);
      localStorage.removeItem("token");
      alert("Token expired or unauthorized.");
      window.location.href = "/user/index.html";
      return;
    }

    const users = await res.json();
    const table = document.getElementById("userTable");
    if (!table) return;

    table.innerHTML = "<tr><th>ID</th><th>Name</th><th>Username</th><th>Email</th></tr>";
    users.forEach(user => {
      table.innerHTML += `<tr><td>${user.id}</td><td>${user.name}</td><td>${user.username}</td><td>${user.email}</td></tr>`;
    });

  } catch (err) {
    console.error("🚨 Error loading users:", err);
    alert("Unexpected error occurred.");
    window.location.href = "/user/index.html";
  }
})();

async function createUser() {
  const name = document.getElementById("userName").value;
  const username = document.getElementById("userUsername").value;
  const email = document.getElementById("userEmail").value;

  const token = localStorage.getItem("token");
  if (!token) {
    alert("❌ Session expired. Please log in again.");
    window.location.href = "/user/index.html";
    return;
  }

  try {
    const res = await fetch("http://localhost:8093/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ name, username, email }),
    });

    if (res.ok) {
      alert("✅ User created successfully!");
      fetchUsers();
      document.getElementById("userForm").reset();
    } else {
      const error = await res.text();
      alert("❌ Failed to create user:\n" + error);
    }
  } catch (err) {
    console.error("🚨 Error during user creation:", err);
    alert("Unexpected error occurred.");
  }
}
