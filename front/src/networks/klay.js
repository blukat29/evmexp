
class KlayNetwork {
  constructor(name, displayName, baseUrl) {
    this.name = name || "klay";
    this.displayName = displayName || "Klaytn mainnet";
    this.baseUrl = baseUrl || "https://scope.klaytn.com";
  }
  isAddr(addr) {
    return !!addr.match(/^0x[a-fA-F0-9]{40}$/);
  }
  isTxid(txid) {
    return !!txid.match(/^0x[a-fA-F0-9]{64}$/);
  }
  explorerName() {
    return "Klaytnscope";
  }
  addrExplorer(addr) {
    if (this.isAddr(addr)) {
      return this.baseUrl + "/account/" + addr;
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

export default KlayNetwork;
