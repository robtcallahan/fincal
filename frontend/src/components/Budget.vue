<script setup lang="js">
import RawDisplayer from "./raw-displayer.vue";
</script>

<template>
    <div class="container-sm">
        <table class="table">
            <thead>
            <tr>
                <th scope="col"></th>
                <th scope="col">ID</th>
                <th scope="col">Index</th>
                <th scope="col">Name</th>
                <th scope="col">Group</th>
                <th scope="col">Amount</th>
                <th scope="col">Period</th>
            </tr>
            </thead>
            <draggable
                    :list="categories"
                    tag="tbody"
                    item-key="id"
                    @start="dragging=true"
                    @end="dragging=false"
            >
                <template #item="{ element: category }">
                    <tr>
                        <td class="drag-icon handle">
                            <b-icon-grip-vertical></b-icon-grip-vertical>
                        </td>
                        <td class="input" style="height: 28px;">{{ category.id }}</td>
                        <td style="height: 28px;">{{ category.column_index }}</td>
                        <td :class="[category.color]">
                            <b-form-input v-model="category.name"
                                          :value=category.name
                                          @click="startEditing(category, 'name')"
                                          @keydown="handleKeydownEvent($event, category.name, category, 'name')"
                                          :class="[category.color]"
                                          class="input"
                                          style="height: 28px;"
                            />
                        </td>
                        <td>
                            <b-form-select v-model="category.budget_group"
                                           :id="getCellId('select', category.id, 'group')"
                                           :options="groups"
                                           @change="stopEditingAndSave(category, 'column_id')"
                                           :class="[category.color]"
                                           class="input"
                                           style="height: 28px; width:90%;">

                            </b-form-select>
                        </td>
                        <td>
                            <b-form-input v-model="category.budget_amount"
                                          class="input"
                                          style="height: 28px;"
                            />
                        </td>
                        <td>
                            <b-form-select v-model="category.budget_period"
                                           :options="periods"
                                           class="input"
                                           style="height: 28px; width:90%;">
                            </b-form-select>
                        </td>
                    </tr>
                </template>
            </draggable>
        </table>
    </div>
    <rawDisplayer class="col-3" :value="categories" title="Categories"/>
</template>

<script lang="js">
import draggable from 'zhyswan-vuedraggable'

// import "sortablejs/Sortable.min.js"
// import "vuedraggable/dist/vuedraggable.umd.js"

import axios from 'axios';

const axiosInstance = axios.create({
    baseURL: 'http://localhost:9000/api/',
    timeout: 10000,
});

draggable.compatConfig = { MODE: 3 };

export default {
    name: "Budget",
    props: ['groups', 'periods'],
    components: {
        draggable,
    },
    data() {
        return {
            categories: [],
            dragging: false,
            fields: [
                {
                    key: "id",
                    label: " ",
                    element: "",
                    thStyle: {"width": "20px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"},
                    editable: false,
                    sortable: false
                },
                {
                    key: "name",
                    label: "Name",
                    element: "input",
                    type: "text",
                    thStyle: {"width": "300px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"},
                    editable: true,
                    sortable: true
                },
                {
                    key: "color",
                    label: "Color",
                    element: "input",
                    type: "text",
                    editable: true,
                    sortable: true,
                    class: "bank-name-col",
                    thStyle: {"background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"}
                },
                {
                    key: "budget_group",
                    label: "Budget Group",
                    element: "input",
                    type: "text",
                    editable: true,
                    sortable: true,
                    thStyle: {"width": "200px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"}
                },
                {
                    key: "budget_amount",
                    label: "Budget Amount",
                    element: "input",
                    type: "text",
                    editable: true,
                    sortable: true,
                    thStyle: {"width": "200px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"}
                },
                {
                    key: "budget_period",
                    label: "Budget Period",
                    element: "select",
                    type: "text",
                    editable: false,
                    sortable: true, filterOn: false,
                    thStyle: {"width": "90px", "background-color": "#31475E", "color": "#bfe8fc", "padding": "2px"}
                }
            ],
            sleepMs: 300,
            selectedCell: {},
            previousCell: {
                id: -1,
                item: {},
                fieldKey: ""
            },
            fieldKeys: {},
            groups: [
                {value: null, text: 'Please select an group'},
                {value: "Discretionary", text: "Discretionary"},
                {value: "Non-Discretionary", text: "Non-Discretionary"},
                {value: "Loans", text: "Loans"},
                {value: "Savings", text: "Savings"},
            ],
            periods: [
                {value: null, text: 'Please select an period'},
                {value: "Weekly", text: "Weekly"},
                {value: "Monthly", text: "Monthly"},
                {value: "Every 2 Weeks", text: "Every 2 Weeks"},
                {value: "Twice Monthly", text: "Twice Monthly"},
                {value: "Yearly", text: "Yearly"},
            ]
        }
    },
    methods: {
        async handleKeydownEvent(event, value, item, fieldKey) {
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
                    }
                    break;
                default:
            }
        },
        async startEditing(item, fieldKey) {
            this.selectedCell = {
                fieldKey: fieldKey,
                editing: true,
                id: item.id,
                previousId: item.id,
                currentValue: item[fieldKey],
                newValue: item[fieldKey],
                dirty: false
            };
            this.categories[this.selectedCell.rowIndex].editing = true;

            const element = this.fields[this.fieldKeys[fieldKey]].element;
            const elementId = this.getCellId(element, item.id, fieldKey);
            await this.focus(elementId, this.sleepMs);
            document.getElementById(elementId).select();
        },
        cancelEdit(item, fieldKey) {
            if (!this.selectedCell.editable) {
                return;
            }
            this.categories[this.selectedCell.rowIndex].editing = false;
            item[fieldKey] = this.selectedCell.currentValue;
        },
        async stopEditingAndSave(item, fieldKey) {
            if (!this.selectedCell.editable) {
                return;
            }
            this.categories[this.selectedCell.rowIndex].editing = false;

            if (item[fieldKey] !== this.selectedCell.currentValue) {
                this.selectedCell.dirty = true;

                let newSpanValue = "";
                if (fieldKey === "column_id") {
                    this.selectedCell.newValue = item.column_id;

                    // set select text to its new value so the span will display it after editing is set to false
                    this.categories.forEach(function (c) {
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
        async updateCategory(fieldKey) {
            if (!this.selectedCell.dirty) {
                this.selectedCell.editing = false;
                this.categories[this.rowsById[this.selectedCell.rowIndex]].editing = false;
                return;
            }
            const j = `{"id": ` + this.selectedCell.id + `, "column": "` + fieldKey + `", "value": "` + this.selectedCell.newValue + `"}`;
            console.log("json: " + j);
            axiosInstance
                .put("update_category", j, {
                    headers: {"Content-Type": "application/json"}
                })
                .then((response) => {
                    console.log("category id: " + response.data.id + ", value: " + response.data.value + ", status: " + response.status);
                    this.selectedCell.dirty = false;
                })
                .catch(function (error) {
                    alert(error);
                    this.categories[this.selectedCell.rowIndex][this.selectedCell.fieldKey].value = this.selectedCell.currentValue;
                    this.selectedCell.item[this.selectedCell.fieldKey] = this.selectedCell.currentValue;
                });
            this.selectedCell.editing = false;
            this.categories[this.selectedCell.rowIndex].editing = false;
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
        handleDragStart(e) {
            e.dataTransfer.dropEffect = 'move'
            e.dataTransfer.effectAllowed = 'move'
            e.dataTransfer.setData('text/plain', e.currentTarget.dataset.index)
        },
        handleDragEnter(e) {
            e.preventDefault()
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
            // this.handleListChange(itemIndex, droppedIndex)
            const oldIndex = parseInt(itemIndex, 10),
                newPositionIndex = parseInt(droppedIndex, 10),
                movedItem = this.categories[oldIndex]
            this.categories.splice(oldIndex, 1)
            this.categories.splice(newPositionIndex, 0, movedItem)
        },
        getCategories() {
            axiosInstance
                .get("get_categories")
                .then((response) => {
                    this.categories = response.data;

                    this.categories.map(function (element, index) {
                        this.categories[index].editing = false;
                    }, this);

                    this.fields.map(function (element, index) {
                        this.fieldKeys[element.key] = index;
                    }, this);
                })
                .catch(function (error) {
                    alert(error);
                })
        }
    },
    mounted() {
        this.getCategories();
    }
}

</script>

<style scoped>
.container {
    display: flex;
    height: 100vh;
    justify-content: center;
    align-items: center;
}

.container table tbody tr[draggable=true] {
    cursor: move;
}

.container table tbody tr[draggable=true].hover {
    margin-top: 49px;
    opacity: 0.2;
}

.drop-zone {
    background-color: #eee;
    margin-bottom: 10px;
    padding: 10px;
}

.drag-el {
    background-color: #fff;
    margin-bottom: 10px;
    padding: 5px;
}

.drag-icon {
    width: 24px;
    height: 24px;
    cursor: pointer;
}

td {
    padding: 0 !important;
}

.input {
    border: 0;
    border-bottom: 0.01rem solid lightgray;
    vertical-align: bottom;
    padding: 5px 0 0 2px;
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
