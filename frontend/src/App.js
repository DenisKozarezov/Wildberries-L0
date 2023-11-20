import logo from './logo.svg';
import './App.css';
import { useState } from 'react';

const ordersEndpoint = "http://localhost:8080/api/orders"

function OrdersForm() {
  const[data, setData] = useState('')
  const[inputData, setInputData] = useState('');

  const handleInputChange = (event) => {
    setInputData(event.target.value);
  }
  
  function handleSubmit() {
    fetch(`${ordersEndpoint}?order_uid=${inputData}`, {
      method: "get",
      headers: {
        "Content-Type": "application/json; charset=utf-8",
        "Access-Control-Allow-Origin": "http://localhost:3000",
        "Access-Control-Allow-Credentials": "true",
        "Access-Control-Allow-Methods": "GET",
        "Access-Control-Allow-Headers": "Content-Type, access-control-allow-origin, access-control-allow-headers, access-control-allow-methods, access-control-allow-credentials",
        Accept: "application/json",
      },
      crossorigin: true,    
    })
    .then(async response => await response.json())
    .then(jsonData => setData(JSON.stringify(jsonData, undefined, 2)))
    .catch(error => {
      console.error(error);
    });
  }

  return (
    <div>
      <input name="myInput" value={inputData} placeholder="Введите ID заказа" onChange={handleInputChange}/>
      <button onClick={handleSubmit}>
        Найти
      </button>
      <div>
        <strong>Ваш заказ</strong>
        <p>{data}</p>
      </div>
    </div>
  );
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Wildberries Internship
        </a>
          (Level 0)
        <OrdersForm />
      </header>
    </div>
  );
}

export default App;
