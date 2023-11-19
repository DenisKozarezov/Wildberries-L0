import logo from './logo.svg';
import './App.css';
import { useState } from 'react';

const ordersEndpoint = "http://localhost:8080/api/orders"

function ordersSubmitHandler(orderUid) {
  fetch(`${ordersEndpoint}?order_uid=${orderUid}`, {
      mode: 'no-cors',
      method: "get",
      headers: {
         "Content-Type": "application/json; charset=utf-8"
      },
    })
    .then(async response => {
      const data = await response.json()

      if (!response.ok) {
        // get error message from body or default to response statusText
        const error = (data && data.message) || response.statusText;
        return Promise.reject(error);
      }
    })
    .catch(error => {
      console.error('There was an error!', error);
  });
}

function OrdersForm() {
  const [orderUID, setOrderUID] = useState('');
  return (
    <div>
      <input name="myInput" value={orderUID} placeholder="Введите ID заказа" onChange={e => setOrderUID(e.target.value)}/>
      <button onClick={() => ordersSubmitHandler(orderUID)}>
        Find
      </button>
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
