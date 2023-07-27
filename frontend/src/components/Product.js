import React from "react";
import { Table, Button, Form } from "react-bootstrap";

class Product extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      id: 0,
      name: "",
      price: "",
      products: [],
    };
  }

  componentDidMount() {
    this.getProducts();
  }

  componentWillUnmount() {}

  createProduct = (product) => {
    fetch("http://localhost:8000/product", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(product),
    }).then((response) => {
      if (response.status === 201) {
        this.getProducts();
      } else {
        alert("Não foi possível cadastrar o produto");
      }
    });
  };

  getProduct = (id) => {
    fetch("http://localhost:8000/product/" + id, { method: "GET" })
      .then((response) => response.json())
      .then((product) => {
        this.setState({
          id: product.id,
          name: product.name,
          price: product.price,
        });
      });
  };

  getProducts = () => {
    fetch("http://localhost:8000/products")
      .then((response) => response.json())
      .then((dados) => {
        this.setState({ products: dados });
      });
  };

  updateProduct = (product) => {
    fetch("http://localhost:8000/product/" + product.id, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(product),
    }).then((response) => {
      if (response.ok) {
        this.getProducts();
      } else {
        alert("Não foi possível atualizar o produto");
      }
    });
  };

  deleteProducts = (id) => {
    fetch("http://localhost:8000/product/" + id, { method: "DELETE" }).then(
      (response) => {
        if (response.ok) {
          this.getProducts();
        }
      }
    );
  };

  renderTable() {
    return (
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Nome</th>
            <th>Preço</th>
            <th>Opções</th>
          </tr>
        </thead>
        <tbody>
          {this.state.products.map((product) => (
            <tr>
              <td>{product.name}</td>
              <td>{product.price}</td>
              <td>
                <Button
                  variant="secondary"
                  onClick={() => this.getProduct(product.id)}
                >
                  Atualizar
                </Button>
                <Button
                  variant="danger"
                  onClick={() => this.deleteProducts(product.id)}
                >
                  Excluir
                </Button>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    );
  }

  updateName = (e) => {
    this.setState({
      name: e.target.value,
    });
  };

  updatePrice = (e) => {
    const price = parseInt(e.target.value, 10);

    if (!isNaN(price)) {
      this.setState({
        price: price,
      });
    }
  };

  submit = () => {
    if (this.state.id === 0) {
      const product = {
        name: this.state.name,
        price: this.state.price,
      };
      this.createProduct(product);
    } else {
      const product = {
        id: this.state.id,
        name: this.state.name,
        price: this.state.price,
      };
      this.updateProduct(product);
    }
  };

  render() {
    return (
      <div>
        <Form>
          <Form.Group className="mb-3" controlId="formBasic">
            <Form.Label>Email address</Form.Label>
            <Form.Control
              type="name"
              placeholder="Enter name product"
              value={this.state.name}
              onChange={this.updateName}
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBasic">
            <Form.Label>Price</Form.Label>
            <Form.Control
              type="number"
              placeholder="Enter price product (only integers)"
              value={this.state.price}
              onChange={this.updatePrice}
            />
          </Form.Group>
          <Button variant="primary" type="submit" onClick={this.submit}>
            Submit
          </Button>
        </Form>

        {this.renderTable()}
      </div>
    );
  }
}

export default Product;
