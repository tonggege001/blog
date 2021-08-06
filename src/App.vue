<template>
  <nav-bar></nav-bar>
  <router-view/>
</template>

<script>
import NavBar from './components/NavBar.vue'
import axios from 'axios'
import 'es6-promise/auto'
import store from "./store"


export default {
  components: {NavBar},
  store: store,

  created(){
    axios
      .get("./__post/meta")
      .then(response => {
          this.$store.commit("InitialData", response.data);

      }).catch(error => {
        console.log(error);
      })
  },
}
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body{
  margin: 0;
  padding: 0;
  font: 300 21px Lato, 'Helvetica Neue', Helvetica, Arial, sans-serif;
  background-color: #333333;
}


</style>
