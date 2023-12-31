import axios from "axios";

const {
  API_HOST = 'http://localhost:9000'
} = process.env;

axios.defaults.baseURL = `${API_HOST}`;
// axios.defaults.headers.post["Content-Type"] = "application/x-www-form-urlencoded";
axios.defaults.withCredentials = true;
// axios.defaults.crossDomain = true;

class APIHelper {
  health() {
    return axios.get("/health").then((r) => {
      console.log(r);
      return r.data;
    });
  }

  login(payload) {
    return axios
      .post(
        "/user-service/login",
        {},
        {
          auth: {
            username: payload.account,
            password: payload.passwd
          }
        }
      )
      .then((r) => r.status === 200);
  }

  register(payload) {
    return axios
      .post(
        "user-service/register",
        {
          email: payload.email,
          name: payload.name,
          password: payload.passwd
        },
      )
  }

  logout() {
    return axios.post("/user-service/logout", {}).then((r) => r.data);
  }

  profile() {
    return axios.get("/user").then((r) => r.data);
  }

  video(vid) {
    return axios
      .get(`/video/${vid}`)
      .then((r) => r.data);
  }

  videos() {
    return axios.get(`/videos`).then((r) => r.data);
  }

  updateVideo(id, data) {
    return axios.post(`/video/${id}`, data).then(r => r.data)
  }

  deleteVideo(id) {
    return axios.delete(`/video/${id}`).then(r => r.data)
  }

  transferVideo(name, desc, file, perm) {
    let fd = new FormData()
    fd.append('name', name);
    fd.append('description', desc);
    fd.append('type', 'FACE_BLUR');
    fd.append('perm', perm)
    fd.append('file', file)

    return axios
      .post(
        `/video-service/transfer`,
        fd,
        {
          headers: {
              'Content-Type': 'multipart/form-data'
          }
      }
      )
      .then((r) => r.data);
  }
}

const apiHelper = new APIHelper();
export default apiHelper;
