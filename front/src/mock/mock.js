import crypto from 'crypto';

import axios from 'axios';
//import queryString from 'query-string';
import MockAdapter from 'axios-mock-adapter';

import * as sampleContract from './eth_usdt.json';

function isDevServer() {
  return (typeof webpackHotUpdate == 'function');
}

function sha256(data) {
  return crypto.createHash('sha256').update(data).digest().toString('hex')
}

function delayed(msec, code, res) {
  return new Promise((resolve) => {
    setTimeout(() => { resolve([code, res]); }, msec);
  });
}

function installMock() {
  if (!isDevServer()) {
    return;
  }
  console.log('Installing API mock');
  var mock = new MockAdapter(axios)

  axios.interceptors.request.use(request => {
    console.log({
      method: request.method,
      url: request.url,
      data: request.data,
    });
    return request
  });

  var db = {
    codes: {},
  };

  mock.onGet("/api/addr/eth-mainnet-0xdac17f958d2ee523a2206206994597c13d831ec7")
    .reply(() => delayed(1000, 200, {
      error: null,
      extendedCodeHash: "evm-generic-6d967f98f2f3843065688dc2065248e3686b56fc0b6ddfa82007df016148becb",
    }));

  mock.onGet(/\/api\/addr\/\w+/)
    .reply(() => delayed(0, 404, {
      error: "No such contract",
    }));

  mock.onGet("/api/deco/evm-generic-6d967f98f2f3843065688dc2065248e3686b56fc0b6ddfa82007df016148becb")
    .reply(() => delayed(2000, 500, sampleContract.default));

  mock.onPost("/api/code/upload")
    .reply((req) => {
      var data = JSON.parse(req.data);

      var format = data.format || "evm-generic";
      var binary = data.binary || "";
      binary = binary.trim();
      binary = Buffer.from(binary, 'hex');

      var codeHash = sha256(binary);
      if (binary.length == 0) {
        codeHash = "6d967f98f2f3843065688dc2065248e3686b56fc0b6ddfa82007df016148becb";
      }
      var extendedCodeHash = format + "-" + codeHash;

      if (db.codes[extendedCodeHash]) {
        // cached
        return delayed(100, 200, {
          error: null,
          extendedCodeHash: extendedCodeHash,
        });
      } else {
        // not cached
        db.codes[extendedCodeHash] = {
          format: format,
          binary: binary,
          extendedCodeHash: extendedCodeHash,
        };
        return delayed(3000, 200, {
          error: null,
          extendedCodeHash: extendedCodeHash,
        });
      }
    });
}

export default installMock;
