import EthNetwork from './eth.js';
import KlayNetwork from './klay.js';
import {extid} from '../util.js';

const allNetworks = [
  new EthNetwork(),
  new EthNetwork('eth_ropsten', "Ethereum Ropsten testnet", "https://ropsten.etherscan.io"),
  new KlayNetwork(),
];

const networkMap = new Map();
allNetworks.forEach(n => {
  networkMap[n.name] = n;
});

const Networks = {
  all: function() {
    return allNetworks;
  },
  get: function(name) {
    return networkMap[name];
  },
  explorerName: function(extId) {
    var eid = extid.decodeId(extId)
    if (!eid)
      return null;
    var net = Networks.get(eid.network)
    return net.explorerName();
  },
  addrExplorer: function(extAddr) {
    var eid = extid.decodeAddr(extAddr)
    if (!eid)
      return null;
    var net = Networks.get(eid.network)
    return net.addrExplorer(eid.addr);
  },
};

export default Networks;
