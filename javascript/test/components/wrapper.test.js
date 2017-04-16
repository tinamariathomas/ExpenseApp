import {expect} from 'chai';
import { shallow, render, mount } from 'enzyme';
import React from 'react';
import sinon from 'sinon';

import '../../test/initdom';

import ExpenseStore from '../../src/stores/expensestore';
import Wrapper from '../../src/components/wrapper';
import Actions from '../../src/actions/expense';

describe('Wrapper', function() {
  let wrapper;


  it('should listen to expense store', ()=>{
    wrapper = mount(<Wrapper />);

    expect(wrapper.node.stores.length).to.equal(1);
    expect(wrapper.node.stores[0].constructor.name).to.equal("ExpenseStore")
  });

  it('should call getExpenses on mount', () =>{
    let stub = sinon.stub(Actions, "getExpenses");

    wrapper = mount(<Wrapper />);

    expect(stub.calledOnce).to.be.true;
    stub.restore();
  });
});