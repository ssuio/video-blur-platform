import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:9000';
axios.defaults.headers.post['Content-Type'] ='application/json;charset=utf-8';
axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

class APIHelper {
  login(acc, pass){
    return axios.post('/user-service/login', {}, {auth: {username: acc,password: pass}}).then(r => r.data)
  }

  logout(){
    return axios.get('/user-service/logout', {}).then(r => r.data)
  }

  profile(){
    return axios.get('/user', {}, {auth: {username: acc,password: pass}}).then(r => r.data)
  }

  video(vid){
    return axios.get(`/video/${vid}`, {}, {auth: {username: acc,password: pass}}).then(r => r.data)
  }

  videos(acc, pass){
    return axios.get(`/videos`, {auth: {username: acc,password: pass}}).then(r => r.data)
  }

  processVideo(name, desc, file){
    return axios.post(`/video-service/videos`, {
      name, desc, file
    }).then(r => r.data)
  }
}

const apiHelper = new APIHelper();
export default apiHelper