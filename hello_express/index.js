var express = require('express');
// var logger = require('morgan');

var app = express();
// app.use(logger('dev'));

app.get('/', function (req, res) {
  res.status(200).json({
    hello: 'world'
  });
});

app.listen(3000, function () {
  console.log('Listening on port 3000!')
});

module.exports = app;
