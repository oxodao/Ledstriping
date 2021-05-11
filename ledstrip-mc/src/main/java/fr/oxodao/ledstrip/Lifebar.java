package fr.oxodao.ledstrip;

import net.minecraft.client.MinecraftClient;
import net.minecraft.client.network.ClientPlayerEntity;

import java.awt.*;

public class Lifebar {

    private static float lastHealth = -1.0f;

    public static void setLifebarColor() {
        if (!Settings.getSettings().enabled || !Settings.getSettings().lifebar.enabled) {
            return;
        }

        ClientPlayerEntity player = MinecraftClient.getInstance().player;
        if (player == null) {
            return;
        }

        if (Lifebar.lastHealth < 0 || Lifebar.lastHealth != player.getHealth()) {
            Ledstrip.getInstance().apiClient.setBrightness(Settings.getSettings().lifebar.brightness);
            Ledstrip.getInstance().apiClient.setColor(getColorForHealth(
                    false,
                    player.getHealth(),
                    player.getMaxHealth()
            ));

            Lifebar.lastHealth = player.getHealth();
        }
    }

    private static String getColorForHealth(boolean progressive, float health, float maxHealth) {
        if (progressive) {
            float percentage = health/maxHealth;
            float minColor = 0f; // Red
            float maxColor = .277f; // Green-ish

            Color c = Color.getHSBColor(minColor + ((maxColor - minColor) * percentage), 1f, 1f);
            return String.format( "#%02X%02X%02X", c.getRed(), c.getGreen(), c.getBlue());
        } else {
            float quarter = maxHealth / 4;

            if (health <= quarter)
                return Settings.getColor(Settings.getSettings().lifebar.oneQuarter);

            if (health <= 2 * quarter)
                return Settings.getColor(Settings.getSettings().lifebar.half);

            if (health <= 3 * quarter)
                return Settings.getColor(Settings.getSettings().lifebar.threeQuarter);

            return Settings.getColor(Settings.getSettings().lifebar.full);
        }
    }
}
