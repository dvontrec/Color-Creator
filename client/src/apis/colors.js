import axios from 'axios';

// creates an axios object that can call methods to this address
export default axios.create({
  baseURL: '/'
});
