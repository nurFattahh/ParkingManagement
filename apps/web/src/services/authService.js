export async function login(username, password) {
  const response = await fetch("http://localhost:8080/api/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: username,
      password: password,
    }),
  });

  if (!response.ok) {
    throw new Error("Login gagal");
  }

  return response.json();
}

export async function register(username, password, full_name, phone, email) {
  const response = await fetch("http://localhost:8080/api/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: username,
      password: password,
      full_name: full_name,
      phone: phone,
      email: email,

    }),
  });

  if (!response.ok) {
    throw new Error("Registrasi gagal");
  }

  return response.json();
}