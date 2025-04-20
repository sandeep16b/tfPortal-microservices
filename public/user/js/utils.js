function isTokenExpired(token) {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    const now = Math.floor(Date.now() / 1000);
    return !payload.exp || now >= payload.exp;
  } catch {
    return true; // treat as expired if broken
  }
}
