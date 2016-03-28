"use strict";

var webdriverjs = require("webdriverio");
var selenium = require('selenium-standalone'),
    util = require('util');
 
describe('Google Search', function() {

  this.timeout(999999);
  var client;

  before(function(done){
    selenium.install({
      logger: function (message) {
        console.log(message);
      }
    }, function (err) {
      if (err) return done(err);
      console.log('Selenuim server installed.');
      selenium.start({
        spawnOptions: {
          stdio: 'inherit'
        }
      },function (err, child) {
        if (err) return done(err);
        selenium.child = child;
        const phantomjsPath = process.cwd() + '/node_modules/phantomjs-prebuilt/bin/phantomjs';
        client = webdriverjs.remote({
          desiredCapabilities: { 
              browserName: 'phantomjs',
             'phantomjs.binary.path': phantomjsPath 
            }
        });
        console.log('client ' + util.inspect(client, {depth: 2}));
        console.log('Selenium server started.');
        client.init().then(function() {
          console.log('client ' + util.inspect(client, {depth: 2}));
          done()
        });
      });
    })}
  );

  it('should work', function (done) {
  // This example shows how you can run a webdriverjs client without starting a selenium server
  // separately.
    console.log('runing test.');
    client
      .url("https://github.com/")
      .getTitle()
      .then(function(title) {
          console.log();
          console.log("GITHUB TITLE: %s", title);
          console.log();
          done();
      });
  });

  after(function(done) {
    client.end().then(function() {
      console.log('Webdriver client stoped. Killing selenium server.');
      selenium.child.kill('SIGKILL');
      console.log('That\'s it!');
      done();
    });
  });
});
