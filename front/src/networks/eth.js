
const etherscanBaseUrl = {
  "mainnet": "https://etherscan.io",
  "ropsten": "https://ropsten.etherscan.io",
  "kovan":   "https://kovan.etherscan.io",
  "rinkeby": "https://rinkeby.etherscan.io",
};

const EthNetwork = {
  isAddr: function(addr) {
    return !!addr.match(/^0x[a-fA-F0-9]{40}$/);
  },
  isTxid: function(txid) {
    return !!txid.match(/^0x[a-fA-F0-9]{64}$/);
  },
  explorerName: function() {
    return "etherscan.io";
  },
  addrExplorer: function(addr, subnet) {
    subnet = subnet || "mainnet";
    if (this.isAddr(addr)) {
      var base = etherscanBaseUrl[subnet];
      if (base) {
        return base + "/address/" + addr;
      }
    }
    return null;
  },
  txExplorer: function(txid, subnet) {
    subnet = subnet || "mainnet";
    if (this.isAddr(txid)) {
      var base = etherscanBaseUrl[subnet];
      if (base) {
        return base + "/tx/" + txid;
      }
    }
    return null;
  },
};

export default EthNetwork;
