import { createStore } from 'vuex'
import { GetDateFromString } from '../util';

export default createStore({
    state() {
        return {
            post_list: [],
            tags_table: {},
            date_table: {},
        }
    },
    getters: {
        yearList(state){
            return Object.keys(state.date_table).sort().reverse()
        },
        tagList(state){
            return Object.keys(state.tags_table)
        }
    },
    mutations: {
      InitialData(state, rawData){
        rawData = rawData.trim()
        const dataArray = rawData.split("\n");
        for (let entryline of dataArray){
            let entryArr = entryline.split(",");
            let postEntry = {
                modifyDate: parseInt(entryArr[1]),
                html: entryArr[2],
                title: entryArr[3],
                date: GetDateFromString(entryArr[4]),
                tags: entryArr[5].split(" "),
                desc: entryArr[6]
            }
            state.post_list.push(postEntry);
        }
        // 先按照创建时间进行排序
        state.post_list.sort((a,b) => {
            if(a.date.year < b.date.year){
                return 1
            }
            if(a.date.year > b.date.year) {
                return -1
            }

            if(a.date.month < b.date.month) {
                return 1
            }
            if (a.date.mnonth > b.date.month){
                return -1
            }

            if(a.date.day < b.date.day) {
                return 1
            }
            if(a.date.day > b.date.day){
                return -1
            }
            return 0
        })

        // 更新tags_table  
        for(let idx in state.post_list){
             console.log(typeof idx)
            for(let tag of state.post_list[idx].tags) {
                if (tag in state.tags_table){
                    state.tags_table[tag].push(parseInt(idx))
                }
                else {
                    state.tags_table[tag] = [parseInt(idx)]
                }
            }
        }

        // 和date_table  
        for(let idx in state.post_list) {
            if(state.post_list[idx].date.year in state.date_table) {
                state.date_table[state.post_list[idx].date.year].push(parseInt(idx))
            }
            else{
                state.date_table[state.post_list[idx].date.year] = [parseInt(idx)]
            }
        }
      }
    }
  });




