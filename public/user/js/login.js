function isTokenExpired(token) {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    const now = Math.floor(Date.now() / 1000);

    console.log("🧾 JWT payload:", payload);
    console.log("⏰ Current time:", now);
    console.log("⌛ Expiry time:", payload.exp);
    console.log("⌛ Expires in:", payload.exp - now, "seconds");
    fmt.Println("🕒 Server time:", time.Now().Unix())
    console.log("🕒 Server time:", time.Now().Unix())

    return !payload.exp || now >= payload.exp;
  } catch (err) {
    console.error("❌ JWT parse error:", err);
    return true;
  }
}


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
