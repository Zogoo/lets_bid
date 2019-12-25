import Vue from 'vue';
import axios from 'axios';

const API_URL = "http://localhost:8080";

axios.defaults.headers.common = {
  "Accept": "application/json",
  "Content-type": "application/x-www-form-urlencoded",
  "X-Requested-With": "XMLHttpRequest",
  "Access-Control-Allow-Origin": "http://localhost:8080"
};

const Client = {
  HTTP_POST: "post",
  HTTP_GET: "get",
  HTTP_PUT: "put",
  HTTP_PATCH: "patch",
  HTTP_DELETE: "delete",

  axiosClient: axios.create({
    baseURL: API_URL,
    json: true,
  }),

  login(data) {
    return this.request(this.HTTP_POST, "/cas/login", data);
  },

  getHelloWorld() {
    return this.request(this.HTTP_GET, "/sso/my_page");
  },

  getAllItems() {
    return this.request(this.HTTP_GET, this.itemPath(null));
  },

  addNewItem(data) {
    return this.request(this.HTTP_PUT, this.itemPath(), data);
  },

  updateItem(itemId, data) {
    return this.request(this.HTTP_PATCH, this.itemPath(itemId), data);
  },

  deleteItem(itemId) {
    return this.request(this.HTTP_DELETE, this.itemPath(itemId));
  },

  async request(method, path, data) {
    let accessToken = await Vue.prototype.$auth.getAccessToken();

    return this.axiosClient({
      method: method,
      url: path,
      data: data,
      headers: {
        "Authorization": `Bearer: ${accessToken}`,
        "Access-Control-Allow-Origin": "*"
      }
    }).then(req => {
      return req.data;
    });
  },

  itemPath(itemId = "") {
    if (itemId == null) {
      return "items";
    } else {
      return `item/${itemId}`;
    }
  }
};

export default Client;