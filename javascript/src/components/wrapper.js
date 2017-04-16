import React from 'react';
import Reflux from 'reflux';

import ExpenseStore from '../stores/expensestore';
import Actions from '../actions/expense';


export default class Wrapper extends Reflux.Component {

  constructor(props)
  {
    super(props);
    this.stores = [ExpenseStore];
  }

  componentDidMount(){
    Actions.getExpenses();
  }

  render () {
    return(<div>Hi</div>);
  }
}