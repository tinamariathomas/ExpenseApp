import nock from 'nock';
import sinon from 'sinon';
import {expect} from 'chai';

import ExpenseStore from '../../src/stores/expensestore';

describe('Expense store', function() {
    let expenseStore;
    beforeEach(()=>{
      expenseStore = new ExpenseStore();
    });

    it('should make call to getch expenses', function(done) {
      const expenses = [
        {"Id":15,"Description":"Shoes","Amount":1000},
        {"Id":16,"Description":"Shoes","Amount":1000},
      ];
      const stub = sinon.stub(expenseStore,"setState").callsFake((data)=>{
        expect(data.expenses).to.deep.equal(expenses);
        done();
      });
      const request = nock('http://localhost')
        .get('/expenses')
        .reply(200, expenses);

      expenseStore.onGetExpenses();
    });

    it('state should be initialized', ()=>{
      expect(expenseStore.state.expenses).to.be.empty;
    })
});