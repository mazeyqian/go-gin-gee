// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VRouter from 'vue-router'
import Vuex from 'vuex'
import Apple from './components/apple'
import RedApple from './components/redapple'
import Banana from './components/banana'

Vue.use(VRouter)
Vue.use(Vuex)

let store = new Vuex.store({

})

let router = new VRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: '/apple'
    },
    {
      path: '/apple',
      component: {
        viewA: Apple,
        viewB: RedApple
      },
      name: 'applePage',
      children: [
        {
          path: 'red',
          component: RedApple
        }
      ]
    },
    {
      path: '/banana',
      component: Banana
    }
  ]
})

router.push(
  {
    path: 'apple'
  }
)

let mazey = new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})

Vue.use({
  mazey
})
