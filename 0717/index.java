class Index{
    public static void main(String[] args){
        int[] arr = {1,2,3,4,5,6,7,8,9};
        int index = getIndex(arr,8);
        System.out.println("index:"+index);
    }
    public static int getIndex(int[] arr,int value){
        for(int x=0;x<arr.length;x++){
            if(arr[x]==value){
                return x;
            }
        }
        return -1;
    }
}