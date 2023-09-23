<script setup lang="js">
import CurrencyTable from "./CurrencyTable.vue";
import RawDisplayer from "./RawDisplayer.vue";
</script>

<template>
    <div class="container-sm budget" style="float: left;">
        <div style="float:left;width:30%;">
            <table class="income">
                <tr>
                    <td>Income:</td>
                    <td>
                        <b-form-input v-model="income"
                                      id="income"
                                      class="cell"
                                      style="height: 24px;"
                        />
                    </td>
                </tr>
                <tr>
                    <td>Pay Period:</td>
                    <td>
                        <b-form-select v-model="payPeriod"
                                       :options="periods"
                                       id="pay_period"
                                       class="cell"
                                       style="height: 24px;">
                        </b-form-select>
                    </td>
                </tr>
            </table>
        </div>
        <div style="float:right;margin-right:4px;margin-top:35px;">
            <b-button variant="primary"
                      @click="this.addRow();"
            >Add Category</b-button>
        </div>
        <table class="table">
            <thead>
            <tr>
                <th></th>
                <th>Name</th>
                <th>Group</th>
                <th style="width: 130px;">Budget Amount</th>
                <th style="width: 150px; text-align: center;">Budget Period</th>
                <th style="width: 170px;">{{ payPeriod }} Contribution</th>
                <th></th>
            </tr>
            </thead>
            <draggable
                    :list="categories"
                    tag="tbody"
                    item-key="id"
                    handle=".handle"
                    @start="dragging=true"
                    @end="this.onDragEnd($event)"
            >
                <template #item="{ element: category, index }">
                    <tr @dragstart="onDragStart($event, category)">
                        <td class="drag-icon handle">
                            <b-icon-grip-vertical></b-icon-grip-vertical>
                        </td>
                        <td class="cell" :class="[category.color]">
                            <b-form-input v-model="category.name"
                                          :value=category.name
                                          :id="getCellId('input', category.id, 'name')"
                                          @click="startEditing(category, 'name', index)"
                                          @keydown="handleKeydownEvent($event, category.name, category, 'name')"
                                          :class="[category.budget_group]"
                                          class="cell"
                                          style="height: 28px;"
                            />
                        </td>
                        <td>
                            <b-form-select v-model="category.budget_group"
                                           :id="getCellId('select', category.id, 'budget_group')"
                                           :options="groups"
                                           @change="save(category, 'budget_group')"
                                           :class="[category.budget_group]"
                                           class="cell"
                                           style="height: 28px; width:90%;">

                            </b-form-select>
                        </td>
                        <td @click="startEditing(category, 'budget_amount', index)">
                            <b-form-input v-if="editing"
                                          v-model="category.budget_amount"
                                          :id="getCellId('input', category.id, 'budget_amount')"
                                          @blur="stopEditingAndSave(category, 'budget_amount')"
                                          @keydown="handleKeydownEvent($event, category.budget_amount, category, 'budget_amount')"
                                          class="cell"
                                          style="height: 28px; text-align: right;"
                            />
                            <currency-table v-else :value=category.budget_amount />
                        </td>
                        <td>
                            <b-form-select v-model="category.budget_period"
                                           :options="periods"
                                           :id="getCellId('select', category.id, 'budget_period')"
                                           @change="save(category, 'budget_period')"
                                           class="cell"
                                           style="height: 28px; width:90%; text-align: center;">
                            </b-form-select>
                        </td>
                        <td>
                            <currency-table :value=this.getContribution(category,index) />
                        </td>
                        <td>
                            <div class="cell"
                                 style="margin-left: 20px;"
                                 @click="this.onDeleteClick(category)">
                                <img src="../../red-cross.png" width="20" height="20">
                            </div>
                        </td>
                    </tr>
                </template>
                <template #footer>
                    <tr>
                        <td class="cell-footer" colspan="4"></td>
                        <td class="cell"><b>Total:</b></td>
                        <td><currency-table :value=totalContribution /></td>
                        <td class="cell-footer">&nbsp;</td>
                    </tr>
                    <tr>
                        <td class="cell-footer" colspan="4"></td>
                        <td class="cell"><b>Difference:</b></td>
                        <td :class="{ 'text-danger': incomeDiffClass}"><currency-table :value=incomeDifference /></td>
                        <td class="cell-footer">&nbsp;</td>
                    </tr>
                </template>
            </draggable>
        </table>
    </div>
    <div>
        <b-modal id="modal-dup-name" title="Duplicate Name" hide-footer>
            <p class="my-4"><b>{{ duplicateName }}</b> category already exists.</p>
            <b-button class="mt-3" variant="primary" @click="this.resetDuplicateName">OK</b-button>
        </b-modal>
    </div>
<!--    <raw-displayer class="col-3" :value="categories" title="Categories"/>-->
</template>

<script lang="js">
import draggable from 'zhyswan-vuedraggable'
import axios from 'axios';

const axiosInstance = axios.create({
    // baseURL: 'http://localhost:9000/api/',
    timeout: 10000,
});

draggable.compatConfig = {MODE: 3};

export default {
    name: "Budget",
    components: {
        draggable,
    },
    data() {
        return {
            duplicateName: "",
            income: 600,
            payPeriod: 'Weekly',
            editing: false,
            newCategoryIndex: 1,
            categories: [],
            dragging: false,
            draggedItem: {},
            fields: [
                {
                    key: "id",
                    label: " ",
                    element: "",
                    editable: false,
                    sortable: false
                },
                {
                    key: "name",
                    label: "Name",
                    element: "input",
                    type: "text",
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
                },
                {
                    key: "budget_group",
                    label: "Budget Group",
                    element: "input",
                    type: "text",
                    editable: true,
                    sortable: true,
                },
                {
                    key: "budget_amount",
                    label: "Budget Amount",
                    element: "input",
                    type: "text",
                    editable: true,
                    sortable: true,
                },
                {
                    key: "budget_period",
                    label: "Budget Period",
                    element: "select",
                    type: "text",
                    editable: false,
                    sortable: true
                }
            ],
            sleepMs: 300,
            selectedCell: {},
            fieldKeys: {},
            groups: [
                {value: null, text: 'Please select a group'},
                {value: "Discretionary", text: "Discretionary"},
                {value: "Non-Discretionary", text: "Non-Discretionary"},
                {value: "Charities", text: "Charities"},
                {value: "Loans", text: "Loans"},
                {value: "Savings", text: "Savings"},
            ],
            periods: [
                {value: null, text: 'Please select a period'},
                {value: "Weekly", text: "Weekly"},
                {value: "Monthly", text: "Monthly"},
                {value: "Every 2 Weeks", text: "Every 2 Weeks"},
                {value: "Twice Monthly", text: "Twice Monthly"},
                {value: "Yearly", text: "Yearly"},
            ]
        }
    },
    computed: {
        totalContribution() {
            let total = 0;
            this.categories.forEach(function (item) {
                total += Number.parseFloat(item.pay_period_contribution);
            });
            return total;
        },
        incomeDifference() {
            return 600 - this.totalContribution;
        },
        incomeDiffClass()  {
            if (this.incomeDifference < 0) {
                return true;
            }
            return false;
        }
    },
    methods: {
        async addRow() {
            let newName = "New Category " + this.newCategoryIndex;
            this.categories.forEach(function(item) {
                if (item.name === newName) {
                    this.newCategoryIndex++;
                    newName = "New Category " + this.newCategoryIndex;
                }
            }, this)
            const currentIndex = this.categories[this.categories.length-1].column_index;
            let newCategory = {
                id: -1,
                column_index: currentIndex+1,
                name: newName,
                color: "yellow",
                budget_group: "Non-Discretionary",
                budget_amount: 0,
                budget_period: "Monthly",
                pay_period_contribution: 0
            }
            this.categories[this.categories.length] = newCategory;

            await this.createCategory(newCategory);

            const cellId = this.getCellId("input", this.categories[this.categories.length-1].id, "name");
            const el = document.getElementById(cellId);
            el.scrollIntoView({ behavior: "smooth", block: "start", inline: "nearest" });
        },
        onDragStart(event, item) {
            this.draggedItem = item;
        },
        async onDragEnd(event, item) {
            console.log("oldIndex: " + event.oldIndex + ", newIndex: "  + event.newIndex);
            console.log(this.draggedItem);

            let before = 0;
            let after = this.categories.length;
            if (this.categories[event.newIndex-1])  {
                before = this.categories[event.newIndex-1].column_index;
            }
            if (this.categories[event.newIndex+1]) {
                after = this.categories[event.newIndex+1].column_index;
            }
            this.draggedItem.column_index = (before + after) / 2;
            await this.updateCategory(this.draggedItem, "column_index");
        },
        getContribution(category, index) {
            let contribution = 0;
            switch (category.budget_period) {
                case "Weekly":
                    contribution = category.budget_amount;
                    break;
                case "Monthly":
                    contribution = category.budget_amount * 12 / 52;
                    break;
                case "Yearly":
                    contribution = category.budget_amount / 52;
                    break;
            }
            this.categories[index].pay_period_contribution = contribution;
            return contribution;
        },
        async handleKeydownEvent(event, value, item, fieldKey) {
            if (event.key !== "Enter" && event.key !== "Tab" && event.key !== "Escape") {
                return;
            }
            switch (event.key) {
                case "Escape":
                    await this.cancelEdit(item, fieldKey);
                    break;
                case "Enter":
                    this.selectedCell.previousId = item.id;
                    await this.stopEditingAndSave(item, fieldKey);
                    break;
                case "Tab":
                    event.preventDefault();
                    await this.stopEditingAndSave(item, fieldKey);
                    break;
                default:
            }
        },
        async startEditing(item, fieldKey, index) {
            this.editing = true;
            this.selectedCell = {
                item: item,
                fieldKey: fieldKey,
                rowIndex: index,
                id: item.id,
                previousId: item.id,
                currentValue: item[fieldKey],
                newValue: item[fieldKey],
                dirty: false
            };

            const element = this.fields[this.fieldKeys[fieldKey]].element;
            const elementId = this.getCellId(element, item.id, fieldKey);
            await this.focus(elementId, this.sleepMs);
        },
        async cancelEdit(item, fieldKey) {
            this.editing = false;
            const elementId = this.getCellId('input', this.selectedCell.item.id, 'name');
            await this.blur(elementId, this.sleepMs);
        },
        async stopEditingAndSave(item, fieldKey) {
            this.editing = false;

            // check for duplicate name
            if (fieldKey === 'name') {
                let duplicateCount = 0;
                this.categories.forEach(function(category) {
                    if (category.name === item.name) {
                        duplicateCount++;
                    }
                }, this);

                if (duplicateCount > 1) {
                    this.duplicateName = item.name;
                    this.selectedCell.dirty = false;
                    this.$bvModal.show("modal-dup-name");
                    return;
                }
            }

            if (item[fieldKey] !== this.selectedCell.currentValue) {
                this.selectedCell.dirty = true;
                this.selectedCell.newValue = item[fieldKey];
                await this.updateCategory(item, fieldKey);
            }

            const elementId = this.getCellId('input', this.selectedCell.item.id, 'name');
            await this.blur(elementId, this.sleepMs);
        },
        async resetDuplicateName() {
            this.$bvModal.hide("modal-dup-name");
            this.categories[this.selectedCell.rowIndex].name = this.selectedCell.currentValue;

            const elementId = this.getCellId('input', this.selectedCell.item.id, 'name');
            await this.focus(elementId, this.sleepMs);
        },
        async save(item, fieldKey) {
            this.editing = false;
            await this.updateCategory(item, fieldKey);

            const elementId = this.getCellId("select", item.id, fieldKey);
            document.getElementById(elementId).blur();
        },
        async updateCategory(item, fieldKey) {
            const j = `{"id": ` + item.id + `, "column": "` + fieldKey + `", "value": "` + item[fieldKey] + `"}`;
            console.log("json: " + j);
            axiosInstance
                .put(this.backendURL + "/update_category", j, {
                    headers: {"Content-Type": "application/json"}
                })
                .then((response) => {
                    console.log("category id: " + response.data.id + ", value: " + response.data.value + ", status: " + response.status);
                    this.selectedCell.dirty = false;
                })
                .catch(function (error) {
                    alert(error);
                    // this.categories[this.selectedCell.rowIndex][this.selectedCell.fieldKey].value = this.selectedCell.currentValue;
                    // this.selectedCell.item[this.selectedCell.fieldKey] = this.selectedCell.currentValue;
                });
        },
        async createCategory(item) {
            const c = {
                column_index: item.column_index,
                name: item.name,
                budget_amount: item.budget_amount,
                color: item.color,
                budget_group: item.budget_group,
                budget_period: item.budget_period
            }
            const j = JSON.stringify(c);
            console.log("json: " + j);
            axiosInstance
                .post(this.backendURL + "/create_category", j, {
                    headers: {"Content-Type": "application/json"}
                })
                .then((response) => {
                    const newCategory = response.data;
                    // console.log(newCategory);
                    this.categories[this.categories.length-1] = newCategory;
                    this.selectedCell.dirty = false;
                    return newCategory;
                })
                .catch(function (error) {
                    alert(error);
                });
        },
        onDeleteClick(category) {
            this.confirmDeleteEntry(category);
        },
        async confirmDeleteEntry(category) {
            this.$bvModal.msgBoxConfirm('Are you sure you want to delete the category "' + category.name + '"?', {
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
                        await this.deleteEntry(category);
                    }
                })
                .catch(err => {
                    console.log(err);
                })
        },
        async deleteEntry(category) {
            const j = `{"id": ` + category.id + `}`;
            console.log("deleteCategory() id: ", category.id, ", json: " + j);
            axiosInstance
                .post(this.backendURL + "/delete_category", j, {
                    headers: {"Content-Type": "application/json"}
                })
                .then((response) => {
                    if (response.status === 200) {
                        console.log("category id: " + category.id + " deleted");
                        const index = this.categories.findIndex(item => item.id === category.id) // find the post index
                        if (~index) {
                            this.categories.splice(index, 1) //delete the merchant
                        }
                    } else {
                        console.log(response.statusText);
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        getCellId(el, id, fieldKey) {
            return el + '_' + id + '_' + fieldKey;
        },
        async blur(elementId, ms) {
            await this.sleep(ms);
            const el = document.getElementById(elementId);
            el.blur();
        },
        async focus(elementId, ms) {
            await this.sleep(ms);
            const el = document.getElementById(elementId);
            el.focus();
            el.select();
        },
        sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
        },
        getCategories() {
            axiosInstance
                .get(this.backendURL + "/get_categories")
                .then((response) => {
                    this.categories = response.data;

                    this.fields.map(function(element, index) {
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
.budget {
}

.drag-icon {
    width: 24px;
    height: 24px;
    cursor: pointer;
}

.income {
    margin: 10px;
}

.income td {
    margin: 10px;
    padding: 0 5px 5px 0 !important;
}

th {
    background-color: #31475E !important;
    color: #bfe8fc !important;
    padding: 2px !important;
}

td {
    padding: 0 !important;
}

.cell {
    border: 0;
    border-bottom: 0.01rem solid lightgray;
    vertical-align: bottom;
    --bs-border-radius: 0;
}
.cell-footer {
    border: 0;
    vertical-align: bottom;
    --bs-border-radius: 0;
}


.Discretionary {
    background-color: #80FF00;
}
.Non-Discretionary {
    background-color: #FFF58C;
}
.Charities {
    background-color: violet;
}
.Loans {
    background-color: indianred;
}
.Savings {
    background-color: #00CCFF;
}
</style>
