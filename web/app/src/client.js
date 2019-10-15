import axios from 'axios';

const API_URL = "http://localhost:8089";

const Client = {
  HTTP_POST: "post",
  HTTP_GET: "get",
  HTTP_PUT: "put",
  HTTP_PATCH: "patch",
  HTTP_DELETE: "delete",

  axiosClient: axios.create({
    baseURL: API_URL,
    json: true
  }),

  login(data) {
    this.request(this.HTTP_POST, "/login", data);
  },

  getAllItems() {
    this.request(this.HTTP_GET, this.itemPath(null));
  },

  addNewItem(data) {
    this.request(this.HTTP_PUT, this.itemPath(), data);
  },

  updateItem(itemId, data) {
    this.request(this.HTTP_PATCH, this.itemPath(itemId), data);
  },

  deleteItem(itemId) {
    this.request(this.HTTP_DELETE, this.itemPath(itemId));
  },

  async request(method, path, data) {
    // let accessToken = await Vue.prototype.$auth.getAccessToken();

    return this.axiosClient({
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
  itemPath(itemId = "") {
    if (itemId == null) {
      return "items";
    } else {
      return `item/${itemId}`;
    }
  }
};

export default Client;