<template>
  <b-container>
    <b-row align-v="center">
      <b-col>
        <b-card
        title="Encryption">
          <b-card-text>Encrypt and store data</b-card-text>
          <input type="text" class="form-control" placeholder="Data" id="1">
          <input type="text" class="form-control" placeholder="Key" id="2">
          <b-button variant="primary" @click="addPosts">Submit</b-button>
        </b-card>
      </b-col>
      <b-col>
        <b-card
        title="Fetching">
          <b-card-text>Get data from IPFS</b-card-text>
          <input type="text" class="form-control" placeholder="CID" id="3">
          <input type="text" class="form-control" placeholder="Key" id="4">
          <b-button variant="primary" @click="getPosts">Submit</b-button>
        </b-card>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import axios from "axios"
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data(){
    return {
      formData: {
        data: 'TXkgbmFtZSBpcyBBc3dpbg==',
        key: 'example key 1234',
      },
      formData2: {
        cid: 'QmNkQCdCm2wLsmpCP4itv6bxTuTdPBTzNXGHSbMaLG9fGw',
        key: 'example key 1234',
      },
    }
  },methods: {
    getPosts(){
      // Code to fetch the data from textbox and update into formData2
      axios
        .post('http://localhost:1323/get', this.formData2)
        .then((response) => {
          console.log(response.data)
          this.posts = response.data
        })
        .catch((error) => {
          console.log(error)
        })
    },
    addPosts(){
      // Code to fetch the data from textbox and update into formData
      axios.post('http://localhost:1323/add', this.formData)
        .then(response => console.log(response))
        .catch(error => console.log(error))
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
