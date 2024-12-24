package Task1;

import java.util.Scanner;

public class BE_Pemrog_ZahidWazifa {
    
    // Fungsi untuk mengecek apakah memori bisa dialokasikan
    public static int hitungAlokasi(int n) {
        // Membuat array boolean untuk tracking nilai yang mungkin
        boolean[] dp = new boolean[n + 1];
        // Array untuk menyimpan jumlah bit yang digunakan
        int[] bitCount = new int[n + 1];
        
        // Inisialisasi kasus dasar
        dp[0] = true;
        
        // Cek untuk setiap nilai dari 1 sampai n
        for (int i = 1; i <= n; i++) {
            // Coba setiap jenis bit [3,5,7]
            for (int bit : new int[]{3, 5, 7}) {
                if (i >= bit && dp[i - bit]) {
                    dp[i] = true;
                    bitCount[i] = bitCount[i - bit] + 1;
                }
            }
        }
        
        // Jika tidak bisa dialokasikan, return -1
        return dp[n] ? bitCount[n] : -1;
    }
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        // Baca jumlah test case
        int t = scanner.nextInt();
        
        // Proses setiap test case
        for (int i = 0; i < t; i++) {
            int n = scanner.nextInt();
            int hasil = hitungAlokasi(n);
            
            // Output hasil
            if (hasil == -1) {
                System.out.println("TIDAK");
            } else {
                System.out.println(hasil);
            }
        }
        
        scanner.close();
    }
}