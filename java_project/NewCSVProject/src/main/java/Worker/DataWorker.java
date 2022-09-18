package Worker;

import Data.Data;

public class DataWorker implements Runnable {

    private Data data;

    public DataWorker(Data d) {
        this.data = d;
    }
    @Override
    public void run() {
        this.data.CalculateAverageBalance();
        this.data.SetThread1ID(((int) Thread.currentThread().getId()));

        this.data.CalculateBenefitFreeTransfer();
        this.data.SetThread2aID(((int) Thread.currentThread().getId()));

        this.data.CalculateBenefitBonusBalance();
        this.data.SetThread2bID(((int) Thread.currentThread().getId()));

        this.data.AddLimitedBonusBalance();
        this.data.SetThread3ID(((int) Thread.currentThread().getId()));
    }
}
