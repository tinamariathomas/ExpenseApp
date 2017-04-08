import React from 'react';
import ReactDOM from 'react-dom';


class App extends React.Component {
  render () {
    return <p> But so is React!</p>;
  }
}

ReactDOM.render(<App/>, document.getElementById('root'));