<template>
    <table style="width: 100%;">
        <tr>
            <td style="text-align: left;">$</td>
            <td style="text-align: right;">{{ fixedPointValue }}</td>
        </tr>
    </table>
</template>

<script>

export default {
    name: 'currency-table',
    props: {
        value: {
            required: true,
            type: Number
        },
    },
    computed: {
        fixedPointValue() {
            return Number.parseFloat(this.value).toFixed(2);
        }
    }
};
</script>