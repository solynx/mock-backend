export default defineNuxtPlugin((/* nuxtApp */) => {
 
    const data1 = useUriQP();
 
    return {
      provide: {
        myPlugin: () => 
            "?"+data1.value.map((item) => `${item.param}=${item.value}`).join("&")
   
      
      
      }
    }
  })