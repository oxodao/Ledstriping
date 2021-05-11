package fr.oxodao.ledstrip;

import com.google.gson.Gson;
import fr.oxodao.ledstrip.settings.Lifebar;
import fr.oxodao.ledstrip.settings.Spark;

import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;

public class Settings {

    public boolean enabled = true;
    public String url = "http://localhost:8121/";

    public int color = 0x12586E;
    public int brightness = 80;

    public Lifebar lifebar = new Lifebar();
    public Spark spark = new Spark();

    public static Settings load() {
        Settings s = new Settings();

        File file = new File("./config/ledstrip.json");
        Gson gson = new Gson();
        if (file.exists()) {
            try {
                FileReader fileReader = new FileReader(file);
                s = gson.fromJson(fileReader, Settings.class);
                fileReader.close();
            } catch (IOException e) {
                Ledstrip.LOGGER.warn("Could not load ledstrip settings: " + e.getLocalizedMessage());
            }
        }

        return s;
    }

    public void save() {
        Gson gson = new Gson();
        File file = new File("./config/ledstrip.json");
        if (!file.getParentFile().exists()) {
            file.getParentFile().mkdir();
        }
        try {
            FileWriter fileWriter = new FileWriter(file);
            fileWriter.write(gson.toJson(this));
            fileWriter.close();
        } catch (IOException e) {
            Ledstrip.LOGGER.warn("Could not save ledstrip settings: " + e.getLocalizedMessage());
        }
    }

    public static String getColor(int color) {
        return String.format("#%06X", (0xFFFFFF & color));
    }

    public static Settings getSettings() {
        return Ledstrip.getInstance().settings;
    }
}
