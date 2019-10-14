import React, { Component } from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions'; //index.js

class TripForm extends Component {
  state = {
    trip: ''
  };

  handleChange = (e) => {
    this.setState({ trip: e.target.value })
  };

  handleSubmit = (e) => {
    e.preventDefault();

    this.props.addTrip(this.state.trip);

    this.setState({ trip: '' })
  };

  render() {
    return (
      <div className='col-4'>
        <form onSubmit={this.handleSubmit} className='d-flex m-4 flex-column align-items-center'>
          <input className='form-control' type='text' onChange={this.handleChange} value={this.state.trip} />
          <button className='btn btn-outline-light m-2' type='submit'>Add Trip</button>
        </form>
      </div>
    );
  }
}

export default connect(null, actions)(TripForm);