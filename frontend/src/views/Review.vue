<template
    ><div>
        <div class="modal">
            <div class="modal-background"></div>
            <div class="modal-card">
                <header class="modal-card-head">
                    <p class="modal-card-title">Add new Review</p>
                    <button
                        class="delete"
                        aria-label="close"
                        @click="hideAddDialog()"
                    ></button>
                </header>
                <section class="modal-card-body">
                    <form @submit.prevent="addReview">
                        <div class="field">
                            <label for="name" class="label">Review</label>
                            <div class="control">
                                <input
                                    type="text"
                                    name="name"
                                    class="input"
                                    v-model="newReview"
                                    placeholder="Very Good"
                                    required
                                />
                            </div>
                        </div>
                        <div class="field">
                            <label for="name" class="label">Rating</label>
                            <div class="control">
                                <input
                                    type="number"
                                    name="nrating"
                                    class="input"
                                    placeholder="5"
                                    required
                                    min="0"
                                    max="5"
                                    v-model="newRating"
                                />
                            </div>
                        </div>

                        <div class="field is-grouped">
                            <div class="control">
                                <button
                                    type="submit"
                                    name="movie"
                                    class="button is-success"
                                >
                                    Add Review
                                </button>
                            </div>
                            <div class="control">
                                <button
                                    type="button"
                                    class="button is-light"
                                    @click="hideAddDialog()"
                                >
                                    Cancel
                                </button>
                            </div>
                        </div>
                    </form>
                </section>
            </div>
        </div>
        <section class="section">
            <div class="container">
                <b-notification
                    v-if="error"
                    type="is-danger"
                    aria-close-label="Close notification"
                    role="alert"
                    >{{ error }}
                </b-notification>
                <div class="level">
                    <div class="level-left" v-if="!notFound">
                        <div class="level-item">
                            <h1 class="title">{{ reviews.name }}</h1>
                        </div>
                        <div class="level-item">
                            <p class="subtitle">{{ reviews.rating }} stars</p>
                        </div>
                    </div>
                    <div class="level-left" v-if="notFound">
                        <div class="level-item">
                            <p class="title is-4">Be the first to add review</p>
                        </div>
                    </div>
                    <div class="level-right">
                        <div class="level-item">
                            <button class="button" @click="showAddDialog()">
                                Add Review
                            </button>
                        </div>
                    </div>
                </div>
                <div class="columns" v-if="!notFound">
                    <div class="column">
                        <div
                            class="card is-6-tablet is-5-desktop is-3-widescreen"
                        >
                            <div
                                class="card-content"
                                v-for="review in reviews.review"
                                :key="
                                    review.Name + review.Review + review.Rating
                                "
                            >
                                <p class="title">“{{ review.Review }}”</p>
                                <p class="subtitle">
                                    {{ review.Name }} ({{ review.Rating }}
                                    stars)
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div></template
>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            newReview: '',
            newRating: 5,
            error: '',
            reviews: {},
            notFound: false
        }
    },
    created() {
        this.initData()
    },
    methods: {
        showAddDialog() {
            let modal = document.querySelector('.modal')
            modal.classList.add('is-active')
        },
        hideAddDialog() {
            let modal = document.querySelector('.modal')
            modal.classList.remove('is-active')
        },
        async addReview() {
            this.hideAddDialog()
            if (this.newReview.length < 1 && !this.newRating) {
                return
            }
            this.$store.state.loading = true
            try {
                let temp = {
                    review: this.newReview,
                    rating: parseInt(this.newRating)
                }

                await axios.post(
                    this.$store.state.baseUrl +
                        '/api/review/' +
                        this.$route.params.id +
                        '/add',
                    temp
                )

                this.error = ''
                this.newReview = ''
                this.newRating = 5
                this.initData()
            } catch (e) {
                if (e) {
                    this.error = 'Problem communication with the server'
                }
            }
        },
        async initData() {
            axios.defaults.headers.common = {
                Authorization: `Bearer ${this.$store.state.token}`
            }
            this.$store.state.loading = true
            try {
                let result = await axios.get(
                    this.$store.state.baseUrl +
                        '/api/review/' +
                        this.$route.params.id
                )
                if (result) {
                    this.reviews = result.data
                }
                this.error = ''
                this.notFound = false
            } catch (e) {
                if (e.response.status == '404') {
                    this.notFound = true
                    return
                } else this.error = 'Problem communication with the server'
            } finally {
                this.$store.state.loading = false
            }
        }
    }
}
</script>

<style></style>
