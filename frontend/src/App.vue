<template>
  <div id="app">
    <img alt="Wails logo" src="./assets/images/logo.png" class="logo zoomIn item1" >
    <Main class="item2"/>
    <Console class="console"/>
  </div>
</template>

<script>
import Main from "./components/Main.vue";
import Console from "./components/Console.vue"
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Main,
    Console
  }
};
</script>

<style scoped>
  #app {
    height: 100%;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    grid-template-rows: 20px 1fr 1fr minmax(50px, 1fr) 20px;
  }

  .item1 {
    grid-row: 2/3;
    grid-column: 2/3;
    align-self: center;
    justify-self: center;
  }
  .item2 {
    grid-row-start: 3;
    grid-row-end: 4;
    grid-column-start: 2;
    grid-column-end: 3;
  }
  .console{
    grid-row: 4/5;
    grid-column: 1/4;
  }
</style>
