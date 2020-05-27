const express = require('express')
const app = express()
const axios = require('axios');
const port = 80

app.get('/', function (req, res) {
  axios.get('http://my-backend:8081')
    .then(response => {
      console.log(response.data);
      res.send('This has been partly fetched from '+ response.data)
    })
    .catch(error => {
      console.log(error);
      res.send(500)
    });
})

app.listen(port, () => console.log(`Example app listening at http://localhost:${port}`))
