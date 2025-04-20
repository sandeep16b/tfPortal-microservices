window.onload = async function () {
    const token = localStorage.getItem("token");
  
    if (!token) {
      alert("Please log in first");
      window.location.href = "/post/index.html";
      return;
    }
  
    const res = await fetch("http://localhost:8080/posts", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  
    if (!res.ok) {
      alert("Unauthorized or session expired");
      window.location.href = "/post/index.html";
      return;
    }
  
    const users = await res.json();
    const list = document.getElementById("userList");
    users.forEach(u => {
      const item = document.createElement("li");
      item.innerText = `${u.id}: ${u.name} (${u.email})`;
      list.appendChild(item);
    });
  
    document.getElementById("welcomeMsg").innerText = "✅ Logged in!";
  };
  