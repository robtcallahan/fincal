<script setup>
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
</script>

<template>
    <nav v-show="authStore.user" class="navbar navbar-expand navbar-dark bg-dark sticky-top">
        <div class="navbar-nav">
<!--            <router-link to="/" class="nav-item nav-link logo"><img src="../../public/64084-united-symbol-dollar-sign-states-currency_32x32.ico">FinCal</router-link>-->
            <router-link to="/" class="nav-item nav-link logo"><img src="../../public/favicon.ico"> FinCal</router-link>
            <router-link to="/" class="nav-item nav-link">Home</router-link>
            <router-link to="/users" class="nav-item nav-link">Users</router-link>
            <router-link to="/budget" class="nav-item nav-link">Budget</router-link>
            <router-link to="/merchants" class="nav-item nav-link">Merchants</router-link>
            <router-link to="/register" class="nav-item nav-link">Register</router-link>
            <button @click="authStore.logout()" class="btn btn-link nav-item nav-link">Logout</button>
        </div>
    </nav>
</template>

<style scoped>
.logo {
    font-size: 12pt;
    font-weight: bold;
    color: #ffe605;
}
</style>