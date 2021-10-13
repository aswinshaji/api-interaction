<template>
  <b-container>
    <b-row align-v="center">
      <b-col>
        <b-card
        title="Encryption">
          <b-card-text>Encrypt and store data</b-card-text>
          <input type="text" class="form-control" placeholder="Data">
          <input type="text" class="form-control" placeholder="Key">
          <b-button variant="primary">Submit</b-button>
        </b-card>
      </b-col>
      <b-col>
        <b-card
        title="Fetching">
          <b-card-text>Get data from IPFS</b-card-text>
          <input type="text" class="form-control" placeholder="CID">
          <input type="text" class="form-control" placeholder="Key">
          <b-button variant="primary">Submit</b-button>
        </b-card>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
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
