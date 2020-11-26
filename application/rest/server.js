const express = require('express');
const app = express();
const mysql = require('mysql');

var path = require('path');
var sdk = require('./sdk');

const PORT = 8080;
const HOST = '192.168.154.130';

var num = 1;

const connection = mysql.createConnection({
    host     : 'localhost',
    user     : 'share',
    password : '1124',
    database : 'sharing_platform'
});

connection.connect();

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
app.post('/api/startShare', function (req, res) {
    var id = req.query.id;
    var target = req.query.target;
    var location = req.query.location;
    var longitude = req.query.longitude;
    var latitude = req.query.latitude;

    var random_string = Math.random().toString(36).substr(2,11);
    random_string = random_string + num;
    num++;
    console.log(random_string);

    connection.query('insert into identity(id, str) values(?, ?)', [id, random_string],(error) => {
        if (error){
            console.log("Query fail...:" + error);
            res.send("Error!");
        }
        else console.log('Query Success');
    });

    let args = [random_string, target, location, longitude, latitude];

    sdk.send(true, 'startShare', args, res, req.query.ccp);
});

app.post('/api/endShare', function (req, res) {
    var id = req.query.id;
    var target = req.query.target;
    var location = req.query.location;
    var longitude = req.query.longitude;
    var latitude = req.query.latitude;

    var random_string = Math.random().toString(36).substr(2,11);
    random_string = random_string + num;
    num++;
    console.log(random_string);

    connection.query('insert into identity(id, str) values(?, ?)', [id, random_string],(error) => {
        if (error){
            console.log("Query fail...:" + error);
            res.send("Error!");
        }
        else console.log('Query Success');
    });

    let args = [random_string, target, location, longitude, latitude];

    sdk.send(true, 'endShare', args, res, req.query.ccp);
});

// agency
app.post('/api/setPlace', function (req, res) {
    var id = req.query.id;
    var location = req.query.location;
    var longitude = req.query.longitude;
    var latitude = req.query.latitude;

    let args = [id, location, longitude, latitude];

    sdk.send(true, 'setPlace', args, res, req.query.ccp);
});

// monitor
app.get('/api/getUserShareRecord', function (req, res) {
    var id = req.query.id;

    let args = [];

    connection.query("select * from identity where id = '" + id + "'", (error, rows, fields) => {
        if (error) throw error;
        else {
            for(var i = 0; i < rows.length; i++){
                console.log(rows[i].str);
                args.push(rows[i].str);
            }
        }
    });

    sdk.send(false, 'getUserShareRecord', args, res, req.query.ccp);
});

app.use(express.static(path.join(__dirname, './client')));

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
