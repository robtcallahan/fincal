<script setup lang="js">

// import {inject} from "vue"
// import {defineComponent} from 'vue'
import "./Merchants.vue";

</script>

<template>
    <div id="table">
        <b-table
                primary-key="id"
                inheritAttrs
                striped
                hover
                small
                bordered
                sticky-header
                head-variant="dark"
                label-sort-asc=""
                label-sort-desc=""
                label-sort-clear=""
                :items="items"
                :fields="fields"
        >
            <template #cell(tax_deductible)="data">
                <span v-if="data.value">Yes</span>
                <span v-else>No</span>
            </template>
<!--            <template v-slot:cell()="{ value, item, field: { key }}">-->
<!--                <template v-if="!item.isEditing" @click="handleEdit">{{ value }}</template>-->
<!--                <b-form-input v-else v-model="item[key]" />-->
<!--            </template>-->
            <template v-slot:cell(name)="row" v-if="edit">
                <b-form-input v-if="items[data.index].isEditing && selectedCell === 'name'" type="text"
                              v-model:name="items[data.index].name"></b-form-input>
                <span v-else @click="handleEdit(data, 'name')">{{ items[data.index].name }}</span>
            </template>
            <template #cell(bank_name)="data">
                <b-form-input v-if="items[data.index].isEditing && selectedCell === 'bank_name'" type="text"
                              v-model:name="items[data.index].bank_name"
                              :value=items[data.index].bank_name></b-form-input>
                <span v-else @click="handleEdit(data, 'bank_name')">{{ data.value }}</span>
            </template>
<!--            <template #cell(column_id)="data">-->
<!--                <b-form-input v-if="items[data.index].isEdit && selectedCell === 'column_id'" type="text"-->
<!--                              v-model:name="items[data.index].column_id"-->
<!--                              :value=items[data.index].column_id></b-form-input>-->
<!--                <span v-else @click="editCellHandler(data, 'column_id')">{{ data.value }}</span>-->
<!--            </template>-->
<!--            <template #cell(tax_deductible)="data">-->
<!--                <b-form-input v-if="items[data.index].isEdit && selectedCell === 'tax_deductible'" type="text"-->
<!--                              v-model:name="items[data.index].tax_deductible"-->
<!--                              :value=items[data.index].tax_deductible></b-form-input>-->
<!--                <span v-else @click="editCellHandler(data, 'tax_deductible')">{{ data.value }}</span>-->
<!--            </template>-->
<!--            <template #cell(created_at)="data">-->
<!--                <b-form-datepicker v-if="items[data.index].isEdit && selectedCell === 'created_at'" type="text"-->
<!--                                   v-model:name="items[data.index].created_at" :value=items[data.index].created_at-->
<!--                                   locale="en"></b-form-datepicker>-->
<!--                <span v-else @click="editCellHandler(data, 'created_at')">{{ data.value }}</span>-->
<!--            </template>-->
<!--            <template #cell(updated_at)="data">-->
<!--                <b-form-datepicker v-if="items[data.index].isEdit && selectedCell === 'updated_at'" type="text"-->
<!--                                   v-model:name="items[data.index].updated_at" :value=items[data.index].updated_at-->
<!--                                   locale="en"></b-form-datepicker>-->
<!--                <span v-else @click="editCellHandler(data, 'updated_at')">{{ data.value }}</span>-->
<!--            </template>-->
        </b-table>
<!--        <pre>{{ items }}</pre>-->
    </div>
</template>

<script lang="js">

import axios from "axios";
const instance = axios.create({
    baseURL: 'https://localhost:9000/api/',
    timeout: 1000,
});

export default {
    name: "Merchants",
    data() {
        return {
            fields: [
                {
                    key: "name",
                    label: "Name",
                    type: "text",
                    thStyle: { "width": "300px" },
                    editable: true,
                    sortable: true
                },
                {
                    key: "bank_name",
                    label: "Bank Name",
                    type: "text",
                    editable: true,
                    sortable: true,
                    class: "bank-name-col",
                },
                {
                    key: "column_id",
                    label: "Col ID",
                    type: "text",
                    editable: true,
                    sortable: true,
                    thStyle: { "width": "70px" }
                },
                {
                    key: "tax_deductible",
                    label: "Tax Ded",
                    type: "checkbox",
                    editable: true,
                    sortable: true,
                    thStyle: { "width": "90px" }
                },
                {
                    key: "created_at",
                    label: "Created",
                    editable: false,
                    sortable: true,
                    type: "text",
                    tdClass: "date-col",
                    thStyle: { "width": "180px" },
                    formatter: this.formatDateTime
                },
                {
                    key: "updated_at",
                    label: "Updated",
                    editable: false,
                    sortable: true,
                    type: "text",
                    tdClass: "date-col",
                    thStyle: { "width": "180px" },
                    formatter: this.formatDateTime
                }
            ],
            items: [],
            isEditing: false,
            selectedCell: null,
        };
    },
    methods: {
        async fetchMerchants() {
            instance
                .get("get_merchants")
                .then((response) => {
                    this.items = response.data
                });
        },
        formatDateTime(value) {
            const d = new Date(value);
            const dateString = `${d.getFullYear()}-${this.padZero(d.getMonth()+1)}-${this.padZero(d.getDate())} ` +
                `${this.padZero(d.getHours())}:${this.padZero(d.getMinutes())}:${this.padZero(d.getSeconds())}`;
            return dateString
        },
        padZero(value) {
            if (value <= 9) {
                return "0" + value.toString();
            }
            return value
        },
        handleEdit(data, key) {
            this.items = this.items.map(item => ({...item, isEditing: false}));
            this.items[data.index].isEditing = true;
            this.selectedCell = key
        },
        handleSubmit(data, update) {
            this.rowUpdate = {
                edit: false,
                // id: data.id,
                action: update ? "update" : "cancel",
            };
        },
        handleDelete(data) {
            // this.rowUpdate = {id: data.id, action: "delete"};
            this.rowUpdate = {action: "delete"};
        },
    },
    async mounted() {
        await this.fetchMerchants()
        this.items = this.items.map(item => ({...item, isEdit: false}));
    },
    watch: {
        items: {
            handler(val, oldVal) {
                // console.log(oldVal + ' --> ' + val)
            },
            deep: true
        }
    }
}
</script>

<style scoped>
#app {
    text-align: center;
    margin: 60px;
}

thead, tbody, tfoot, tr, td, th {
    text-align: left;
    width: 100px;
    vertical-align: middle;
    margin: 10px;
}

span {
    cursor: pointer;
}

div.table-container {
    margin: 20px;
}

.date-col {
    text-align: justify;
}

.bank-name-col {
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    width: 800px;
    max-width: 800px;
}
input.form-control  {
    border-radius: inherit;
    padding: 0;
}
pre {
    text-align: left;
    color: #d63384;
}
</style>