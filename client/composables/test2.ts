
export const useQueryTest1 = () => {
    const { $myPlugin } = useNuxtApp();
 
    return useState("query_param",()=> $myPlugin())
   
}