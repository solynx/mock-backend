export const useQueryTest = (data:Array) => {

  
    return useState("query_param",()=> data.map((item) => `${item.param}=${item.value}`).join("&"))
   
}