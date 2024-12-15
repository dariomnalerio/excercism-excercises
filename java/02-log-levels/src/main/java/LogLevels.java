public class LogLevels {

    public static String message(String logLine) {
        int i = logLine.indexOf(":");
        return logLine.substring(i + 1).trim();
    }

    public static String logLevel(String logLine) {
        int i = logLine.indexOf("]");
        return logLine.substring(1, i).toLowerCase();
    }

    public static String reformat(String logLine) {
        return String.format("%s (%s)", message(logLine), logLevel(logLine));
    }
}
