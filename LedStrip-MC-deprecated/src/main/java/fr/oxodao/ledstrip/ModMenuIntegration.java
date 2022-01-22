package fr.oxodao.ledstrip;

import com.terraformersmc.modmenu.api.ModMenuApi;
import com.terraformersmc.modmenu.api.ConfigScreenFactory;
import net.minecraft.client.MinecraftClient;

public class ModMenuIntegration implements ModMenuApi {
    @Override
    public ConfigScreenFactory<?> getModConfigScreenFactory() {
        return (parent)-> Ledstrip.getInstance().settingsScreen.getConfigScreen(parent, MinecraftClient.getInstance().world != null);
    }
}
