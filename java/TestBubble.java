
public class TestBubble {

	public static void main(String[] args) {
		long init = System.currentTimeMillis();
		for (int i = 0 ; i < 1000000 ; i++){
			bubble();
		}
		long end = System.currentTimeMillis();
		System.out.println((end - init)+" ms");
	}
	
	public static void bubble(){
		int[] arrayB = {3,4,1,3,5,1,92,2,4124,424,52,12};
		for (int i = 0; i < arrayB.length; i++) {
	        for (int y = 0; y < arrayB.length - 1; y++) {
	            if (arrayB[y+1] < arrayB[y]) {
	                int t = arrayB[y];
	                arrayB[y] = arrayB[y + 1];
	                arrayB[y + 1] = t;
	            }
	        }
   		}
	}	
}