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

module.exports = {
  getAxiosError,
};
