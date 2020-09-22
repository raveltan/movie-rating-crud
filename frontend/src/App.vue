<template>
    <div>
        <Navbar v-if="login"></Navbar>
        <b-loading is-full-page v-model="loading"></b-loading>
        <Login v-if="!login"></Login>
        <router-view v-if="login"></router-view>
    </div>
</template>

<script>
import Login from '@/views/Login.vue'
import Navbar from '@/components/Navbar.vue'
export default {
    components: {
        Login,
        Navbar
    },
    data() {
        return {
            login: false,
            loading: true
        }
    },
    created() {
        this.initLogin()
    },
    methods: {
        async initLogin() {
            await this.$store.dispatch('getLogin')
            this.login = this.$store.state.token != null
            this.$store.state.loading = false
        }
    },
    watch: {
        '$store.state.token': function() {
            if (this.$store.state.token && this.$store.state.refresh) {
                this.login = true
            } else {
                this.login = false
            }
        },
        '$store.state.loading': function() {
            this.loading = this.$store.state.loading
        }
    }
}
</script>

<style></style>
