

import Data.Data;
import Worker.DataWorker;
import com.opencsv.CSVWriterBuilder;
import com.opencsv.ICSVWriter;

import java.io.File;
import java.io.FileWriter;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class LearnCsv {
    public static void main(String[] args) throws Exception {

        ArrayList<Data> data = new ArrayList<>();
        String splitBy = ";";
        Scanner sc = new Scanner(new File("/Users/user/Documents/IntellijFolder/NewCSVProject/src/main/resources/BeforeEOD.csv"));

        while (sc.hasNext())  //returns a boolean value
        {
            System.out.println(sc.nextLine());
            String[] row = sc.next().split(splitBy);
            Data inp = new Data(
                    Integer.parseInt(row[0]),
                    row[1],
                    Integer.parseInt(row[2]),
                    Integer.parseInt(row[3]),
                    Integer.parseInt(row[4]),
                    Integer.parseInt(row[5]),
                    Integer.parseInt(row[6])
                    );

            data.add(inp);
        }
        sc.close();  //closes the scanner

        ExecutorService executor = Executors.newFixedThreadPool(8);

        data.forEach(d -> {
//            d.CalculateAverageBalance();
//            d.CalculateBenefitFreeTransfer();
//            d.CalculateBenefitBonusBalance();
//            d.AddLimitedBonusBalance();
//            System.out.println(d);

            Runnable worker = new DataWorker(d);

            executor.execute(worker);
        });

        executor.shutdown();
        while (!executor.isTerminated()) {}

        List<String[]> csvData = createCsvData(data);

        try (ICSVWriter writer = new CSVWriterBuilder(new FileWriter("/Users/user/Documents/IntellijFolder/NewCSVProject/src/main/resources/AfterEOD.csv")).withSeparator(';').build()) {
            writer.writeAll(csvData);
        }
    }

    private static List<String[]> createCsvData(ArrayList<Data> data) {
        List<String[]> list = new ArrayList<>();

        String[] header = {
                "id",
                "Name",
                "Age",
                "Balanced",
                "No 2b Thread-No",
                "No 3 Thread-No",
                "Previous Balanced",
                "Average Balanced",
                "No 1 Thread-No",
                "Free Transfer",
                "No 2a Thread-No"
        };

        list.add(header);

        data.forEach(d -> list.add(d.BuildData()));

        return  list;
    }
}


