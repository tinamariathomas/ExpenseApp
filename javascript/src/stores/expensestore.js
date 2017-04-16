import Reflux from 'reflux';
import Actions from '../actions/expense';
import request from 'superagent';

export default class ExpenseStore extends Reflux.Store
{

  constructor() {
    super();
    this.listenables = Actions;
    this.state = {expenses:[]};

  }
  onGetExpenses() {
    let that = this;
    request
      .get('/expenses')
      .end(function(err, res){
        that.setState({expenses: res.body})
      });
  }

}