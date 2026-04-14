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
        int levelStart = logLine.indexOf("[");
        int levelEnd = logLine.indexOf("]");
        String level = logLine.substring(levelStart + 1, levelEnd).toLowerCase();
    
        int messageStart = logLine.indexOf(":");
        String message = logLine.substring(messageStart + 1).trim();
    
        return message + " (" + level + ")";
    }
}