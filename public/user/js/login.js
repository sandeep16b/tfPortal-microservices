

window.onload = function () {
  const token = localStorage.getItem("token");
  console.log("📦 Token:", token);

  if (token && !isTokenExpired(token)) {
    window.location.href = "/user/home.html";
  } else {
    localStorage.removeItem("token");
    //localStorage.clear(); // 🧹 Clear expired or invalid token
  }
};

async function login(event) {
  event.preventDefault();

  const username = document.getElementById("username").value.trim();
  const password = document.getElementById("password").value.trim();

  if (!username) {
    alert("⚠️ Please enter a username");
    return;
  }

  if (!password) {
    alert("⚠️ Please enter a password");
    return;
  }

  try {
    const res = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password }),
    });

    const data = await res.json();

    if (res.ok && data.token) {
      localStorage.setItem("token", data.token);
      window.location.href = "/user/home.html";
    } else {
      alert("❌ " + (data.error || "Invalid username or password"));
    }
  } catch (err) {
    console.error("Login error:", err);
    alert("❌ Server error. Please try again.");
  }
}
