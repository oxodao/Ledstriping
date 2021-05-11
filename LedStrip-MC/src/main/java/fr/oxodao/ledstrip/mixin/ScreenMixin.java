package fr.oxodao.ledstrip.mixin;

import fr.oxodao.ledstrip.Ledstrip;
import fr.oxodao.ledstrip.Lifebar;
import fr.oxodao.ledstrip.Settings;
import net.minecraft.client.MinecraftClient;
import net.minecraft.client.gui.screen.Screen;
import org.jetbrains.annotations.Nullable;
import org.spongepowered.asm.mixin.Mixin;
import org.spongepowered.asm.mixin.injection.At;
import org.spongepowered.asm.mixin.injection.Inject;
import org.spongepowered.asm.mixin.injection.callback.CallbackInfo;

@Mixin(MinecraftClient.class)
public class ScreenMixin {

    @Inject(at = @At("HEAD"), method = "openScreen(Lnet/minecraft/client/gui/screen/Screen;)V")
    public void openScreen(@Nullable Screen screen, CallbackInfo ci) {
        if (!Settings.getSettings().enabled)
            return;

        if (MinecraftClient.getInstance().world != null) {
            Lifebar.setLifebarColor();
            return;
        }

        Ledstrip.getInstance().apiClient.setBrightness(Settings.getSettings().brightness);
        Ledstrip.getInstance().apiClient.setColor(Settings.getColor(Settings.getSettings().color));
    }

}
