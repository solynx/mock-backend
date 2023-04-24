import { defineStore } from "pinia";

export const useURIStore = defineStore({
    id:'uri_params',
    state: () =>({
        uri:"",
        queries: [] as Object[],
        all_param:"",
        url_changed: ""
    }),
    actions: {
        convertUrl(){
          this.all_param  =  this.queries.length > 0  ? "?"+this.queries.map((item) => `${item.param}=${item.value}`).join("&"):''
           this.url_changed = this.uri+this.all_param
        },
        test(){
            console.log(this.queries)
        }
    }
})