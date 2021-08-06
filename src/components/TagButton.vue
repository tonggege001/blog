<template>
    <div class="tagButton" @mouseover="mouseOver" @mouseleave="mouseLeave" @click="tagBtnClick" ref="elem">{{tagname}}</div>
</template>

<script>
export default {
    name: "TagButton",
    data(){
        return {
            backgroundColor: "",
        }
    },
    props: {
        tagname: String,
    },
    methods: {
        getColor(){
            const color_list = ["#9fb2dd", "#609cb1", "#72a063", "#46a3e9", "#f6bb41", "#e15b25", 
            "#e6b644", "#c93f69","#c63f45", "#7fb254", "#4792ca"];
            const idx = Math.floor(Math.random() * color_list.length)
            return color_list[idx];
        },

        mouseOver(){
            this.$refs.elem.style.setProperty("background-color", "#39839e")
        },

        mouseLeave(){
            this.$refs.elem.style.setProperty("background-color", this.$data.backgroundColor);
        },

        tagBtnClick(){
            this.$router.push("/tags/"+this.tagname);
        }

        

    },

    mounted:function(){
        var element = document.getElementsByClassName("tagButton");
        for(let elem of element) {
            let color = this.getColor()
            elem.style.setProperty('background-color', color);
            this.$data.backgroundColor = color;
        }
    }
}
</script>


<style>
.tagButton{
    height: 2rem;
    border-radius: 0.5rem;
    font-weight: 700;
    font-size: 1rem;
    font-family: "YouYuan";
    letter-spacing: 0.1rem;
    color: #ffffff;
    box-sizing: border-box;
    display: inline-block;
    margin-inline-end: auto;
    line-height: 0.8rem;
    padding: 0.6rem;
    margin: 0.6rem 0.8rem 0.6rem 0.8rem;
}

.tagButton:hover{
    cursor: pointer;
    transform: scale(1.1);
}


</style>


