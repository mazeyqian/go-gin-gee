<template>
  <div id="app">
    <img src="./assets/logo.png">
    <transition name="fade">
      <keep-alive>
        <router-view></router-view>
      </keep-alive>
    </transition>

    <router-link :to="'apple'" tag="li">to apple</router-link>
    <router-link :to="{name: 'applePage'}">to banana</router-link>
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
