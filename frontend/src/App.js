import "./App.css";
import Product from "./components/Product";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  return (
    <div className="App">
      <h1>Go-Products</h1>
      <Product></Product>
    </div>
  );
}

export default App;
