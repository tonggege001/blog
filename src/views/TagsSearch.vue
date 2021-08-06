<template>
    <div class="titleBox">标签</div>
    <div class="HomeMainContainer">
        <item-container v-for="idx in TagList" :key="idx" :item="$store.state.post_list[idx]"></item-container>
    </div>
</template>

<script>
import ItemContainer from "../components/ItemContainer.vue";

export default {
    "name": "TagsSearch",
    components:{
        ItemContainer,
    },

    data(){
        return {
            TagList:[]
        }
    },

    created(){
        const tagname = this.$route.params.tagname;
        this.TagList = this.$store.state.tags_table[tagname]

    }
}
</script>




