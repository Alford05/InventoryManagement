import React, { useState } from "react";
import { createOrder } from "../api/api";

export default function OrderForm() {
  const [customerId, setCustomerId] = useState("");
  const [items, setItems] = useState([{ product_id: "", quantity: 1 }]);

  const handleAddItem = () => setItems([...items, { product_id: "", quantity: 1 }]);
  const handleChangeItem = (index, field, value) => {
    const newItems = [...items];
    newItems[index][field] = value;
    setItems(newItems);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const order = await createOrder({ customer_id: parseInt(customerId), items });
    console.log("Order created:", order);
    alert(`Order #${order.id} created!`);
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Customer ID: 
        <input type="number" value={customerId} onChange={e => setCustomerId(e.target.value)} required />
      </label>

      {items.map((item, idx) => (
        <div key={idx}>
          <input
            type="number"
            placeholder="Product ID"
            value={item.product_id}
            onChange={e => handleChangeItem(idx, "product_id", parseInt(e.target.value))}
            required
          />
          <input
            type="number"
            placeholder="Quantity"
            value={item.quantity}
            onChange={e => handleChangeItem(idx, "quantity", parseInt(e.target.value))}
            min="1"
            required
          />
        </div>
      ))}

      <button type="button" onClick={handleAddItem}>Add Another Item</button>
      <button type="submit">Create Order</button>
    </form>
  );
}
