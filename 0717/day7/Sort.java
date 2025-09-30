public class Sort{
    public static void main(String[] args){
        int[] arr = {9999,2,3,100,5,1,7,8,9};
        // printArray(arr);
        System.out.println("----------------");
        diySort(arr);
    }

    public static void diySort(int[] arr){
        for (int y=0;y<arr.length-1;y++){
            int minX = y;
            for(int x=y;x<arr.length;x++){
                if(arr[minX]>= arr[x]){
                  minX = x;
                }
            }
            int temp = arr[y];
            arr[y] = arr[minX];
            arr[minX] = temp;
          

        }
        printArray(arr);
        System.out.println("----------------");
    }

    public static void printArray(int[] arr){
        System.out.print("[");
        for(int x=0;x<arr.length;x++){
            if(x==arr.length-1){
                System.out.print(arr[x]);
            }else{
                System.out.print(arr[x]+",");
            }
        }
        System.out.println("]");
    }
}