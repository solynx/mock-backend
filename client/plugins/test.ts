export default defineNuxtPlugin((/* nuxtApp */) => {
 
    const data1 = useUriQP();
 
    return {
      provide: {
        myPlugin: {
            
            showUri : () =>"?"+data1.value.map((item) => `${item.param}=${item.value}`).join("&")
   
        },
        newUri: (uri:string,query:string) => `${uri}${query}` 
      
      }
    }
  })