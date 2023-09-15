<script setup>

</script>

<template>
    <table :handleListChange="handleListChange" class="table table-hover">
        <thead>
        <th>ID</th>
        <th>Index</th>
        <th>Name</th>
        <th>Group</th>
        <th>Budget Amount</th>
        <th>Budget Period</th>
        </thead>
        <tbody>
        <tr v-for="(row, index) in categories"
            :key="index"
            :data-index="index"
            draggable="true"
            @dragstart="handleDragStart"
            @dragenter="handleDragEnter"
            @dragleave="handleDragLeave"
            @dragover="handleDragOver"
            @drop="handleDrop">
            <td class="input" style="height: 28px;">{{ row.id }}</td>
            <td class="input" style="height: 28px;">{{ row.column_index }}</td>
            <td :class="[row.color]">
                <b-form-input :value=row.name
                              @click="startEditing(row, 'name')"
                              @keydown="handleKeydownEvent($event, row.name, row, 'name')"
                              :class="[row.color]"
                              class="input"
                              style="height: 28px;"/>
            </td>
            <td>
                <b-form-select v-model="row.group"
                               :id="getCellId('select', row.id, 'group')"
                               :options="groups"
                               @change="stopEditingAndSave(row, 'column_id')"
                               :class="[row.color]"
                               class="input"
                               style="height: 28px; width:90%;">

                </b-form-select>
            </td>
            <td><b-form-input v-model="row.budget_amount" class="input" style="height: 28px;"/></td>
            <td><b-form-select v-model="row.budget_period" :options="periods" class="input" style="height: 28px; width:90%;"></b-form-select></td>
        </tr>
        </tbody>
    </table>
</template>

<script lang="js">
import axios from 'axios';
const axiosInstance = axios.create({
    baseURL: 'http://localhost:9000/api/',
    timeout: 10000,
});

export default {
    name: 'budget-table',
    props: ['handleListChange', 'list', 'groups', 'periods'],
    data() {
        return {
        }
    },
    methods: {
    }
}

</script>

<style scoped>
</style>