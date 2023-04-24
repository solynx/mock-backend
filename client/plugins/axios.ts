import axios from "axios";
export default defineNuxtPlugin((nuxtApp) => {
  
return {
    provide: {
      api_post: async (url:string,data?:object) => {
          await axios.post(url,data).then(function (response) {
            return response.data;
          })
      },
    },
  };
});