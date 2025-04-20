// On successful login response
fetch("http://localhost:8080/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username: "admin", password: "password123" }),
  })
    .then(res => res.json())
    .then(data => {
      if (data.token) {
        localStorage.setItem("token", data.token); // ✅ Store in localStorage
        window.location.href = "/user/index.html";       // Redirect to home
      } else {
        alert("Login failed");
      }
    });
  