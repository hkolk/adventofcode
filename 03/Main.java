import java.io.*;
import java.util.regex.*;

public class Main {
  public static void main(String[] args) throws IOException {
    BufferedReader reader = new BufferedReader(new FileReader("input.txt"));
    Pattern p = Pattern.compile("\\#(\\d+) \\@ (\\d+)\\,(\\d+)\\: (\\d+)x(\\d+)");
    String l = reader.readLine();

    int[][] a = new int[1000][1000];
    int c = 0;
    while (l != null) {
      Matcher m = p.matcher(l);
      m.find();

      int x = Integer.parseInt(m.group(2));
      int y = Integer.parseInt(m.group(3));
      int w = Integer.parseInt(m.group(4));
      int h = Integer.parseInt(m.group(5));

      for (int ii = x; ii < x + w; ii++) {
        for (int i = y; i < y + h; i++) {
          if (a[ii][i] <=1) {
            if (a[ii][i] ==1){
              c++;
            }
            a[ii][i]++;
          }
        }
      }
      l = reader.readLine();
    }
    System.out.println(c);
  }
}
