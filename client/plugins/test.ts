export default defineNuxtPlugin((/* nuxtApp */) => {
 
    const data1 = useUriQP();
 const test =query_actived();
 const full_url_params = full_url();
    return {
      provide: {
        myPlugin: {
            
            showUri : () =>test.value.length > 0 ?"?"+test.value.map((item) => `${item.param}=${item.value}`).join("&"):''
   
        },
        test:(msg:string) => {
          
        
          full_url_params.value = msg ;
          
          return msg
        }
        ,
        newUri: (uri:string,query:string) => `${uri}${query}` 
        
      }
    }
  })