import React from "react";
import ProductList from "./components/ProductList";
import OrderForm from "./components/OrderForm";
import OrderList from "./components/OrderList";

function App() {
  return (
    <div>
      <h1>Inventory Management</h1>
      <ProductList />
      <hr />
      <OrderForm />
      <hr />
      <OrderList />
    </div>
  );
}

export default App;
