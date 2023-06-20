package bugfixdemo;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;

@ActivityInterface
public interface BugFixDemoActivities {
    @ActivityMethod()
    String stepOne(String input);

    @ActivityMethod()
    String stepTwo(String input);

    @ActivityMethod()
    String stepThree(String input);
}