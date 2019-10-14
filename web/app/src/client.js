import Vue from 'vue';
import axios from 'axios';

const API_URL = 'http://localhost:8089';

const axiosClient = axios.create({
  baseURL: API_URL,
  json: true
});

const HTTP_POST = 'post';
const HTTP_GET = 'get';
const HTTP_PUT = 'put';
const HTTP_PATCH = 'patch';
const HTTP_DELETE = 'delete';

const Client = {

  login(data) {
     this.request(HTTP_POST, '/login', data);
  },

  getAllItems() {
    this.request(HTTP_GET, this.itemPath(null));
  },

  addNewItem(data) {
    this.request(HTTP_PUT, this.itemPath(), data);
  },

  updateItem(itemId, data) {
    this.request(HTTP_PATCH, this.itemPath(itemId), data);
  },

  deleteItem(itemId) {
    this.request(HTTP_DELETE, this.itemPath(itemId));
  },

  async request(method, path, data) {
    // let accessToken = await Vue.prototype.$auth.getAccessToken();

    return axiosClient({
      method,
      url: path,
      data,
      headers: {
        // Authorization: `Bearer: ${accessToken}`
      }
    }).then(req => {
      return req.data;
    });
  },
  itemPath(itemId = '') {
    if (itemId == null) {
      return 'items'
    } else {
      return `item/${itemId}`;
    }
  }
};

export default Client;