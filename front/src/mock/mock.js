import axios from 'axios';
//import queryString from 'query-string';
import MockAdapter from 'axios-mock-adapter';

import * as sampleContract from './eth_usdt.json';

function isDevServer() {
  return (typeof webpackHotUpdate == 'function');
}

function installMock() {
  if (!isDevServer()) {
    return;
  }
  console.log('Installing API mock');
  var mock = new MockAdapter(axios)

  mock.onGet("/api/deco/code0000").reply(200, sampleContract);
  //mock.onPost("/api/contract/upload").reply(function(req) { });
}

export default installMock;
