import axios from "axios";

const FETCH_DATA_URL = "http://localhost:1323/add";

const STORE_DATA_URL = "http://localhost:1323/get";

class ipfsapi {

    getCID(){
        return axios.post(FETCH_DATA_URL);
    }

    getKEYstore(){
        return axios.get(STORE_DATA_URL);
    }

    getKEYfetch(){
        return axios.post(FETCH_DATA_URL);
    }

    getData(){
        return axios.get(STORE_DATA_URL);
    }
}

export default new ipfsapi()