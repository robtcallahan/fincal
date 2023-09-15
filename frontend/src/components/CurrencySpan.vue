<template>
    <span ref="inputRef">{{ ref }}</span>
</template>

<script>
import {useCurrencyInput} from 'vue-currency-input'

export default {
    name: 'CurrencySpan',
    props: {
        modelValue: Number, // Vue 2: value
        options: {
            accountingSign: true,
            currency: 'USD',
            currencyDisplay: 'hidden',
            autoDecimalDigits: true,
            precision: 2
        }
    },
    setup(props) {
        const {inputRef} = useCurrencyInput(props.options)
        return {inputRef}
    }
}
</script>