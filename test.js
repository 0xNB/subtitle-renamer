const fetch = require('node-fetch');

fetch(`http://www.google.com`).then(response => response.text()).then(text => console.log(text));