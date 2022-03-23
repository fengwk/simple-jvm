package test;

public class Test {
    
    public int testInt = 123123;
    public long testLong = 2L;
    public float testFloat = 3.14f;
    public double testDouble = 999.8;
    public Object testObj = new Object();

    public void testHello() {
        System.out.println("hello");
    }

    public static void main(String[] args) {
        Test test = new Test();
        test.testHello();
    }

}
