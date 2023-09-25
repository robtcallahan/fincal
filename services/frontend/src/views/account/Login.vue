<script setup lang="js">
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';

const schema = Yup.object().shape({
    username: Yup.string().required('Username is required'),
    password: Yup.string().required('Password is required')
});

</script>

<template>
    <div class="card m-3 login">
        <h4 class="card-header">Login</h4>
        <div class="card-body">
            <Form @submit="this.login" :validation-schema="schema" v-slot="{ errors, isSubmitting }">
                <div class="form-group">
                    <label>Username</label>
                    <Field name="username" type="text" class="form-control" :class="{ 'is-invalid': errors.username }" />
                    <div class="invalid-feedback">{{ errors.username }}</div>
                </div>
                <div class="form-group">
                    <label>Password</label>
                    <Field name="password" type="password" class="form-control" :class="{ 'is-invalid': errors.password }" />
                    <div class="invalid-feedback">{{ errors.password }}</div>
                </div>
                <div class="form-group">
                    <button class="btn btn-primary" :disabled="isSubmitting">
                        <span v-show="isSubmitting" class="spinner-border spinner-border-sm mr-1"></span>
                        Login
                    </button>
                </div>
            </Form>
        </div>
    </div>
</template>

<script lang="js">
import axios from 'axios';
import { router } from "@/router/index.js";
import { useAlertStore } from '@/stores';

const axiosInstance = axios.create({
    timeout: 5000
});

export default {
    methods: {
        async login(values) {
            const { username, password } = values;
            const user = {
                username: username,
                password: password
            }
            const json = JSON.stringify(user);
            if (this === 'development') {
                console.log("login() POST: " + this.backendURL + "/auth/login");
                console.log(json);
            }

            axiosInstance
                .post(this.backendURL + "/auth/login", json, {
                    headers: {"Content-Type": "application/json"}
                })
                .then((response) => {
                    if (this === 'development') {
                        console.log("login() response:")
                        console.log(response.data);
                    }
                    // redirect to previous url or default to home page
                    router.push(this.returnUrl || '/');
                })
                .catch((response) => {
                    const r = response.response;
                    console.log(r);
                    const alertStore = useAlertStore();
                    alertStore.error(r.statusText + ': ' + r.data);
                });
        }
    }
}

</script>


<style scoped>
.login {
    width: 350px;
}
label {
    float: left;
}
.form-group {
    margin-top: 20px;
    clear: both;
}
</style>