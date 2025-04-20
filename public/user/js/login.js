window.onload = function () {
  const token = localStorage.getItem("token");
  if (token) {
    window.location.href = "/user/home.html";
  }
};

// login.js
async function login(event) {
  event.preventDefault();

  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

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
    alert("❌ Login failed: " + (data.error || "Invalid credentials"));
  }
}
