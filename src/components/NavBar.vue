<template>
    <div class="navbar">
        <img class="logoimg" src="../assets/bar_logo.png">
        <div class="btnBox">
            <div v-for="(item,index) in nav_list" v-bind:key="item.name" :class="['navBtn', currentIdx==index?'active':'inactive']" @click="navBtnClick(index)">{{item.name}}</div>
        </div>
        <div class="logoImgBox">
            
            <a id="zhihu" href="https://www.zhihu.com/people/tong-shuai-37-29">
                <img src="../assets/zhihu_logo.png">
            </a>
            
            <a id="github"  href="https://github.com/tonggege001">
                <img src="../assets/github_logo.png">
            </a>
        </div>
    </div>
    
</template>


<script>
export default {
    name: "NavBar",
    data(){
        return {
            nav_list: [
                {
                    route: "/home",
                    name: "首页",
                },
                {
                    route: "/archive",
                    name: "归档",
                },
                {
                    route: "/tags",
                    name: "标签",
                },
                {
                    route: "/milestone",
                    name: "里程碑",
                },
                {
                    route: "/about",
                    name: "关于"
                }
            ],
        }
    },
    computed: {
        currentIdx() {
            if (this.$route.path == "/home"){
                return 0;
            }
            if (this.$route.path == "/archive"){
                return 1;
            }
            if (this.$route.path == "/tags"){
                return 2;
            }
            if (this.$route.path == "/milestone"){
                return 3;
            }
            if (this.$route.path == "/about"){
                return 4;
            }
            return -1;
        },
    },
    methods: {
        navBtnClick(idx){
            this.$router.push({
                path: this.nav_list[idx].route,
            })

        }
    }

}
</script>

<style>
.navbar{
    width: 100%;
    height: 80px;
    display: grid;
    grid-template-columns: 3fr 4fr 3fr;
    grid-template-rows: auto;
    background-color: #2b2b2b;
    align-items: center;
    
    color: #f8cc46;
    font-size: 0.75em;
    font-weight: 900;
}

.navbar .logoimg {
    position: relative;
    top: 5px;
    height: 50px;
    padding-left: 100px;
}

.btnBox{
    height: 35px;
    display: flex;
    box-sizing: border-box;
    align-items: center;
    flex-direction: row;
    justify-content: space-around;
}

.navbar .navBtn {
    border-radius: 5px;
    padding: 8px 16px 8px 16px;
}

.navbar .navBtn:hover{
    cursor: pointer;
}

.active {
    background-color: #f8cc46;
    color: #222222;
}

.inactive:hover {
    background-color: #252525;
    box-shadow: 0px 0px 8px 2px rgba(7, 7, 7, 0.2) inset;
}

.logoImgBox {
    height: 60px;
    display: flex;
    box-sizing: border-box;
    flex-direction: row;
    justify-content:left;
    padding-top: 0.7rem;
    padding-left: 10rem;
}

.logoImgBox img{
    width: 35px;
    height: 35px;
    padding-right: 0.5rem;
}


</style>
