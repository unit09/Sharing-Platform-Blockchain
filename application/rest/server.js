const express = require('express');
const app = express();

var path = require('path');
var sdk = require('./sdk');

const PORT = 8080;
const HOST = 'localhost';

// common
app.get('/api/getAllShareRecord', function (req, res) {
    let args = [];
    sdk.send(false, 'getAllShareRecord', args, res, req.query.ccp);
});

app.get('/api/getShareStart', function (req, res) {
    let args = [];
    sdk.send(false, 'getShareStart', args, res, req.query.ccp);
});

app.get('/api/getShareEnd', function (req, res) {
    let args = [];
    sdk.send(false, 'getShareEnd', args, res, req.query.ccp);
});

app.get('/api/getShareRecordByLocation', function (req, res) {
    var location = req.query.location;

    let args = [location];

    sdk.send(false, 'getShareRecordByLocation', args, res, req.query.ccp);
});

app.get('/api/countAllShareRecordByLocation', function (req, res) {
    let args = [];
    sdk.send(false, 'countAllShareRecordByLocation', args, res, req.query.ccp);
});

app.get('/api/getAllPlace', function (req, res) {
    let args = [];
    sdk.send(false, 'getAllPlace', args, res, req.query.ccp);
});


// share1
app.get('/api/startShare', function (req, res) {
    var id = req.query.id;
    var target = req.query.target;
    var location = req.query.location;

    let args = [id, target, location];

    sdk.send(true, 'startShare', args, res, req.query.ccp);
});

app.get('/api/endShare', function (req, res) {
    var id = req.query.id;
    var target = req.query.target;
    var location = req.query.location;

    let args = [id, target, location];

    sdk.send(true, 'endShare', args, res, req.query.ccp);
});

// agency
app.get('/api/setPlace', function (req, res) {
    var id = req.query.id;
    var location = req.query.location;

    let args = [id, location];

    sdk.send(true, 'setPlace', args, res, req.query.ccp);
});

// monitor
app.get('/api/getUserShareRecord', function (req, res) {
    var id_set = String(req.query.id_set);
    console.log(id_set);

    let args = [];

    args = id_set.split(",");
    console.log(args);

    sdk.send(false, 'getUserShareRecord', args, res, req.query.ccp);
});
/*
app.get('/api/setMusic', function (req, res) {
    var title = req.query.title;
    var singer = req.query.singer;
    var price = req.query.price;
    var walletid = req.query.walletid;

    let args = [title, singer, price, walletid];
    sdk.send(true, 'setMusic', args, res);
});
app.get('/api/getAllmusic', function (req, res) {
    let args = [];
    sdk.send(false, 'getAllMusic', args, res);
});
app.get('/api/purchaseMusic', function (req, res) {
    var walletid = req.query.walletid;
    var key = req.query.musickey;
    
    let args = [walletid, key];
    sdk.send(true, 'purchaseMusic', args, res);
});
*/

app.use(express.static(path.join(__dirname, './client')));

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);