package bugfixdemo;

public class BugFixDemoActivitiesImpl implements BugFixDemoActivities {
    private final boolean HAS_BUG = false;

    @Override
    public String stepOne(String input) {
        System.out.println("Running step 1 with input " + input);
        return "One Done!";
    }

    @Override
    public String stepTwo(String input) {
        if (HAS_BUG) {
            throw new UnsupportedOperationException("Cannot run this thing!");
        } else {
            System.out.println("Running step 2 with input " + input);
        }
        return "Two Done!";
    }

    @Override
    public String stepThree(String input) {
        System.out.println("Running step 3 with input " + input);
        return "Three Done!";
    }
}
