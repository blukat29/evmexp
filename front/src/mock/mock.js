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

  mock.onGet("/api/deco/evm-generic-6d967f98f2f3843065688dc2065248e3686b56fc0b6ddfa82007df016148becb").reply(200, sampleContract);
  //mock.onPost("/api/contract/upload").reply(function(req) { });
}

export default installMock;
