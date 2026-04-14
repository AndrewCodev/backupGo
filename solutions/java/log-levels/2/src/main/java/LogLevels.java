public class LogLevels {
    
    public static String message(String logLine) {
        int pos = logLine.indexOf(":");
        return logLine.substring(pos + 1).trim();
    }

    public static String logLevel(String logLine) {
        int initPos = logLine.indexOf("[");
        int endPos = logLine.indexOf("]");
        return logLine.substring(initPos + 1, endPos).toLowerCase();
    }

    public static String reformat(String logLine) {    
        return  message(logLine) + " (" + logLevel(logLine) +")";
    }
}