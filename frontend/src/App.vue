<script setup>
import {ref, computed} from 'vue'
import Register from './components/Register.vue'
import Merchants from './components/Merchants.vue'
import Budget from './components/Budget.vue'

const routes = {
    '/': Budget,
    '/budget': Budget,
    '/register': Register,
    '/merchants': Merchants
}

const currentPath = ref(window.location.hash)

window.addEventListener('hashchange', () => {
    currentPath.value = window.location.hash
})

const currentView = computed(() => {
    return routes[currentPath.value.slice(1) || '/'] || NotFound
})
</script>

<style src="../style.css" type="text/css"></style>

<template>
    <b-navbar type="dark" variant="primary" class="navbar-brand">
        <b-navbar-brand href="#">FinCal</b-navbar-brand>
        <b-navbar-nav>
            <b-nav-item href="#">Home</b-nav-item>
            <b-nav-item href="#/budget">Budget</b-nav-item>
            <b-nav-item href="#/merchants">Merchants</b-nav-item>
            <b-nav-item href="#/register">Register</b-nav-item>
        </b-navbar-nav>
    </b-navbar>

    <component :is="currentView"/>
</template>

<style scoped>
#app {
    text-align: center;
    margin: 60px;
}

b-nav-item:active {
    color: white !important;
}
</style>


