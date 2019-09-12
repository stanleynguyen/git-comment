module.exports = {
  handleGenericHTTPErr: function(err, res) {
    if (err.send) return err.send(res);
    console.trace(err); // eslint-disable-line no-console
    res.status(500).json({ message: 'Internal Server Error' });
  },
};
