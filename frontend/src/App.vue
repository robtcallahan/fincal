<script setup>
import Register from './components/Register.vue'
import Merchants from './components/Merchants.vue'
import Budget from './components/Budget.vue'
</script>

<style src="../style.css" type="text/css"></style>

<template>
    <nav class="navbar navbar-expand-lg bg-body-tertiary">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">FinCal</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <a class="nav-link" aria-current="page" href="#">Home</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Register</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Categories</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link active" href="#">Merchants</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <Budget />
</template>

<style scoped>
#app {
    text-align: center;
    margin: 60px;
}
</style>


