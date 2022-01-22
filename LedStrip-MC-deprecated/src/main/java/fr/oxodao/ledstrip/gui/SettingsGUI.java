package fr.oxodao.ledstrip.gui;

import fr.oxodao.ledstrip.Settings;
import me.shedaniel.clothconfig2.api.ConfigBuilder;
import me.shedaniel.clothconfig2.api.ConfigCategory;
import me.shedaniel.clothconfig2.api.ConfigEntryBuilder;
import me.shedaniel.clothconfig2.gui.entries.SubCategoryListEntry;
import me.shedaniel.clothconfig2.impl.builders.SubCategoryBuilder;
import net.minecraft.client.gui.screen.Screen;
import net.minecraft.text.LiteralText;

public class SettingsGUI {
    private ConfigEntryBuilder ceb;

    public Screen getConfigScreen(Screen parent, boolean isTransparent) {
        ConfigBuilder builder = ConfigBuilder.create().setParentScreen(parent).setTitle(new LiteralText("Ledstrip")).setTransparentBackground(isTransparent);
        this.ceb = builder.entryBuilder();
        builder.setSavingRunnable(Settings.getSettings()::save);

        ConfigCategory root = builder.getOrCreateCategory(new LiteralText("Root"));
        root.addEntry(getGeneralCategory());
        root.addEntry(getLifebarCategory());
        root.addEntry(getSparkCategory());

        return builder.build();
    }

    private SubCategoryListEntry getGeneralCategory() {
        SubCategoryBuilder sub = ceb.startSubCategory(new LiteralText("General"));

        sub.add(ceb.startBooleanToggle(new LiteralText("Enabled"), Settings.getSettings().enabled)
                .setDefaultValue(true)
                .setSaveConsumer(val -> Settings.getSettings().enabled = val)
                .build());

        sub.add(ceb.startStrField(new LiteralText("API Url"), Settings.getSettings().url)
                .setDefaultValue("http://localhost:8121/")
                .setSaveConsumer(val -> Settings.getSettings().url = val)
                .build());

        sub.add(ceb.startColorField(new LiteralText("Default color"), Settings.getSettings().color)
                .setDefaultValue(0x12586E)
                .setSaveConsumer(val -> Settings.getSettings().color = val)
                .build());

        sub.add(ceb.startIntSlider(new LiteralText("Default brightness"), Settings.getSettings().brightness, 1, 255)
                .setDefaultValue(80)
                .setSaveConsumer(val -> Settings.getSettings().brightness = val)
                .build());

        return sub.build();
    }

    private SubCategoryListEntry getLifebarCategory() {
        SubCategoryBuilder sub = ceb.startSubCategory(new LiteralText("Lifebar"));

        sub.add(ceb.startBooleanToggle(new LiteralText("Enabled"), Settings.getSettings().lifebar.enabled)
                .setDefaultValue(true)
                .setSaveConsumer(val -> Settings.getSettings().lifebar.enabled = val)
                .build());

        sub.add(ceb.startIntSlider(new LiteralText("Led strip brightness"), Settings.getSettings().lifebar.brightness, 1, 255)
                .setDefaultValue(150)
                .setSaveConsumer(val -> Settings.getSettings().lifebar.brightness = val)
                .build());

        sub.add(ceb.startColorField(new LiteralText("1/4 life"), Settings.getSettings().lifebar.oneQuarter)
                .setDefaultValue(0xFF0000)
                .setSaveConsumer(val -> Settings.getSettings().lifebar.oneQuarter = val)
                .build());

        sub.add(ceb.startColorField(new LiteralText("Half life"), Settings.getSettings().lifebar.half)
                .setDefaultValue(0xFF7300)
                .setSaveConsumer(val -> Settings.getSettings().lifebar.half = val)
                .build());

        sub.add(ceb.startColorField(new LiteralText("3/4 life"), Settings.getSettings().lifebar.threeQuarter)
                .setDefaultValue(0xFFFF00)
                .setSaveConsumer(val -> Settings.getSettings().lifebar.threeQuarter = val)
                .build());

        sub.add(ceb.startColorField(new LiteralText("Full"), Settings.getSettings().lifebar.full)
                .setDefaultValue(0x15FF00)
                .setSaveConsumer(val -> Settings.getSettings().lifebar.full = val)
                .build());

        return sub.build();
    }

    private SubCategoryListEntry getSparkCategory() {
        SubCategoryBuilder sub = ceb.startSubCategory(new LiteralText("Spark"));

        sub.add(ceb.startBooleanToggle(new LiteralText("Mob kill"), Settings.getSettings().spark.mobKill)
                .setDefaultValue(true)
                .setSaveConsumer(val -> Settings.getSettings().spark.mobKill = val)
                .build());

        sub.add(ceb.startColorField(new LiteralText("Mob kill color"), Settings.getSettings().spark.mobKillColor)
                .setDefaultValue(0xFFFFFF)
                .setSaveConsumer(val -> Settings.getSettings().spark.mobKillColor = val)
                .build());

        sub.add(ceb.startIntField(new LiteralText("Mob kill duration"), Settings.getSettings().spark.mobKillDuration)
                .setDefaultValue(100)
                .setSaveConsumer(val -> Settings.getSettings().spark.mobKillDuration = val)
                .build());

        sub.add(ceb.startBooleanToggle(new LiteralText("On death"), Settings.getSettings().spark.death)
                .setDefaultValue(true)
                .setSaveConsumer(val -> Settings.getSettings().spark.death = val)
                .build());

        sub.add(ceb.startColorField(new LiteralText("On death color"), Settings.getSettings().spark.deathColor)
                .setDefaultValue(0xFF0000)
                .setSaveConsumer(val -> Settings.getSettings().spark.deathColor = val)
                .build());

        sub.add(ceb.startIntField(new LiteralText("On death duration"), Settings.getSettings().spark.deathDuration)
                .setDefaultValue(250)
                .setSaveConsumer(val -> Settings.getSettings().spark.deathDuration = val)
                .build());

        return sub.build();
    }

}
