package fr.oxodao.ledstrip;

import fr.oxodao.ledstrip.gui.SettingsGUI;
import net.fabricmc.api.ModInitializer;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class Ledstrip implements ModInitializer {

	public static final Logger LOGGER = LogManager.getLogger();
	private static Ledstrip instance;

	public Settings settings;
	public ApiCall apiClient;
	public SettingsGUI settingsScreen;

	@Override
	public void onInitialize() {
	    Ledstrip.instance = this;

	    this.settings = Settings.load();
		this.apiClient = new ApiCall();
		this.settingsScreen = new SettingsGUI();

		this.apiClient.setBrightness(Settings.getSettings().brightness);
		this.apiClient.setColor(Settings.getColor(Settings.getSettings().color));
	}

	public static Ledstrip getInstance() {
		return instance;
	}

}
