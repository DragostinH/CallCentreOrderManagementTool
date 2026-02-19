import { defineStore } from "pinia";
import { ref } from "vue";

const useSelectedItemsStore = defineStore('selectedItems', () => {
    const selectedProductIds = ref<number[]>([])

    const handleSelectAll = ()=>{
        
    }
})