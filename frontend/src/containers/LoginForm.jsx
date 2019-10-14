import React, { Component } from 'react';

class LoginScreen extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="font-weight-bold">TRIPIT</h1>
          <br></br>
          <form>
            <div class="form-group">
              <label for="username">Username</label>
              <input class="form-control" id="username" name="username" placeholder="Enter Username"></input>
            </div>
            <div class="form-group">
              <label for="password">Password</label>
              <input class="form-control" id="password" name="password" placeholder="Enter Password" type="password"></input>
            </div>
            <div class="form-group">
              <button type="submit" class="btn btn-primary" onClick={() => logIn()}>Log In</button>
            </div>
            <div class="form-group">
              <button type="submit" class="btn btn-success" onClick={createAccount}>Create Account</button>
            </div>
          </form>
        </header>
      </div>
    );
  }
}

function logIn() {
  //document.getElementById("result").textContent = document.getElementById("username").value
}

async function createAccount(e) {
  e.preventDefault();
  const d = new FormData(e.target);
  const data = 'user=' + d.get('username') + '&name=A&pass=' + d.get('password');

  let res = await fetch('http://18.219.140.44:3000/user/create', {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: body
  }).json();

  const auth_token = res.body.auth;
}

export default App;