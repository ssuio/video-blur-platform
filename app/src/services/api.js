import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:9000';
axios.defaults.headers.post['Content-Type'] ='application/json;charset=utf-8';
axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

export default class APIHelper {
  static request() {
    axios.get('/videos', {
      auth: {
        username: 'xssuio@gmail.com',
        password: 'Abcd1234777'
      }
    })
      .then(function (response) {
        // handle success
        console.log(response.data);
      })
      .catch(function (error) {
        // handle error
        console.log(error);
      })
  }
}