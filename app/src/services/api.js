import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:9000';

class APIHelper {
  health(){
    return axios.get('/health').then(r => {
      console.log(r)
      return r.data
    })
  }

  login(acc, pass){
    return axios.post('/user-service/login', {}, {auth: {username: acc,password: pass}}).then(r => r.status === 200)
  }

  logout(){
    return axios.post('/user-service/logout', {}, {withCredentials: true}).then(r => r.data)
  }

  profile(){
    return axios.get('/user', {withCredentials: true}).then(r => r.data)
  }

  video(vid){
    return axios.get(`/video/${vid}`, {withCredentials: true}).then(r => r.data)
  }

  videos(acc, pass){
    return axios.get(`/videos`, {withCredentials: true}).then(r => r.data)
  }

  processVideo(name, desc, file){
    return axios.post(`/video-service/videos`, {
      name, desc, file
    }, {withCredentials: true}).then(r => r.data)
  }
}

const apiHelper = new APIHelper();
export default apiHelper