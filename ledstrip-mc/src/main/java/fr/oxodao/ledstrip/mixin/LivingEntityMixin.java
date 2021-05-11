package fr.oxodao.ledstrip.mixin;

import fr.oxodao.ledstrip.Ledstrip;
import fr.oxodao.ledstrip.Lifebar;
import fr.oxodao.ledstrip.Settings;
import net.minecraft.client.MinecraftClient;
import net.minecraft.entity.LivingEntity;
import net.minecraft.entity.damage.DamageSource;
import net.minecraft.entity.player.PlayerEntity;
import org.spongepowered.asm.mixin.Mixin;
import org.spongepowered.asm.mixin.injection.At;
import org.spongepowered.asm.mixin.injection.Inject;
import org.spongepowered.asm.mixin.injection.callback.CallbackInfo;

import java.awt.*;

@Mixin(net.minecraft.entity.LivingEntity.class)
public abstract class LivingEntityMixin {

    private static boolean isAlive = true;

    @Inject(at = @At("HEAD"), method = "onDeath(Lnet/minecraft/entity/damage/DamageSource;)V")
    private void entityDeath(DamageSource ds, CallbackInfo info) {
        if (!Settings.getSettings().enabled)
            return;

        if (ds.getAttacker() == null || MinecraftClient.getInstance().player == null) {
            return;
        }

        String playerUuid = MinecraftClient.getInstance().player.getUuidAsString();
        String attackerUuid = ds.getAttacker().getUuidAsString();

        // On killing an entity
        if (Settings.getSettings().spark.mobKill && attackerUuid.equals(playerUuid)) {
            Ledstrip.getInstance().apiClient.spark(
                    Settings.getColor(Settings.getSettings().spark.mobKillColor),
                    Settings.getSettings().spark.mobKillDuration
            );
        }
    }

    @Inject(at=@At("TAIL"), method= "tick()V")
    public void tick(CallbackInfo ci) {
        if (!Settings.getSettings().enabled)
            return;

        /* Do not listen to IDE's complain. That's plain just wrong */
        if (MinecraftClient.getInstance().player == null || !(((LivingEntity)(Object)this) instanceof PlayerEntity)) {
            return;
        }

        // On Death
        if (!MinecraftClient.getInstance().player.isAlive()) {
            if (Settings.getSettings().spark.death && LivingEntityMixin.isAlive) {
                Ledstrip.getInstance().apiClient.spark(
                        Settings.getColor(Settings.getSettings().spark.deathColor),
                        Settings.getSettings().spark.deathDuration
                );
            }

            LivingEntityMixin.isAlive = false;
            return;
        }

        LivingEntityMixin.isAlive = true;

        Lifebar.setLifebarColor();
    }

}
