<template>
<header>
  <h1>Encryption Key</h1>
  <input type="text" placeholder="Key" v-model="key">
  <br>
  <br>
  <b-container>
    <b-row align-v="center">
      <!-- <form v-on:submit.prevent="getPosts"> -->
      <b-col>
        <b-card
        title="Storing">
          <b-card-text>Encrypt and store data</b-card-text>
          <input type="text" class="form-control" placeholder="Data" v-model="data">
          <b-button variant="primary" @click="addPosts">Store</b-button>
        </b-card>
      </b-col>
      <b-col>
        <b-card
        title="Fetching">
          <b-card-text>Get data from IPFS</b-card-text>
          <input type="text" class="form-control" placeholder="CID" v-model="cid">
          <!-- <input type="text" class="form-control" placeholder="Key" v-model="key_1"> -->
          <!-- <button class="btn btn-primary">Submit</button> -->
          <b-button variant="primary" @click="getPosts">Fetch</b-button>
        </b-card>
      </b-col>
      <!-- </form> -->
    </b-row>
  </b-container>
</header>
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
      
        data: '',
        key: '',
        cid: '',
      
    }
  },methods: {
    getPosts(){
      // Code to fetch the data from textbox and update into formData2
      console.log('CID: ', this.cid)
      console.log('Key: ', this.key)
      console.log('CID: ', this.cid)
      axios
        .post('http://localhost:1323/get', {
          cid: this.cid,
          key: this.key,

          })
        .then((response) => {
          console.log(response.data)
          this.posts = response.data
        })
        .catch((error) => {
          console.log(error)
          this.$msg('Hello')
        })
    },
    addPosts(){
      // Code to fetch the data from textbox and update into formData
      axios.post('http://localhost:1323/add', {
        data: this.data,
        key: this.key
      })
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
