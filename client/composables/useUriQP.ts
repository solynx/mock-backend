type RowData = {
    key: string;
    param: string;
    value: string;
    description: string;
  };
  const createData = (): RowData[] => [
    {
      key: "aaaa",
      param: "12",
      value: "àdsdfdsf",
      description: "",
    },
    {
      key: "a11",
      param: "2",
      value: "asdasd",
      description: "",
    },
    {
      key: "aa",
      param: "3",
      value: "",
      description: "",
    },
  ];
 
export const useUriQP = () =>{
    return useState("data", () => createData());
}

