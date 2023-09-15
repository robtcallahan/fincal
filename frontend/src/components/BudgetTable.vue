<script setup>

</script>

<template>
    <table class="table table-hover">
        <thead>
        <th>ID</th>
        <th>Index</th>
        <th>Name</th>
        <th>Group</th>
        <th>Budget Amount</th>
        <th>Budget Period</th>
        </thead>
        <tbody>
        <tr v-for="(row, index) in list"
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
            <td :class="[row.color]"><b-form-input :value=row.name :class="[row.color]" class="input" style="height: 28px;"/></td>
            <td><b-form-select v-model="row.group" :options="groups" :class="[row.color]" class="input" style="height: 28px; width:90%;"></b-form-select></td>
            <td><b-form-input v-model="row.budget_amount" class="input" style="height: 28px;"/></td>
            <td><b-form-select v-model="row.budget_period" :options="periods" class="input" style="height: 28px; width:90%;"></b-form-select></td>
        </tr>
        </tbody>
    </table>
</template>

<script lang="js">

export default {
    name: 'category-table',
    props: ['handleListChange', 'list', 'groups', 'periods'],
    data() {
        return {
            selected: 1,
            groups: [
                { value: null, text: 'Please select an group' },
                { value: "Discretionary", text: "Discretionary" },
                { value: "Non-Discretionary", text: "Non-Discretionary" },
                { value: "Loans", text: "Loans" },
                { value: "Savings", text: "Savings" },
            ],
            periods: [
                { value: null, text: 'Please select an period' },
                { value: "Weekly", text: "Weekly" },
                { value: "Monthly", text: "Monthly" },
                { value: "Every 2 Weeks", text: "Every 2 Weeks" },
                { value: "Twice Monthly", text: "Twice Monthly" },
                { value: "Yearly", text: "Yearly" },
            ]
        }
    },
    methods: {
        handleDragStart(e) {
            e.dataTransfer.setData('text/plain', e.currentTarget.dataset.index)
        },
        handleDragEnter(e) {
            e.currentTarget.classList.add('hover')
        },
        handleDragLeave(e) {
            e.currentTarget.classList.remove('hover')
        },
        handleDragOver(e) {
            e.preventDefault()
        },
        handleDrop(e) {
            e.preventDefault()
            console.log('current Target', e.currentTarget)
            const itemIndex = e.dataTransfer.getData('text/plain'),
                droppedIndex = e.target.parentNode.dataset.index
            Array.from(e.currentTarget.parentNode.children).map(tr => {
                tr.classList.remove('hover')
            })
            this.handleListChange(itemIndex, droppedIndex)
        }
    }
}

</script>

<style scoped>
td {
    padding: 2px 0 !important;
}
.input {
    border-width: 0;
    vertical-align: bottom;
    padding: 0 0 0 2px;
}
.blue {
    background-color: #00CCFF;
}

.yellow {
    background-color: #FFF58C;
}

.green {
    background-color: #80FF00;
}
</style>