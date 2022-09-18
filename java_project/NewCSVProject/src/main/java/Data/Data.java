package Data;

public class Data {
    int Id;

    String Name;

    int Age;

    int Balanced;

    int Thread2b;

    int Thread3;

    int PreviousBalanced;

    int AverageBalanced;

    int Thread1;

    int FreeTransfer;

    int Thread2a;

    public Data(int Id, String Name, int Age, int Balanced, int PreviousBalanced, int AverageBalanced, int FreeTransfer) {
        this.Id = Id;
        this.Name = Name;
        this.Age = Age;
        this.Balanced = Balanced;
        this.PreviousBalanced = PreviousBalanced;
        this.AverageBalanced = AverageBalanced;
        this.FreeTransfer = FreeTransfer;
    }

    public void SetThread1ID(int id) {
        this.Thread1 = id;
    }

    public void SetThread2aID(int id) {
        this.Thread2a = id;
    }

    public void SetThread2bID(int id) {
        this.Thread2b = id;
    }

    public void SetThread3ID(int id) {
        this.Thread3 = id;
    }

    public void CalculateAverageBalance() {
        this.AverageBalanced = (this.Balanced + this.PreviousBalanced)/2;
    }

    public void CalculateBenefitFreeTransfer() {
        if (this.Balanced >= 100 && this.Balanced <= 150) {
            this.FreeTransfer = 5;
        }
    }

    public void CalculateBenefitBonusBalance() {
        if (this.Balanced > 150) {
            this.Balanced += 25;
        }
    }

    public void AddLimitedBonusBalance() {
        if (this.Id <= 100) {
            this.Balanced += 10;
        }
    }

    public String[] BuildData() {
        String[] data = {
                String.valueOf(this.Id),
                this.Name,
                String.valueOf(this.Age),
                String.valueOf(this.Balanced),
                String.valueOf(this.Thread2b),
                String.valueOf(this.Thread3),
                String.valueOf(this.PreviousBalanced),
                String.valueOf(this.AverageBalanced),
                String.valueOf(this.Thread1),
                String.valueOf(this.FreeTransfer),
                String.valueOf(this.Thread2a),
        };
        return  data;
    }
}
