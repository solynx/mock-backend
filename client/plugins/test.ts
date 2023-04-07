export default defineNuxtPlugin((/* nuxtApp */) => {
 
    const data1 = useUriQP();
 const test =query_actived();
    return {
      provide: {
        myPlugin: {
            
            showUri : () =>test.value.length > 0 ?"?"+test.value.map((item) => `${item.param}=${item.value}`).join("&"):''
   
        },
        newUri: (uri:string,query:string) => `${uri}${query}` 
      
      }
    }
  })