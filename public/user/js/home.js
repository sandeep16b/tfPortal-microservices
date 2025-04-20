window.onload = async function () {
  const token = localStorage.getItem("token");
  if (!token) {
    // ⛔ No token means not logged in → go to login page
    alert("Please log in first.");
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
      // ⛔ Token is invalid or expired
      alert("Unauthorized or session expired. Please log in again.");
      localStorage.removeItem("token"); // Clear invalid token
      window.location.href = "/user/index.html";
      return;
    }
    // ✅ Load user list
    const users = await res.json();
    const table = document.getElementById("userTable");
    table.innerHTML = "<tr><th>ID</th><th>Name</th><th>Username</th><th>Email</th></tr>";
    users.forEach(user => {
      table.innerHTML += `<tr><td>${user.id}</td><td>${user.name}</td><td>${user.username}</td><td>${user.email}</td></tr>`;
    });
  } catch (err) {
    console.error("Error loading users:", err);
    alert("Something went wrong. Try again later.");
    window.location.href = "/user/index.html";
  }
};