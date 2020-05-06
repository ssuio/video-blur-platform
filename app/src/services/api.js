import axios from "axios";

axios.defaults.baseURL = "http://localhost:9000";
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

  login(acc, pass) {
    return axios
      .post(
        "/user-service/login",
        {},
        {auth: { username: acc, password: pass } }
      )
      .then((r) => r.status === 200);
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

  processVideo(name, desc, file) {
    return axios
      .post(
        `/video-service/videos`,
        {
          name,
          desc,
          file,
        }
      )
      .then((r) => r.data);
  }
}

const apiHelper = new APIHelper();
export default apiHelper;
