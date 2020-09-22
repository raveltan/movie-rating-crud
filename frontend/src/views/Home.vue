<template
    ><section class="section">
        <div class="container">
            <div class="level">
                <div class="left">
                    <div class="item">
                        <h1 class="title is-5 mb-3">
                            Movie ({{ movie.length }})
                        </h1>
                    </div>
                </div>
                <div class="right">
                    <div class="item">
                        <b-button
                            rounded
                            outlined
                            class="mb-3"
                            @click="addDialog = true"
                        >
                            Add Movie
                        </b-button>
                    </div>
                </div>
            </div>

            <b-modal
                v-model="addDialog"
                has-modal-card
                trap-focus
                :destroy-on-hide="false"
                aria-role="dialog"
                aria-modal
            >
                <template>
                    <form action="" @submit.prevent="addMovie">
                        <div class="modal-card" style="width: auto">
                            <header class="modal-card-head">
                                <p class="modal-card-title">Add New Movie</p>
                                <button
                                    type="button"
                                    class="delete"
                                    @click="addDialog = false"
                                />
                            </header>
                            <section class="modal-card-body">
                                <b-field label="Movie Name">
                                    <b-input
                                        type="text"
                                        v-model="newMovie"
                                        placeholder="Hua Mu Lan"
                                    >
                                    </b-input>
                                </b-field>
                            </section>
                            <footer class="modal-card-foot">
                                <button
                                    class="button"
                                    type="button"
                                    @click="addDialog = false"
                                >
                                    Close
                                </button>
                                <b-button type="is-primary" native-type="submit"
                                    >Add</b-button
                                >
                            </footer>
                        </div>
                    </form>
                </template>
            </b-modal>

            <b-notification
                v-if="error"
                type="is-danger"
                aria-close-label="Close notification"
                role="alert"
                >{{ error }}
            </b-notification>
            <b-table
                :data="movie"
                sticky-header
                striped
                height="50vh"
                :default-sort="['ID', 'asc']"
                paginated
                icon-pack="fas"
            >
                <b-table-column
                    field="ID"
                    label="ID"
                    sortable
                    width="40"
                    centered
                    numeric
                    v-slot="props"
                >
                    {{ props.row.ID }}
                </b-table-column>

                <b-table-column
                    field="Name"
                    label="Movie Name"
                    v-slot="props"
                    width="600"
                >
                    {{ props.row.Name }}
                </b-table-column>

                <b-table-column
                    field="Rating"
                    label="Rating"
                    sortable
                    centered
                    width="40"
                    numeric
                    v-slot="props"
                >
                    <span class="tag is-info is-light">
                        <strong class="is-white">{{ props.row.Rating }}</strong>
                        <span class="icon ">
                            <i class="fas fa-star"></i>
                        </span>
                    </span>
                </b-table-column>
                <b-table-column
                    label="Action"
                    centered
                    width="20"
                    v-slot="props"
                >
                    <b-button
                        size="is-small"
                        rounded
                        @click="$router.push('/review/' + props.row.ID)"
                    >
                        Review
                    </b-button>
                </b-table-column>
            </b-table>
        </div>
    </section></template
>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            movie: [],
            error: '',
            addDialog: false,
            newMovie: ''
        }
    },
    created() {
        this.initData()
        setInterval(() => {
            this.initData()
        }, 15000)
    },
    methods: {
        async addMovie() {
            this.addDialog = false
            if (this.newMovie == '') return
            this.$store.state.loading = true
            try {
                await axios.post(this.$store.state.baseUrl + '/api/movie/add', {
                    name: this.newMovie
                })
                this.initData()
            } catch (e) {
                console.log(e)
                if (e) {
                    this.error = 'Problem communication with the server'
                }
            } finally {
                this.newMovie = ''
            }
        },
        async initData() {
            axios.defaults.headers.common = {
                Authorization: `Bearer ${this.$store.state.token}`
            }
            this.$store.state.loading = true
            try {
                let result = await axios.get(
                    this.$store.state.baseUrl + '/api/movies'
                )
                if (result) {
                    console.log(result)
                    this.movie = result.data
                }
            } catch (e) {
                if (e) {
                    console.log(e)
                    this.error = 'Problem communication with the server'
                }
            } finally {
                this.$store.state.loading = false
            }
        }
    }
}
</script>

<style></style>
