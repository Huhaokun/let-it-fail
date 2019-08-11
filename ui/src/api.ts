const axios = require('axios');

function listEndpoints() {
    axios.get('/api/endpoint')
        .then(function (response) {
            console.log("response", response)
        })
        .catch(function (error) {
            console.log(error)
        })
        .finally(function () {
           // do nothing
        })

}