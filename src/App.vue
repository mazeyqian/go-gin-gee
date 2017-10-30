<template>
  <div id="app">
    <img src="./assets/logo.png">
    <router-view></router-view>

    <router-link :to="'apple'">to apple</router-link>
    <router-link :to="{path:'banana', param:{color:'yellow'}}">to banana</router-link>
    <router-link :to="{path:'apple/red'}">to red apple</router-link>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
