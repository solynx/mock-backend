import { defineStore } from "pinia";

export const useFullUrlStore = defineStore("uri_mixed",{
    state:() =>({base_url:"", params:"", detail_url :""}),
    actions:{
        getFullUrl (url :string) {
            this.detail_url = url
        }
    }
})