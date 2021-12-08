
class EthNetwork {
  constructor(name, displayName, baseUrl) {
    this.name = name || "eth";
    this.displayName = displayName || "Ethereum mainnet";
    this.baseUrl = baseUrl || "https://etherscan.io";
  }
  isAddr(addr) {
    return !!addr.match(/^0x[a-fA-F0-9]{40}$/);
  }
  isTxid(txid) {
    return !!txid.match(/^0x[a-fA-F0-9]{64}$/);
  }
  explorerName() {
    return "etherscan.io";
  }
  addrExplorer(addr) {
    if (this.isAddr(addr)) {
      return this.baseUrl + "/address/" + addr;
    } else {
      return null;
    }
  }
  txExplorer(txid) {
    if (this.isAddr(txid)) {
      return this.baseUrl + "/tx/" + txid;
    } else {
      return null;
    }
  }
}

export default EthNetwork;
