import React, {useEffect, useState} from 'react';
import {convertColor, sendCommand} from "./command";
import {List, ListItem, ListItemText, Slider}    from "@material-ui/core";
import {useDebounce}               from "use-debounce";

import styles from './index.module.scss';

const modes: string[] = [
    'Static',
    'Blink',
    'Breath',
    'Color Wipe',
    'Color Wipe Inverse',
    'Color Wipe Reverse',
    'Color Wipe Reverse Inverse',
    'Color Wipe Random',
    'Random Color',
    'Single Dynamic',
    'Multi Dynamic',
    'Rainbow',
    'Rainbow Cycle',
    'Scan',
    'Dual Scan',
    'Fade',
    'Theater Chase',
    'Theater Chase Rainbow',
    'Running Lights',
    'Twinkle',
    'Twinkle Random',
    'Twinkle Fade',
    'Twinkle Fade Random',
    'Sparkle',
    'Flash Sparkle',
    'Hyper Sparkle',
    'Strobe',
    'Strobe Rainbow',
    'Multi Strobe',
    'Blink Rainbow',
    'Chase White',
    'Chase Color',
    'Chase Random',
    'Chase Rainbow',
    'Chase Flash',
    'Chase Flash Random',
    'Chase Rainbow White',
    'Chase Blackout',
    'Chase Blackout Rainbow',
    'Color Sweep Random',
    'Running Color',
    'Running Red Blue',
    'Running Random',
    'Larson Scanner',
    'Comet',
    'Fireworks',
    'Fireworks Random',
    'Merry Christmas',
    'Fire Flicker',
    'Fire Flicker (soft)',
    'Fire Flicker (intense)',
    'Circus Combustus',
    'Halloween',
    'Bicolor Chase',
    'Tricolor Chase',
    'TwinkleFOX',
    'Custom 0',
    'Custom 1',
    'Custom 2',
    'Custom 3',
    'Custom 4',
    'Custom 5',
    'Custom 6',
    'Custom 7',
];

function App() {
    const [ color, setColor ] = useState<string>("");
    const [ debouncedColor ] = useDebounce(color, 500);

    const [ brightness, setBrightness ] = useState<number | number[]>(100);
    const [ debouncedBrightness ] = useDebounce(brightness, 500);

    const [ speed, setSpeed ] = useState<number | number[]>(100);
    const [ debouncedSpeed ] = useDebounce(speed, 500);

    useEffect(() => {
        sendCommand('b ' + debouncedBrightness);
    }, [debouncedBrightness]);

    useEffect(() => {
        sendCommand('c ' + convertColor(debouncedColor));
    }, [debouncedColor]);

    useEffect(() => {
        sendCommand('s ' + debouncedSpeed);
    }, [debouncedSpeed]);

    const setMode = (mode: number) => {
        sendCommand('m ' + mode);
    }

  return (
    <div className={styles.MainApp}>
        <label htmlFor="Color">Color: </label>
        <input id="Color" type="color" value={color} onChange={a => setColor(a.target.value)}/>

        <label htmlFor="Brightness">Brightness: </label>
        <Slider id="Brightness"
                min={0}
                max={255}
                defaultValue={100}
                value={brightness}
                onChange={(_, val) => setBrightness(val)}
                valueLabelDisplay={"auto"}
                marks />

        <label htmlFor="Speed">Speed: </label>
        <Slider id="Speed"
                min={0}
                max={255}
                defaultValue={100}
                value={speed}
                onChange={(_, val) => setSpeed(val)}
                valueLabelDisplay={"auto"}
                marks />

        <label htmlFor="Mode">Mode: </label>
        <List>
            {
                modes.map((val, i) => <ListItem button>
                    <ListItemText onClick={() => setMode(i)}>{val}</ListItemText>
                </ListItem>)
            }
        </List>
    </div>
  );
}

export default App;
