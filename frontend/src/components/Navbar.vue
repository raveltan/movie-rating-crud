<template>
    <b-navbar shadow fixed-top>
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ path: '/' }">
                <h1 class="title is-4">Movie Rater</h1>
            </b-navbar-item>
        </template>

        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <b-button
                        rounded
                        icon-pack="fas"
                        icon-right="sign-out-alt"
                        expanded
                        @click="$store.commit('logout')"
                    >
                        <strong>Sign Out</strong>
                    </b-button>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>
<script>
export default {}
</script>

<style></style>
