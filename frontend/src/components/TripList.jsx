import React, { Component } from 'react';
import { connect } from 'react-redux';

class TripList extends Component {
  renderTrips() {
    return this.props.trips.map(trip => {
      let d = new Date();
      return <li className='list-group-item' key={d}>{trip} <span className='text-muted'>{d.toDateString()}</span></li>;
    });
  }
  
  render() { //passed a list of comments
    console.log(this.props.trips);

    if (this.props.trips.length === 0) {
      return (
        <div id='trip-list' className='d-flex col-6 flex-column'>
          <h2 className='text-center'>List of Trips</h2>
          <p className='text-center'>No trips added yet!</p>
        </div>
      );
    }

    return (
      <div id='trip-list' className='d-flex col-6 flex-column'>
        <h2 className='text-center'>List of Trips</h2>
        <ul className='list-group w-100'>
          {this.renderTrips()}
        </ul>
      </div>
    );
  }
}

// const TripList = ( props, onTripClick ) => {
//   return (
//     <div>
//       {props.trips && <ul id='tripList' className='container'>
//         {props.trips.map((trip, index) => (
//           <Trip key={index}
//                 onClick={() => onTripClick(index)}
//                 starred={trip.starred}
//                 trip_name={trip.trip_name} />
//         ))}
//       </ul>}
//     </div>
//   )
// }

function mapStateToProps(state) {
  return { trips: state.trips };
}

export default connect(mapStateToProps)(TripList);