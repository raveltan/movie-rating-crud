import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Vuex from 'vuex'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'

Vue.use(Buefy)

Vue.config.productionTip = false

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        token: null,
        refresh: null,
        baseUrl: 'https://thawing-chamber-91284.herokuapp.com',
        //baseUrl: 'http://localhost:3000',
        loading: true
    },
    mutations: {
        login(state, data) {
            localStorage.setItem('token', data.token)
            localStorage.setItem('refresh', data.refresh)
            state.token = data.token
            state.refresh = data.refresh
        },

        logout(state) {
            localStorage.removeItem('token')
            localStorage.removeItem('refresh')
            state.refresh = null
            state.token = null
        }
    },
    actions: {
        getLogin(context) {
            if (
                localStorage.getItem('token') &&
                localStorage.getItem('refresh')
            ) {
                context.commit('login', {
                    token: localStorage.getItem('token'),
                    refresh: localStorage.getItem('refresh')
                })
            }
        }
    }
})

new Vue({
    router,
    render: h => h(App),
    store: store
}).$mount('#app')
