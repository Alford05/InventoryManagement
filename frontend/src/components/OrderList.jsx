import React, { useEffect, useState } from "react";
import { fetchOrders, updateOrderStatus, cancelOrder } from "../api/api";

const STATUS_OPTIONS = ["pending", "paid", "shipped", "canceled"];

export default function OrderList() {
  const [orders, setOrders] = useState([]);
  const [customerId, setCustomerId] = useState("");

  const loadOrders = async () => {
    const data = await fetchOrders(customerId || null);
    setOrders(data);
  };

  useEffect(() => {
    loadOrders();
  }, []);

  const handleStatusChange = async (orderId, newStatus) => {
    await updateOrderStatus(orderId, newStatus);
    loadOrders();
  };

  const handleCancel = async (orderId) => {
    if (!window.confirm("Cancel this order?")) return;
    await cancelOrder(orderId);
    loadOrders();
  };

  return (
    <div>
      <h2>Orders</h2>

      <div style={{ marginBottom: "1rem" }}>
        <input
          type="number"
          placeholder="Filter by Customer ID"
          value={customerId}
          onChange={(e) => setCustomerId(e.target.value)}
        />
        <button onClick={loadOrders}>Load</button>
      </div>

      {orders.map((order) => (
        <div
          key={order.id}
          style={{
            border: "1px solid #ccc",
            padding: "1rem",
            marginBottom: "1rem",
          }}
        >
          <h4>
            Order #{order.id} — Status:{" "}
            <strong>{order.status}</strong>
          </h4>

          <p>Customer ID: {order.customer_id}</p>

          <ul>
            {order.items.map((item) => (
              <li key={item.id}>
                Product #{item.product_id} — Qty: {item.quantity} — $
                {item.unit_price}
              </li>
            ))}
          </ul>

          <div style={{ marginTop: "0.5rem" }}>
            <select
              value={order.status}
              onChange={(e) =>
                handleStatusChange(order.id, e.target.value)
              }
              disabled={order.status === "canceled"}
            >
              {STATUS_OPTIONS.map((s) => (
                <option key={s} value={s}>
                  {s}
                </option>
              ))}
            </select>

            <button
              style={{ marginLeft: "1rem" }}
              onClick={() => handleCancel(order.id)}
              disabled={
                order.status === "shipped" ||
                order.status === "canceled"
              }
            >
              Cancel Order
            </button>
          </div>
        </div>
      ))}
    </div>
  );
}
