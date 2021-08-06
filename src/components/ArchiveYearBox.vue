<template>
    <div class="archiveYearBox">
        <h2 class="year">
            {{year}}
        </h2>
        <div class="archive-list">
            <div class="archive-list-item"  v-for="item_idx in $store.state.date_table[year]" :key="item_idx">
                
                <div class="archive-list-item-date">{{formattedDate(item_idx)}}</div>
                <div class="archive-list-item-title">
                    <a :href="'/__post/'+$store.state.post_list[item_idx].html">
                        {{$store.state.post_list[item_idx].title}}
                    </a>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "ArchiveYearBox",
    props: {
        year: String,
    },

    methods: {
        formattedDate(idx){
            return this.$store.state.post_list[idx].date.year+"年"+
            this.$store.state.post_list[idx].date.month+"月"
            +this.$store.state.post_list[idx].date.day+"日"
        }
    }, 
}
</script>

<style>
.archiveYearBox {
    width: 100%;

    display: flex;
    flex-direction: column;
    box-sizing: border-box;
    padding: 2rem 4rem 2rem 4rem;

}

.archiveYearBox .year{
    color: #dddddd;
    font-size: 3rem;
    font-weight: 900;

    display: inline-block;
    /* margin-inline是水平方向，margin-block是垂直方向 */
    margin-block-start: 0; 
    margin-block-end: 0;
    margin-inline-start: 0;
    margin-inline-end: 0;

}

.archiveYearBox .archive-list-item {
    width: 100%;
    font-size: 1.2rem;
    font-weight: 550;
    color: #dddddd;
    box-sizing: border-box;
    display: flex;
    flex-direction: row;
    justify-content: space-between;

    margin-block-end: 1rem;


}

.archiveYearBox .archive-list-item .archive-list-item-date{
    width: 15%;
    font-size: 1rem;
    text-align: left;
    margin-inline-start: 0.5rem;
}

.archiveYearBox .archive-list-item .archive-list-item-title{
    width: 85%;
    text-align: left;
}

.archiveYearBox .archive-list-item .archive-list-item-title a{
    font-size: 1.2rem;
    font-weight: 550;
    color: #dddddd;
    text-decoration: none;
}

</style>
