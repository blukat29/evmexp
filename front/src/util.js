function getAxiosError(err) {
  try {
    var msg = err.response.data.error;
    if (msg)
      return msg;
    else
      return err.message;
  } catch {
    return err.message;
  }
}

const extsep = '-';

const extid = {
  encodeAddr: function(net, addr) {
    return net + extsep + addr;
  },
  decodeAddr: function(extAddr) {
    var parts = extAddr.split(extsep);
    if (parts.length != 2) {
      return {};
    } else {
      return {network: parts[0], addr: parts[1]};
    }
  },

  encodeCodeID: function(net, codeID) {
    return net + extsep + codeID;
  },
  decodeCodeID: function(extCodeID) {
    var parts = extCodeID.split(extsep);
    if (parts.length != 2) {
      return {};
    } else {
      return {network: parts[0], codeID: parts[1]};
    }
  },

  encodeTxid: function(net, txid) {
    return net + extsep + txid;
  },
  decodeTxid: function(extTxid) {
    var parts = extTxid.split(extsep);
    if (parts.length != 2) {
      return {};
    } else {
      return {network: parts[0], txid: parts[1]};
    }
  },

  encodeId: function(net, id) {
    return net + extsep + id;
  },
  decodeId: function(extId) {
    var parts = extId.split(extsep);
    if (parts.length != 2) {
      return {};
    } else {
      return {network: parts[0], id: parts[1]};
    }
  },
};

module.exports = {
  getAxiosError,
  extsep,
  extid,
};
