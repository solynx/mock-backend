type RowData = {
    key: number;
    param: string;
    value: string;
    description: string;
  };
  const createData = (): RowData[] => [
    {
      key: 1,
      param: "",
      value: "",
      description: "",
    },

  
  ];

export const useUriQP = () =>{
    return useState("data", () => createData());
}
export const query_actived = () =>{
  return useState("query_active", () => []);
}
export const full_url = () =>{
  const full_url= useState("full_url_params", ()=>"a");
  return full_url;
}

