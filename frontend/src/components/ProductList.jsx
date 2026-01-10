import React, { useEffect, useState } from "react";
import { fetchProducts } from "../api/api";

export default function ProductList() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetchProducts().then(setProducts);
  }, []);

  return (
    <div>
      <h2>Products</h2>
      <table>
        <thead>
          <tr>
            <th>Name</th><th>Price</th><th>Stock</th><th>Category</th>
          </tr>
        </thead>
        <tbody>
          {products.map(p => (
            <tr key={p.id}>
              <td>{p.name}</td>
              <td>${p.price}</td>
              <td>{p.stock}</td>
              <td>{p.category_id}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
