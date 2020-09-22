<template>
    <section class="hero is-light is-fullheight">
        <div class="hero-body">
            <div class="container">
                <div class="columns is-centered">
                    <div
                        class="column is-7-tablet is-6-desktop is-5-widescreen"
                    >
                        <form class="box rounded" @submit.prevent="login()">
                            <div
                                class="field has-text-centered title is-3 mt-3"
                            >
                                Movie Rater
                            </div>
                            <div
                                class="field has-text-centered title is-6 mb-3"
                            >
                                {{
                                    isLogin
                                        ? 'Login to existing account'
                                        : 'Register for an account'
                                }}
                            </div>
                            <b-notification
                                type="is-danger"
                                aria-close-label="Close notification"
                                role="alert"
                                v-if="errorMessage"
                            >
                                {{ errorMessage }}
                            </b-notification>
                            <b-field
                                label="Email"
                                :type="emailError ? 'is-danger' : null"
                                :message="emailError"
                            >
                                <b-input
                                    v-model="email"
                                    type="email"
                                    placeholder="E.g. john@gemail.com"
                                    rounded
                                    icon-pack="fas"
                                    aria-required
                                    icon="envelope"
                                    :disabled="isLoading"
                                    icon-right="close-circle"
                                    icon-right-clickable
                                    @icon-right-click="email = ''"
                                ></b-input>
                            </b-field>
                            <b-field
                                label="Password"
                                :type="passwordError ? 'is-danger' : null"
                                :message="passwordError"
                            >
                                <b-input
                                    v-model="password"
                                    type="password"
                                    rounded
                                    :disabled="isLoading"
                                    placeholder="********"
                                    password-reveal
                                    icon-pack="fas"
                                    icon="key"
                                >
                                </b-input>
                            </b-field>
                            <b-field
                                v-if="!isLogin"
                                label="First Name"
                                :type="firstNameError ? 'is-danger' : null"
                                :message="firstNameError"
                            >
                                <b-input
                                    v-model="firstName"
                                    type="text"
                                    placeholder="John"
                                    rounded
                                    icon-pack="fas"
                                    aria-required
                                    :disabled="isLoading"
                                    icon="user-alt"
                                    icon-right="close-circle"
                                    icon-right-clickable
                                    @icon-right-click="firstName = ''"
                                ></b-input>
                            </b-field>
                            <b-field
                                v-if="!isLogin"
                                label="Last Name"
                                :type="lastNameError ? 'is-danger' : null"
                                :message="lastNameError"
                            >
                                <b-input
                                    v-model="lastName"
                                    type="text"
                                    placeholder="Purple"
                                    rounded
                                    icon-pack="fas"
                                    :disabled="isLoading"
                                    aria-required
                                    icon="user-alt"
                                    icon-right="close-circle"
                                    icon-right-clickable
                                    @icon-right-click="lastName = ''"
                                ></b-input>
                            </b-field>
                            <div class="field centered">
                                <b-button
                                    rounded
                                    expanded
                                    outlined
                                    :disabled="isLoading"
                                    @click="isLogin = !isLogin"
                                >
                                    {{
                                        isLogin
                                            ? 'Register for an Account'
                                            : 'Already got an account?'
                                    }}
                                </b-button>
                            </div>
                            <div class="field">
                                <b-button
                                    :loading="isLoading"
                                    native-type="submit"
                                    rounded
                                    expanded
                                    type="is-primary"
                                >
                                    {{ isLogin ? 'Sign In' : 'Register' }}
                                </b-button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </section></template
>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            isLogin: false,
            isLoading: false,
            email: '',
            password: '',
            emailError: '',
            passwordError: '',
            firstName: '',
            firstNameError: '',
            lastName: '',
            lastNameError: '',
            errorMessage: ''
        }
    },
    methods: {
        async login() {
            this.errorMessage = ''
            let error = false
            if (this.password.length < 8) {
                this.passwordError = 'Password should be at least 8 characters'
                error = true
            }

            if (!this.email) {
                this.emailError = 'Email should not be empty'
                error = true
            }
            if (this.isLogin && !error) {
                this.isLoading = true
                try {
                    let result = await axios.post(
                        this.$store.state.baseUrl + '/api/login',
                        {
                            email: this.email,
                            password: this.password
                        }
                    )
                    this.$store.commit('login', {
                        token: result.data.token,
                        refresh: result.data.refresh
                    })
                } catch (e) {
                    if (e) {
                        this.errorMessage =
                            'Unable to login, please check your credentials'
                    }
                } finally {
                    this.isLoading = false
                }
            } else {
                if (this.firstName.length < 3) {
                    this.firstNameError =
                        'First Name should be at least 3 characters'
                    error = true
                }

                if (this.lastName.length < 3) {
                    this.lastNameError =
                        'Last Name should be at least 3 characters'
                    error = true
                }
                if (!error) {
                    this.isLoading = true
                    try {
                        let result = await axios.post(
                            this.$store.state.baseUrl + '/api/register',
                            {
                                email: this.email,
                                password: this.password,
                                firstName: this.firstName,
                                lastName: this.lastName
                            }
                        )
                        this.$store.commit('login', {
                            token: result.data.token,
                            refresh: result.data.refresh
                        })
                    } catch (e) {
                        if (e) {
                            if (e.response.data.error == 'User exists') {
                                this.errorMessage =
                                    'User with with this email is already registered.'
                            } else
                                this.errorMessage =
                                    'Unable to Register, please check your credentials'
                        }
                    } finally {
                        this.isLoading = false
                    }
                }
            }
        }
    }
}
</script>
