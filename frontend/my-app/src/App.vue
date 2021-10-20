<template>
  <div id="app">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import axios from "axios"

export default {
  name: 'App',
  data(){
    return {
      fields:['data','key'],
      items:[],
      posts:[]
    }
  },
  getCID(){
    axios
     .get('http://localhost:1323')
     .then(response => {
       console.log(response.data)
     })
     .catch(error => {
       console.log('There was an error:' + error.response)
     })
  },
  components: {
    HelloWorld
  }
}
</script>

<style lang="scss">
@import "~@/assets/scss/vendors/bootstrap-vue/index";
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
