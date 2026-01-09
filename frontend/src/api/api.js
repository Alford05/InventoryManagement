const API_BASE = "http://localhost:8080";

export async function fetchProducts() {
  const res = await fetch(`${API_BASE}/products`);
  return res.json();
}

export async function createProduct(data) {
  const res = await fetch(`${API_BASE}/products`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  return res.json();
}

export async function fetchOrders(customerId) {
  let url = `${API_BASE}/orders`;
  if (customerId) url += `?customer_id=${customerId}`;
  const res = await fetch(url);
  return res.json();
}

export async function createOrder(data) {
  const res = await fetch(`${API_BASE}/orders`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  return res.json();
}


