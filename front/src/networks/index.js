import EthNetwork from './eth.js';

const networkMap = {
  'eth': EthNetwork,
};

const Networks = {
  parseExtendedId: function(extendedSomething) {
    var parts = extendedSomething.split('-');
    if (parts.length < 2)
      return null;
    if (!(parts[0] in networkMap))
      return null;

    var eid = {
      network: networkMap[parts[0]],
      name: parts[0],
    };
    if (parts.length == 2) {
      eid.subnet = null;
      eid.id = parts[1];
    } else if (parts.length == 3) {
      eid.subnet = parts[1];
      eid.id = parts[2];
    }
    return eid;
  },
  explorerName: function(extendedId) {
    var eid = this.parseExtendedId(extendedId);
    if (!eid)
      return null;
    return eid.network.explorerName();
  },
  addrExplorer: function(extendedAddr) {
    var eid = this.parseExtendedId(extendedAddr);
    if (!eid)
      return null;
    return eid.network.addrExplorer(eid.id, eid.subnet);
  },
};

export default Networks;
