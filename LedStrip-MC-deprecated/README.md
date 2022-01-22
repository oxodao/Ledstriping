# Simple ledstrip mod for Minecraft

## Deprecation
This mod is the deprecated one, as I am currently running my server on Forge, I do not plan to update the Fabric version.

As my only hard dependency is ClothConfig and it exists for both Fabric and Forge, there is a world in which I'll make a multi-jar or something like that.

Not sure if this will happen though. For now, use MCForge, sorry for Fabric modpack users

## Setup

This is a simple mod for the Ledstrip software.

Installing it requires Fabric, Fabric API, Cloth config API and ModMenu to set it up

Just drop those jars in the mods directory then run the game.

Go into the mod settings and populate it with your Bridge API URL (If not remote / custom bridge config)

## Demo

Here is a video demo of the mod in its current state.

https://www.youtube.com/watch?v=wDBF8xzos4k

The only thing not displayed here is a WIP feature to gradually change the color depending on the health because it's not convincing enough yet.

There is a known bug seen in the video that prevents the lifebar from directly being set when loading a world, you need to take damage first.
