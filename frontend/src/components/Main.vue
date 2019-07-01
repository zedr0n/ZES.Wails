
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

  .console{
    grid-column: 1/4;
    grid-row: 3/4;
    border: 1px solid white;
    text-align: left;
  }

  .button{
    display: grid;
    grid-template-columns: 1fr 150px 1fr;
    grid-template-rows: 1fr 65px 1fr;
  }

  .container{
    display: grid;
    grid-template-columns: 80px 1fr 80px;
    grid-template-rows: 80px 1fr 0.5fr 20px;
    height: 100%;
  }

  .value{
    grid-row : 1/2;
    grid-column: 2/3;
  }

  .item22{
    grid-row: 2/3;
    grid-column: 2/3;
  }

  ul {
    display: grid;
    list-style-type: none;
  }
  h1 {
  }
  a:hover {
    font-size: 1em;
    border-color: blue;
    background-color: blue;
    color: white;
    border: 3px solid white;
    border-radius: 10px;
    padding: 9px;
    cursor: pointer;
    transition: 500ms;
  }
  a {
    font-size: 1em;
    border-color: white;
    background-color: #121212;
    color: white;
    border: 3px solid white;
    border-radius: 10px;
    padding: 9px;
    cursor: pointer;
  }
</style>


<template>
  <div class="container">
    <h1 class="value">{{message}}</h1>
    <div class="button item22">
      <a @click="getMessage" class="item22">Query number of roots</a>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: " ",
      log : " ",
    };
  },
  methods: {
    getMessage: function() {
      let self = this;
      window.backend.MyClient.NumberOfRoots().then(result => {
        self.message = result + " roots";
      });
    },

    getLog: function() {
      let self = this;
      window.backend.lastLog().then(result => {
        self.log = result;
      })
    }
  }
};
</script>
