import React from 'react';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1 className="font-weight-bold">TRIPIT</h1>
        <br></br>
        <form>
          <div class="form-group">
            <label for="username">Username</label>
            <input class="form-control" id="username" placeholder="Enter Username"></input>
          </div>
          <div class="form-group">
            <label for="password">Password</label>
            <input class="form-control" id="password" placeholder="Enter Password" type="password"></input>
          </div>
          <div class="form-group">
            <button type="submit" class="btn btn-primary" onClick={() => logIn()}>Log In</button>
          </div>
          <div class="form-group">
            <button type="submit" class="btn btn-success" onClick={() => createAccount()}>Create Account</button>
          </div>
        </form>
      </header>
    </div>
  );
}

function logIn() {
  //document.getElementById("result").textContent = document.getElementById("username").value
}

function createAccount() {
  fetch('http://localhost:8080/user/create', {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: "user=jonah&name=jonah&pass=wu"
  }).then((response) => response.json())
}

export default App;
