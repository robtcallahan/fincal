<script setup lang="js">
</script>

<template>
    <b-container fluid>
        <div><h1>Register</h1></div>
        <div id="register" style="height: 800px;">
            <b-table
                    id="register-table"
                    primary-key="id"
                    small
                    bordered
                    sticky-header="100%"
                    class="table-container"
                    label-sort-asc=""
                    label-sort-desc=""
                    label-sort-clear=""
                    :items="register"
                    :fields="fields"
            >
                <template v-slot:cell(source)="{value, item, field: {key, editable, row_id}}">
                    <div v-if="item.row_id % 2 === 0"
                         class="cell-primary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <span v-else class="cell-primary" :id="getCellId('span', item.id, key)">{{ value }}</span>
                    </div>
                    <div v-else
                         class="cell-secondary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <span v-else class="cell-secondary" :id="getCellId('span', item.id, key)">{{ value }}</span>
                    </div>
                </template>
                <template v-slot:cell(date)="{value, item, field: {key, editable, row_id}}">
                    <div v-if="item.row_id % 2 === 0"
                         class="cell-primary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <span v-else class="cell-primary" :id="getCellId('span', item.id, key)">{{ value }}</span>
                    </div>
                    <div v-else
                         class="cell-secondary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <span v-else class="cell-secondary" :id="getCellId('span', item.id, key)">{{ value }}</span>
                    </div>
                </template>
                <template v-slot:cell(name)="{value, item, field: {key, editable, row_id}}">
                    <div v-if="item.row_id % 2 === 0"
                         class="cell-primary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <span v-else class="cell-primary" :id="getCellId('span', item.id, key)">{{ value }}</span>
                    </div>
                    <div v-else
                         class="cell-secondary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <span v-else class="cell-secondary" :id="getCellId('span', item.id, key)">{{ value }}</span>
                    </div>
                </template>
                <template v-slot:cell()="{value, item, field: {key, editable, row_id}}">
                    <div v-if="item.row_id % 2 === 0"
                         class="cell-primary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <div v-else
                             class="currency-row cell-primary">
                            <div class="currency-column-left">$</div>
                            <div class="currency-column-right"
                                 :id="getCellId('span', item.id, key)">{{ value }}</div>
                        </div>
                    </div>
                    <div v-else
                         class="cell-secondary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <div v-else class="currency-row cell-secondary">
                            <div class="currency-column-left">$</div>
                            <div class="currency-column-right"
                                 :id="getCellId('span', item.id, key)">{{ value }}</div>
                        </div>
                    </div>
                </template>
                <template v-slot:cell(bank_register)="{value, item, field: {key, editable, row_id}}">
                    <div v-if="item.row_id % 2 === 0"
                         class="cell-primary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <div v-else
                             class="currency-row cell-primary">
                            <div class="currency-column-left">$</div>
                            <div class="currency-column-right"
                                 :id="getCellId('span', item.id, key)">{{ value }}</div>
                        </div>
                    </div>
                    <div v-else
                         class="cell-secondary"
                         :id="getCellId('div', item.id, key)"
                         :tabindex=0
                         @click="this.onCellClick(value, item, key)"
                         @keydown="this.handleKeydownEvent($event, value, item, key)"
                    >
                        <b-form-input :id="getCellId('input', item.id, key)"
                                      v-if="editable && this.selectedCell.editing"
                                      type="text"
                                      v-model="item[key]"
                                      @click="$event.stopPropagation()"
                                      @keydown="handleKeydownEvent($event, value, item, key)"
                        />
                        <div v-else class="currency-row cell-secondary">
<!--                            <div v-if="item.bank_register !== '' && item.bank_register !== 0" class="currency-column-left">$</div>-->
<!--                            <div v-else class="currency-column-left"></div>-->
                            <div class="currency-column-left"></div>
                            <div class="currency-column-right"
                                 :id="getCellId('span', item.id, key)">
                                <currency-span ref="value" :options="{ currency: 'USD'}" class="cell-secondary"/>
                            </div>
                        </div>
                    </div>
                </template>
            </b-table>
            <pre>{{ register }}</pre>
        </div>
    </b-container>
</template>

<script lang="js">
import CurrencySpan from "./CurrencySpan.vue";

export default {
    name: "Register",
    components: {CurrencySpan},
    data() {
        return {
            fields: [
                {
                    key: "id",
                    label: " ",
                    element: "",
                    thClass: 'd-none',
                    tdClass: 'd-none',
                    editable: false,
                    sortable: false,
                },
                {
                    key: "row_id",
                    label: " ",
                    element: "",
                    thClass: 'd-none',
                    tdClass: 'd-none',
                    editable: false,
                    sortable: false,
                },
                {
                    key: "source",
                    label: "Source",
                    element: "input",
                    type: "text",
                    thStyle: {"width": "140px", "background-color": "#C0C0C0", "padding": "2px"},
                    editable: true,
                    sortable: true,
                },
                {
                    key: "date",
                    label: "Date",
                    element: "input",
                    type: "text",
                    thStyle: {"width": "100px", "background-color": "#C0C0C0", "padding": "2px"},
                    editable: true,
                    sortable: true,
                },
                {
                    key: "name",
                    label: "Name",
                    element: "input",
                    type: "text",
                    thStyle: {"background-color": "#C0C0C0", "padding": "2px"},
                    editable: true,
                    sortable: true,
                },
                {
                    key: "withdrawal",
                    label: "Withdrawals",
                    element: "input",
                    type: "text",
                    formatter: this.formatDollars,
                    thStyle: {"width": "100px", "background-color": "#C0C0C0", "padding": "2px"},
                    editable: true,
                    sortable: true,
                },
                {
                    key: "deposit",
                    label: "Deposits",
                    element: "input",
                    type: "text",
                    formatter: this.formatDollars,
                    thStyle: {"width": "100px", "background-color": "#C0C0C0", "padding": "2px"},
                    editable: true,
                    sortable: true,
                },
                {
                    key: "credit_card",
                    label: "Credit",
                    element: "input",
                    type: "text",
                    formatter: this.formatDollars,
                    thStyle: {"width": "100px", "background-color": "#C0C0C0", "padding": "2px"},
                    editable: true,
                    sortable: true,
                },
                {
                    key: "bank_register",
                    label: "Balance",
                    element: "input",
                    type: "text",
                    formatter: this.formatDollars,
                    thStyle: {"width": "100px", "background-color": "#C0C0C0", "padding": "2px"},
                    editable: true,
                    sortable: true,
                },
            ],
            register: [],
            selectedCell: {},
            previousCell: {
                id: -1,
                item: {},
                fieldKey: ""
            },
            fieldKeys: {},
            editableFieldKeys: {},
            rowsById: {},
            selected: null,
            sleepMs: 300
        }
    },
    methods: {
        formatDollars(value, key, item) {
            if (value === 0) {
                return "-"
            }
            return value;
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
                    this.categories.forEach(function (c, i) {
                        if (c.id === item.column_id) {
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
        async getRegister() {
            axiosInstance
                .get("get_register")
                .then((response) => {
                    this.fieldKeys = {};
                    this.editableFieldKeys = {};
                    this.rowsById = {};

                    this.register = response.data;

                    // this.merchants.map(function (element, index) {
                    //     this.merchants[index].editing = false;
                    //     this.rowsById[element.id] = index;
                    // }, this);
                    //
                    // this.fields.map(function (element, index) {
                    //     this.fieldKeys[element.key] = index;
                    // }, this);
                    //
                    // this.editableFields = this.fields.filter(field => {
                    //     return field.editable === true;
                    // });
                    //
                    // this.editableFields.forEach(function (field, index) {
                    //     this.editableFieldKeys[field.key] = index;
                    // }, this);
                    //
                    // this.filterFields = this.getFilterFields();
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
        await this.getRegister();
    }
}
</script>

<style scoped>
thead, tbody, tfoot, th, td {
    text-align: left;
    width: 100px;
    vertical-align: middle;
    margin: 10px;
}

tr, td {
    text-align: left;
    width: 100px;
    vertical-align: bottom;
    padding: 0 !important;
}

h1 {
    margin: 20px 0 0 20px;
    color: #BFE9FC;
    font-size: 2rem;
}

.table-container {
    margin: 5px 20px 20px 20px;
}

td {
    cursor: default;
}
.currency-column-left {
    float: left;
    width: 10%;
    text-align: left;
}
.currency-column-right {
    float: right;
    width: 90%;
    text-align: right;
}
/* Clear floats after the columns */
.currency-row:after {
    content: "";
    display: table;
    clear: both;
}

div.cell-primary {
    width: 100%;
    height: 100%;
    cursor: default;
    padding: 0 2px;
    border-right: 0.05em solid black;
    vertical-align: bottom;
}

div.cell-secondary {
    width: 100%;
    height: 100%;
    cursor: default;
    background-color: #E7E7E7 !important;
    padding: 0 2px;
    border-right: 0.05em solid black;
    border-bottom: 0.05em solid black;
    vertical-align: bottom;
}

span.cell-primary {
    width: 100%;
    height: 100%;
    cursor: default;
    padding-left: 2px;
    vertical-align: bottom;
}

span.cell-secondary {
    width: 100%;
    height: 100%;
    cursor: default;
    padding-left: 2px;
    vertical-align: bottom;
}

div.cell span input {
    width: 100%;
    cursor: default;
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
