<template>
    <b-container fluid>
        <b-row>
            <b-col class="w-75"><h1>Merchants</h1></b-col>
            <b-col class="w-25" style='margin-right: 20px'>
                <b-form-group
                        label-for="filter-input"
                        label-cols-sm="3"
                        label-align-sm="right"
                        label-size="sm"
                        class="mb-0"
                >
                    <b-input-group size="sm">
                        <b-form-input
                                id="filter-input"
                                v-model="filter"
                                type="search"
                                placeholder=" Type to Search"
                                class="search-input"
                        ></b-form-input>

                        <b-input-group-append>
                            <b-button :disabled="!filter" @click="filter = ''">Clear</b-button>
                        </b-input-group-append>
                    </b-input-group>
                </b-form-group>
            </b-col>
        </b-row>
        <div id="table" style="height: 800px;">
            <b-table
                    id="merchants"
                    primary-key="id"
                    small
                    bordered
                    sticky-header="100%"
                    class="table-container"
                    label-sort-asc=""
                    label-sort-desc=""
                    label-sort-clear=""
                    :items="merchants"
                    :categories="categories"
                    :fields="fields"
                    :filter="filter"
                    :filter-function="filterFunction"
                    :filter-included-fields="filterOn"
            >
                <template v-slot:cell(id)="data">
                    <div class="cell"
                         @click="this.onDeleteClick(data)">
                        <img src="../../public/red-cross.png" width="20" height="20">
                    </div>
                </template>
                <template v-slot:cell(column_name)="{value, item, field: {key, editable}}">
                    <div :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="startEditing(item, key)"
                         @keydown="handleKeydownEvent($event, value, item, key)"
                         class="cell"
                         :style="'font-color: ' + item.color"
                    >
                        <b-form-select
                                v-if="editable && item.editing && this.selectedCell.fieldKey === key"
                                :id="getCellId('select', item.id, key)"
                                :options="categories" size="sm"
                                @change="stopEditingAndSave(item, 'column_id')"
                                v-model="item.column_id"
                                :style="'border-width: 0;'">
                        </b-form-select>
                        <span v-else :id="getCellId('span', item.id, key)">{{ item.column_name }}</span>
                    </div>
                </template>
                <template v-slot:cell(tax_deductible)="data">
                    <!--              <span :class="`${getColumnColor(item)}`">-->
                    <!--              <span :style="`color: ${getColumnColor(item)}`">{{ value }}</span>-->
                    <div :id="getCellId('div', data.item.id, data.fieldkey)"
                         :tabindex=-1
                         class="pointer"
                         @click="this.onCellClick(data.value, data.item, data.field.key)">
                        <b-form-checkbox
                                :tabindex=-1
                                value="true"
                                unchecked-value="false"
                                v-model="data.item[data.field.key]"
                        >
                        </b-form-checkbox>
                    </div>
                </template>
                <template v-slot:cell()="{value, item, field: {key, editable}}">
                    <div :id="getCellId('div', item.id, key)"
                         class="cell"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <!--                        v-if="editable && this.selectedCell.editing && this.selectedCell.fieldKey === key"-->
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <span v-else :id="getCellId('span', item.id, key)">{{ value }}</span>
                    </div>
                </template>
            </b-table>
            <pre>{{ merchants }}</pre>
        </div>
    </b-container>
</template>

<script lang="js">
export default {
    name: "Merchants",
    data() {
        return {
            fields: [
                {
                    key: "id",
                    label: " ",
                    element: "",
                    thStyle: {"width": "20px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"},
                    editable: false,
                    sortable: false,
                    filterOn: false
                },
                {
                    key: "name",
                    label: "Name",
                    element: "input",
                    type: "text",
                    thStyle: {"width": "300px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"},
                    editable: true,
                    sortable: true,
                    filterOn: true
                },
                {
                    key: "bank_name",
                    label: "Bank Name",
                    element: "input",
                    type: "text",
                    editable: true,
                    sortable: true,
                    filterOn: true,
                    class: "bank-name-col",
                    thStyle: {"background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"}
                },
                {
                    key: "column_name",
                    label: "Category",
                    element: "select",
                    type: "text",
                    editable: true,
                    sortable: true,
                    filterOn: true,
                    thStyle: {"width": "200px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"}
                },
                {
                    key: "tax_deductible",
                    label: "Tax Ded",
                    element: "",
                    type: "checkbox",
                    editable: false,
                    sortable: true,
                    filterOn: false,
                    thStyle: {"width": "90px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"}
                },
                {
                    key: "column_id",
                    label: "",
                    thClass: 'd-none',
                    tdClass: 'd-none',
                    filterOn: false
                },
                {
                    key: "column_color",
                    label: "",
                    thClass: 'd-none',
                    tdClass: 'd-none',
                    filterOn: false
                }
            ],
            categories: [],
            merchants: [],
            selectedCell: {},
            previousCell: {
                id: -1,
                item: {},
                fieldKey: ""
            },
            fieldKeys: {},
            editableFieldKeys: {},
            rowsById: {},
            filter: null,
            filterOn: [],
            filterFields: [],
            sleepMs: 300
        }
    },
    computed: {
    },
    methods: {
        getFilterFields() {
            return this.fields.filter(f => f.filterOn);
        },
        filterFunction(item, filter) {
            // return this.filterFields.forEach(f => item[f.key].toLowerCase().includes(filter));
            for (let i=0; i<this.filterFields.length; i++) {
                const field = this.filterFields[i];
                if (item[field.key].toLowerCase().includes(filter)) {
                    return true;
                }
            }
            return false;
        },
        onCellClick(value, item, fieldKey) {
            if (item.editing) {
                return;
            }
            if (item.id !== this.previousCell.id && this.previousCell.id !== -1) {
                this.unHighlightCell(this.previousCell.item, this.previousCell.fieldKey)
            }
            this.highlightCell(item, fieldKey);
            this.previousCell = {
                id: item.id,
                item: item,
                fieldKey: fieldKey
            }
        },
        onDeleteClick(data) {
            this.confirmDeleteMerchant(data);
        },
        async handleKeydownEvent(event, value, item, fieldKey) {
            // TODO: doesn't handle fast key repeats
            console.log(event);

            // if editing, allow cursor and alphanumeric keys to their default behavior
            if (item.editing && event.key !== "Enter" && event.key !== "Tab") {
                return;
            }
            switch (event.key) {
                case "Escape":
                    this.cancelEdit(item, fieldKey);
                    break;
                case "Enter":
                    if (this.selectedCell.editing) {
                        this.selectedCell.previousId = item.id;
                        if (!item.editing) {
                            return;
                        }
                        await this.stopEditingAndSave(item, fieldKey);
                        this.cellDownItem(item, fieldKey);
                    } else {
                        if (event.target === event.currentTarget) {
                            await this.startEditing(item, fieldKey);
                        }
                    }
                    break;
                case "Tab":
                    event.preventDefault();
                    if (this.selectedCell.editing) {
                        await this.stopEditingAndSave(item, fieldKey);
                    } else {
                        this.cellRightKey(item, fieldKey);
                    }
                    break;
                case "ArrowRight":
                    this.cellRightKey(item, fieldKey);
                    break;
                case "ArrowLeft":
                    this.cellLeftKey(item, fieldKey);
                    break;
                case "ArrowDown":
                    this.cellDownItem(item, fieldKey);
                    break;
                case "ArrowUp":
                    this.cellUpItem(item, fieldKey);
                    break;
                default:
            }
        },
        cellDownItem(item, fieldKey) {
            console.log("cellDownItem()");
            const newItem = this.merchants[this.selectedCell.rowIndex + 1]
            this.unHighlightCell(item, fieldKey);
            this.highlightCell(newItem, fieldKey)
            this.previousCell = {
                id: newItem.id,
                item: newItem,
                fieldKey: fieldKey
            }
            return newItem;
        },
        cellUpItem(item, fieldKey) {
            console.log("cellUpItem()");
            const newItem = this.merchants[this.selectedCell.rowIndex - 1]
            this.unHighlightCell(item, fieldKey);
            this.highlightCell(newItem, fieldKey)
            this.previousCell = {
                id: newItem.id,
                item: newItem,
                fieldKey: fieldKey
            }
        },
        cellLeftKey(item, fieldKey) {
            console.log("cellLeftKey()");
            let newFieldKey = "";
            if (this.editableFieldKeys[fieldKey] === 0) {
                this.selectedCell.editableFieldsIndex = this.editableFields.length - 1;
            } else {
                this.selectedCell.editableFieldsIndex--;
            }
            newFieldKey = this.editableFields[this.selectedCell.editableFieldsIndex].key;
            this.unHighlightCell(item, fieldKey);
            this.highlightCell(item, newFieldKey)
            this.previousCell = {
                id: item.id,
                item: item,
                fieldKey: newFieldKey
            }
        },
        cellRightKey(item, fieldKey) {
            console.log("cellRightKey()");
            let newFieldKey = "";
            if (this.editableFieldKeys[fieldKey] === this.editableFields.length - 1) {
                this.selectedCell.editableFieldsIndex = 0;
            } else {
                this.selectedCell.editableFieldsIndex++;
            }
            newFieldKey = this.editableFields[this.selectedCell.editableFieldsIndex].key;
            this.unHighlightCell(item, fieldKey)
            this.highlightCell(item, newFieldKey);
            this.previousCell = {
                id: item.id,
                item: item,
                fieldKey: newFieldKey
            }
        },
        highlightCell(item, fieldKey) {
            let el = {};
            // if (item.editing) {
            //     if (fieldKey === "column_name") {
            //         el = document.getElementById(this.getCellId("select", item.id, fieldKey));
            //     } else {
            //         el = document.getElementById(this.getCellId("input", item.id, fieldKey));
            //     }
            //     el.select();
            // } else {
            //     el = document.getElementById(this.getCellId("div", item.id, fieldKey));
            // }

            el = document.getElementById(this.getCellId("div", item.id, fieldKey));
            el.style.border = "1px solid gray";
            el.focus();

            this.selectedCell = {
                fieldKey: fieldKey,
                editable: this.fields[this.fieldKeys[fieldKey]].editable,
                editableFieldsIndex: this.editableFieldKeys[fieldKey],
                editing: false,
                id: item.id,
                currentValue: item[fieldKey],
                newValue: item[fieldKey],
                dirty: false,
                rowIndex: this.rowsById[item.id]
            };
        },
        unHighlightCell(item, fieldKey) {
            const elementId = this.getCellId("div", item.id, fieldKey);
            const element = document.getElementById(elementId);
            element.style.border = "";
            element.blur();
        },
        async startEditing(item, fieldKey) {
            if (!this.fields[this.fieldKeys[fieldKey]].editable) {
                return;
            }
            this.selectedCell = {
                fieldKey: fieldKey,
                editable: this.fields[this.fieldKeys[fieldKey]].editable,
                editableFieldsIndex: this.editableFieldKeys[fieldKey],
                editing: true,
                id: item.id,
                previousId: item.id,
                currentValue: fieldKey === "column_name" ? item.column_id : item[fieldKey],
                newValue: fieldKey === "column_name" ? item.column_id : item[fieldKey],
                dirty: false,
                rowIndex: this.rowsById[item.id]
            };
            this.merchants[this.selectedCell.rowIndex].editing = true;

            const element = this.fields[this.fieldKeys[fieldKey]].element;
            const elementId = this.getCellId(element, item.id, fieldKey);
            if (fieldKey !== "tax_deductible") {
                await this.focus(elementId, this.sleepMs);
            }
            if (fieldKey !== "column_name" && fieldKey !== "tax_deductible") {
                document.getElementById(elementId).select();
            }
        },
        cancelEdit(item, fieldKey) {
            if (!this.selectedCell.editable) {
                return;
            }
            this.merchants[this.selectedCell.rowIndex].editing = false;
            item[fieldKey] = this.selectedCell.currentValue;
        },
        async stopEditingAndSave(item, fieldKey) {
            if (!this.selectedCell.editable) {
                return;
            }
            this.merchants[this.selectedCell.rowIndex].editing = false;

            if (item[fieldKey] !== this.selectedCell.currentValue) {
                this.selectedCell.dirty = true;

                let newSpanValue = "";
                if (fieldKey === "column_id") {
                    this.selectedCell.newValue = item.column_id;

                    // set select text to its new value so the span will display it after editing is set to false
                    this.categories.forEach(function(c, i) {
                        if (c.id === item.column_id)  {
                            item.column_name = c.column_name;
                            newSpanValue = c.text;
                        }
                    }, this);
                } else {
                    this.selectedCell.newValue = item[fieldKey];
                }
                await this.updateMerchant(fieldKey);

                if (fieldKey === 'column_id') {
                    const spanId = this.getCellId('span', item.id, 'column_name');
                    document.getElementById(spanId).innerHTML = newSpanValue;
                    fieldKey = 'column_name';
                }
            }
            this.highlightCell(item, fieldKey);
        },
        async updateMerchant(fieldKey) {
            if (!this.selectedCell.dirty) {
                this.selectedCell.editing = false;
                this.merchants[this.rowsById[this.selectedCell.rowIndex]].editing = false;
                return;
            }
            const j = `{"id": ` + this.selectedCell.id + `, "column": "` + fieldKey + `", "value": "` + this.selectedCell.newValue + `"}`;
            console.log("json: " + j);
            axiosInstance
                .put("update_merchant", j, {
                    headers: {"Content-Type": "application/json"}
                })
                .then((response) => {
                    console.log("merchant id: " + response.data.id + ", value: " + response.data.value + ", status: " + response.status);
                    this.selectedCell.dirty = false;
                })
                .catch(function (error) {
                    alert(error);
                    this.merchants[this.selectedCell.rowIndex][this.selectedCell.fieldKey].value = this.selectedCell.currentValue;
                    this.selectedCell.item[this.selectedCell.fieldKey] = this.selectedCell.currentValue;
                });
            this.selectedCell.editing = false;
            this.merchants[this.selectedCell.rowIndex].editing = false;
        },
        async confirmDeleteMerchant(data) {
            this.$bvModal.msgBoxConfirm('Are you sure you want to delete ' + data.item.name + '?', {
                title: 'Please Confirm',
                size: 'sm',
                buttonSize: 'sm',
                okVariant: 'danger',
                okTitle: 'YES',
                cancelTitle: 'NO',
                footerClass: 'p-2',
                hideHeaderClose: false,
                centered: true
            })
                .then(async value => {
                    if (value === true) {
                        await this.deleteMerchant(data);
                    }
                })
                .catch(err => {
                    console.log(err);
                })
        },
        async deleteMerchant(data) {
            const j = `{"id": ` + data.item.id + `}`;
            console.log("deleteMerchant() id: ", data.item.id, ", json: " + j);
            axiosInstance
                .post("delete_merchant", j, {
                    headers: {"Content-Type": "application/json"}
                })
                .then((response) => {
                    if (response.status === 200) {
                        console.log("merchant id: " + data.item.id + " deleted");
                        const index = this.merchants.findIndex(merchant => merchant.id === data.item.id) // find the post index
                        if (~index) {
                            this.merchants.splice(index, 1) //delete the merchant
                        }
                        this.filterFunction(data.item, this.filter);
                        // this.merchants = this.merchants.filter((merchant, i) => i !== merchant.index);
                    } else {
                        console.log(response.statusText);
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        async getMerchants() {
            axiosInstance
                .get("get_merchants")
                .then((response) => {
                    this.fieldKeys = {};
                    this.editableFieldKeys = {};
                    this.rowsById = {};

                    this.merchants = response.data;

                    this.merchants.map(function (element, index) {
                        this.merchants[index].editing = false;
                        this.rowsById[element.id] = index;
                    }, this);

                    this.fields.map(function (element, index) {
                        this.fieldKeys[element.key] = index;
                    }, this);

                    this.editableFields = this.fields.filter(field => {
                        return field.editable === true;
                    });

                    this.editableFields.forEach(function (field, index) {
                        this.editableFieldKeys[field.key] = index;
                    }, this);

                    this.filterFields = this.getFilterFields();
                })
                .catch(function (error) {
                    alert(error);
                })
        },
        getCellId(el, id, fieldKey) {
            return el + '_' + id + '_' + fieldKey;
        },
        async focus(elementId, ms) {
            if (!this.selectedCell.editable) {
                return;
            }
            await this.sleep(ms);
            const el = document.getElementById(elementId);
            el.focus();
            el.select();
        },
        sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
        },
        async getCategoriesForSelect() {
            axiosInstance
                .get("get_categories_for_select")
                .then((response) => {
                    this.categories = response.data
                })
                .catch(function (error) {
                    alert(error);
                })
        },
        getColumnColor(item) {
            switch (item.column_color) {
                case "blue":
                    // return "td > span.dodgerblue";
                    // return "dodgerblue";
                    return "/deep/ .table > tbody > tr > td:has(span) dodgerblue";
                case "yellow":
                    // return "td > span.gold";
                    return "gold";
                case "green":
                    // return "td > span.forestgreen";
                    return "forestgreen";
                default:
                    return "black";
            }
        },
        formatDateTime(value) {
            const d = new Date(value);
            return `${d.getFullYear()}-${this.padZero(d.getMonth() + 1)}-${this.padZero(d.getDate())} ` +
                `${this.padZero(d.getHours())}:${this.padZero(d.getMinutes())}:${this.padZero(d.getSeconds())}`;
        },
        padZero(value) {
            if (value <= 9) {
                return "0" + value.toString();
            }
            return value;
        }
    },
    async mounted() {
        await this.getCategoriesForSelect();
        await this.getMerchants();
    }
}
</script>

<style scoped>
thead, tbody, tfoot, tr, td, th {
    text-align: left;
    width: 100px;
    vertical-align: middle;
    margin: 10px;
}

h1  {
    margin: 20px 0 0 20px;
    color:#BFE9FC;
    font-size: 2rem;
}
.search-input {
    height: 2rem;
    padding-left: 10px !important;
}

.pointer {
    cursor: default;
}

div.cell {
    width: 100%;
    height: 100%;
    cursor: default;
}

td {
    cursor: default;
}

div.cell span input {
    width: 100%;
    cursor: default;
}

.table-container {
    margin: 5px 20px 20px 20px;
}

.bank-name-col {
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    width: 800px;
    max-width: 800px;
}

input.form-control {
    line-height: 1.0;
    border-radius: inherit;
    padding: 0;
    --bs-border-width: 0;
}

pre {
    text-align: left;
    color: #d63384;
}
</style>