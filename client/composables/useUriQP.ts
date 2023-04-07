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

